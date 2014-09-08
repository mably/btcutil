package btcutil

import (
	"bytes"
	"github.com/mably/btcwire"
	"io"
	"math/big"
)

type Meta struct {
	generatedStakeModifier bool
	stakeModifier          uint64
	stakeModifierChecksum  uint32 // checksum of index; in-memeory only (main.h)
	hashProofOfStake       *btcwire.ShaHash
	stakeEntropyBit        uint32
	flags                  uint32
	chainTrust             *big.Int
}

func (msg *Meta) Serialize(w io.Writer) error {
	return nil
}

func (msg *Meta) Deserialize(r io.Reader) error {
	return nil
}

func (b *Block) Meta() *Meta {
	return &b.meta
}

func NewBlockFromBytesWithMeta(serializedBlock []byte) (*Block, error) {
	br := bytes.NewReader(serializedBlock)
	b, err := NewBlockFromReader(br)
	if err != nil {
		return nil, err
	}
	err = b.meta.Deserialize(br)
	if err != nil {
		return nil, err
	}
	b.serializedBlock = serializedBlock
	return b, nil
}

func (b *Block) BytesWithMeta() ([]byte, error) {
	// Return the cached serialized bytes if it has already been generated.
	/*if false & len(b.serializedBlock) != 0 {
		return b.serializedBlock, nil
	}*/

	// Serialize the MsgBlock.
	var w bytes.Buffer
	err := b.msgBlock.Serialize(&w)
	if err != nil {
		return nil, err
	}

	serializedBlock := w.Bytes()

	// Serialize Meta.
	err = b.meta.Serialize(&w)
	if err != nil {
		return nil, err
	}

	serializedBlockWithMeta := w.Bytes()

	// Cache the serialized bytes and return them.
	b.serializedBlock = serializedBlock
	return serializedBlockWithMeta, nil
}
