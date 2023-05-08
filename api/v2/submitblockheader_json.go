package v2

import (
	"encoding/hex"
	"encoding/json"
	"strings"

	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

// submitBlockRequestJSON is the spec representation of the struct.
type submitBlockHeaderRequestJSON struct {
	Message                *v1.BidTrace                    `json:"message"`
	ExecutionPayloadHeader *capella.ExecutionPayloadHeader `json:"execution_payload_header"`
	Signature              string                          `json:"signature"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *SubmitBlockHeaderRequest) UnmarshalJSON(input []byte) error {
	var data submitBlockHeaderRequestJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return errors.Wrap(err, "invalid JSON")
	}
	return s.unpack(&data)
}

func (s *SubmitBlockHeaderRequest) unpack(data *submitBlockHeaderRequestJSON) error {
	// field: Message
	if data.Message == nil {
		return errors.New("message missing")
	}
	s.Message = data.Message

	// field: ExecutionPayloadHeader
	if data.ExecutionPayloadHeader == nil {
		return errors.New("execution payload header missing")
	}
	s.ExecutionPayloadHeader = data.ExecutionPayloadHeader

	// field: Signature
	if data.Signature == "" {
		return errors.New("signature missing")
	}
	signature, err := hex.DecodeString(strings.TrimPrefix(data.Signature, "0x"))
	if err != nil {
		return errors.Wrap(err, "invalid signature")
	}
	if len(signature) != phase0.SignatureLength {
		return errors.New("incorrect length for signature")
	}
	copy(s.Signature[:], signature)

	return nil
}
