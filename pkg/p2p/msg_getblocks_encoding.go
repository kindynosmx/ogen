// Code generated by fastssz. DO NOT EDIT.
// Hash: 5cd34b045feacd7dacf0dc09fba0bb0f9eb91f646be77ad5c600c6e74746af33
package p2p

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the MsgGetBlocks object
func (m *MsgGetBlocks) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the MsgGetBlocks object to a target array
func (m *MsgGetBlocks) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'LastBlockHash'
	dst = append(dst, m.LastBlockHash[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the MsgGetBlocks object
func (m *MsgGetBlocks) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 32 {
		return ssz.ErrSize
	}

	// Field (0) 'LastBlockHash'
	copy(m.LastBlockHash[:], buf[0:32])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the MsgGetBlocks object
func (m *MsgGetBlocks) SizeSSZ() (size int) {
	size = 32
	return
}

// HashTreeRoot ssz hashes the MsgGetBlocks object
func (m *MsgGetBlocks) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the MsgGetBlocks object with a hasher
func (m *MsgGetBlocks) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'LastBlockHash'
	hh.PutBytes(m.LastBlockHash[:])

	hh.Merkleize(indx)
	return
}
