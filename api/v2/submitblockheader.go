package v2

import (
	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/phase0"
)

// SubmitBlockHeaderRequest is the v2 request from the builder to submit a block.
type SubmitBlockHeaderRequest struct {
	Message                *v1.BidTrace
	ExecutionPayloadHeader *capella.ExecutionPayloadHeader
	Signature              phase0.BLSSignature `ssz-size:"96"`
}
