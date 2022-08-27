// Copyright © 2022 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/attestantio/go-builder-client/api"
	consensusapi "github.com/attestantio/go-eth2-client/api"
	consensusapiv1 "github.com/attestantio/go-eth2-client/api/v1"
	consensusspec "github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/pkg/errors"
)

// UnblindBlock unblinds a block.
func (s *Service) UnblindBlock(ctx context.Context,
	block *consensusapi.VersionedSignedBlindedBeaconBlock,
) (
	*consensusspec.VersionedSignedBeaconBlock,
	error,
) {
	started := time.Now()

	if block == nil {
		return nil, errors.New("no block supplied")
	}

	switch block.Version {
	case consensusspec.DataVersionBellatrix:
		return s.unblindBellatrixBlock(ctx, started, block.Bellatrix)
	default:
		return nil, fmt.Errorf("unhandled data version %v", block.Version)
	}
}

func (s *Service) unblindBellatrixBlock(ctx context.Context,
	started time.Time,
	block *consensusapiv1.SignedBlindedBeaconBlock,
) (
	*consensusspec.VersionedSignedBeaconBlock,
	error,
) {
	specJSON, err := json.Marshal(block)
	if err != nil {
		monitorOperation(s.Address(), "unblind block", false, time.Since(started))
		return nil, errors.Wrap(err, "failed to marshal JSON")
	}

	contentType, respBodyReader, err := s.post(ctx, "/eth/v1/builder/blinded_blocks", ContentTypeJSON, bytes.NewBuffer(specJSON))
	if err != nil {
		monitorOperation(s.Address(), "unblind block", false, time.Since(started))
		return nil, errors.Wrap(err, "failed to submit unblind block request")
	}

	var dataBodyReader bytes.Buffer
	metadataReader := io.TeeReader(respBodyReader, &dataBodyReader)
	var metadata responseMetadata
	if err := json.NewDecoder(metadataReader).Decode(&metadata); err != nil {
		monitorOperation(s.Address(), "unblind block", false, time.Since(started))
		return nil, errors.Wrap(err, "failed to parse response")
	}
	res := &consensusspec.VersionedSignedBeaconBlock{
		Version: consensusspec.DataVersionBellatrix,
		Bellatrix: &bellatrix.SignedBeaconBlock{
			Message: &bellatrix.BeaconBlock{
				Slot:          block.Message.Slot,
				ProposerIndex: block.Message.ProposerIndex,
				ParentRoot:    block.Message.ParentRoot,
				StateRoot:     block.Message.StateRoot,
				Body: &bellatrix.BeaconBlockBody{
					RANDAOReveal:      block.Message.Body.RANDAOReveal,
					ETH1Data:          block.Message.Body.ETH1Data,
					Graffiti:          block.Message.Body.Graffiti,
					ProposerSlashings: block.Message.Body.ProposerSlashings,
					AttesterSlashings: block.Message.Body.AttesterSlashings,
					Attestations:      block.Message.Body.Attestations,
					Deposits:          block.Message.Body.Deposits,
					VoluntaryExits:    block.Message.Body.VoluntaryExits,
					SyncAggregate:     block.Message.Body.SyncAggregate,
				},
			},
			Signature: block.Signature,
		},
	}

	switch contentType {
	case ContentTypeJSON:
		var resp api.VersionedExecutionPayload
		if err := json.NewDecoder(&dataBodyReader).Decode(&resp); err != nil {
			return nil, errors.Wrap(err, "failed to parse bellatrix response")
		}
		// Ensure that the data returned is what we expect.
		ourExecutionPayloadHash, err := block.Message.Body.ExecutionPayloadHeader.HashTreeRoot()
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate hash tree root for our execution payload header")
		}
		receivedExecutionPayloadHash, err := resp.Bellatrix.HashTreeRoot()
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate hash tree root for the received execution payload")
		}
		if !bytes.Equal(ourExecutionPayloadHash[:], receivedExecutionPayloadHash[:]) {
			return nil, fmt.Errorf("execution payload hash mismatch: %#x != %#x", receivedExecutionPayloadHash[:], ourExecutionPayloadHash[:])
		}
		res.Bellatrix.Message.Body.ExecutionPayload = resp.Bellatrix
	default:
		return nil, fmt.Errorf("unsupported content type %v", contentType)
	}
	monitorOperation(s.Address(), "unblind block", true, time.Since(started))
	return res, nil
}
