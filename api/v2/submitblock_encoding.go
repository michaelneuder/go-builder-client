// Code generated by fastssz. DO NOT EDIT.
// Hash: eeb4b8b9825e95d530784e497b3266dda0c91eecedd1a92edda67ffa6ddefe97
// Version: 0.1.3
package v2

import (
	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/capella"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SubmitBlockRequest object
func (s *SubmitBlockRequest) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SubmitBlockRequest object to a target array
func (s *SubmitBlockRequest) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(344)

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(v1.BidTrace)
	}
	if dst, err = s.Message.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (1) 'ExecutionPayloadHeader'
	dst = ssz.WriteOffset(dst, offset)
	if s.ExecutionPayloadHeader == nil {
		s.ExecutionPayloadHeader = new(capella.ExecutionPayloadHeader)
	}
	offset += s.ExecutionPayloadHeader.SizeSSZ()

	// Field (2) 'Signature'
	dst = append(dst, s.Signature[:]...)

	// Offset (3) 'Transactions'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.Transactions); ii++ {
		offset += 4
		offset += len(s.Transactions[ii])
	}

	// Offset (4) 'Withdrawals'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.Withdrawals) * 44

	// Field (1) 'ExecutionPayloadHeader'
	if dst, err = s.ExecutionPayloadHeader.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (3) 'Transactions'
	if size := len(s.Transactions); size > 1073741824 {
		err = ssz.ErrListTooBigFn("SubmitBlockRequest.Transactions", size, 1073741824)
		return
	}
	{
		offset = 4 * len(s.Transactions)
		for ii := 0; ii < len(s.Transactions); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += len(s.Transactions[ii])
		}
	}
	for ii := 0; ii < len(s.Transactions); ii++ {
		if size := len(s.Transactions[ii]); size > 1073741824 {
			err = ssz.ErrBytesLengthFn("SubmitBlockRequest.Transactions[ii]", size, 1073741824)
			return
		}
		dst = append(dst, s.Transactions[ii]...)
	}

	// Field (4) 'Withdrawals'
	if size := len(s.Withdrawals); size > 16 {
		err = ssz.ErrListTooBigFn("SubmitBlockRequest.Withdrawals", size, 16)
		return
	}
	for ii := 0; ii < len(s.Withdrawals); ii++ {
		if dst, err = s.Withdrawals[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the SubmitBlockRequest object
func (s *SubmitBlockRequest) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 344 {
		return ssz.ErrSize
	}

	tail := buf
	var o1, o3, o4 uint64

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(v1.BidTrace)
	}
	if err = s.Message.UnmarshalSSZ(buf[0:236]); err != nil {
		return err
	}

	// Offset (1) 'ExecutionPayloadHeader'
	if o1 = ssz.ReadOffset(buf[236:240]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 344 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (2) 'Signature'
	copy(s.Signature[:], buf[240:336])

	// Offset (3) 'Transactions'
	if o3 = ssz.ReadOffset(buf[336:340]); o3 > size || o1 > o3 {
		return ssz.ErrOffset
	}

	// Offset (4) 'Withdrawals'
	if o4 = ssz.ReadOffset(buf[340:344]); o4 > size || o3 > o4 {
		return ssz.ErrOffset
	}

	// Field (1) 'ExecutionPayloadHeader'
	{
		buf = tail[o1:o3]
		if s.ExecutionPayloadHeader == nil {
			s.ExecutionPayloadHeader = new(capella.ExecutionPayloadHeader)
		}
		if err = s.ExecutionPayloadHeader.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (3) 'Transactions'
	{
		buf = tail[o3:o4]
		num, err := ssz.DecodeDynamicLength(buf, 1073741824)
		if err != nil {
			return err
		}
		s.Transactions = make([]bellatrix.Transaction, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if len(buf) > 1073741824 {
				return ssz.ErrBytesLength
			}
			if cap(s.Transactions[indx]) == 0 {
				s.Transactions[indx] = bellatrix.Transaction(make([]byte, 0, len(buf)))
			}
			s.Transactions[indx] = append(s.Transactions[indx], buf...)
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (4) 'Withdrawals'
	{
		buf = tail[o4:]
		num, err := ssz.DivideInt2(len(buf), 44, 16)
		if err != nil {
			return err
		}
		s.Withdrawals = make([]capella.Withdrawal, num)
		for ii := 0; ii < num; ii++ {
			if err = s.Withdrawals[ii].UnmarshalSSZ(buf[ii*44 : (ii+1)*44]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SubmitBlockRequest object
func (s *SubmitBlockRequest) SizeSSZ() (size int) {
	size = 344

	// Field (1) 'ExecutionPayloadHeader'
	if s.ExecutionPayloadHeader == nil {
		s.ExecutionPayloadHeader = new(capella.ExecutionPayloadHeader)
	}
	size += s.ExecutionPayloadHeader.SizeSSZ()

	// Field (3) 'Transactions'
	for ii := 0; ii < len(s.Transactions); ii++ {
		size += 4
		size += len(s.Transactions[ii])
	}

	// Field (4) 'Withdrawals'
	size += len(s.Withdrawals) * 44

	return
}

// HashTreeRoot ssz hashes the SubmitBlockRequest object
func (s *SubmitBlockRequest) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SubmitBlockRequest object with a hasher
func (s *SubmitBlockRequest) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(v1.BidTrace)
	}
	if err = s.Message.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'ExecutionPayloadHeader'
	if err = s.ExecutionPayloadHeader.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Signature'
	hh.PutBytes(s.Signature[:])

	// Field (3) 'Transactions'
	{
		subIndx := hh.Index()
		num := uint64(len(s.Transactions))
		if num > 1073741824 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.Transactions {
			{
				elemIndx := hh.Index()
				byteLen := uint64(len(elem))
				if byteLen > 1073741824 {
					err = ssz.ErrIncorrectListSize
					return
				}
				hh.AppendBytes32(elem)
				hh.MerkleizeWithMixin(elemIndx, byteLen, (1073741824+31)/32)
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 1073741824)
	}

	// Field (4) 'Withdrawals'
	{
		subIndx := hh.Index()
		num := uint64(len(s.Withdrawals))
		if num > 16 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.Withdrawals {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 16)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SubmitBlockRequest object
func (s *SubmitBlockRequest) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
