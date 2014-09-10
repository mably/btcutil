package btcutil

import (
	"bytes"
	"encoding/binary"
	"github.com/conformal/btclog"
	"github.com/mably/btcwire"
	"io"
	"math/big"
	"time"
)

type Meta struct {
	StakeModifier         uint64
	StakeModifierChecksum uint32 // checksum of index; in-memeory only (main.h)
	HashProofOfStake      btcwire.ShaHash
	Flags                 uint32
	ChainTrust            big.Int
}

func (m *Meta) Serialize(w io.Writer) error {
	e := binary.Write(w, binary.LittleEndian, &m.StakeModifier)
	if e != nil {
		return e
	}
	binary.Write(w, binary.LittleEndian, &m.StakeModifierChecksum)
	if e != nil {
		return e
	}
	binary.Write(w, binary.LittleEndian, &m.Flags)
	if e != nil {
		return e
	}
	binary.Write(w, binary.LittleEndian, &m.HashProofOfStake)
	if e != nil {
		return e
	}
	bytes := m.ChainTrust.Bytes()
	var blen byte
	blen = byte(len(bytes))
	binary.Write(w, binary.LittleEndian, &blen)
	if e != nil {
		return e
	}
	binary.Write(w, binary.LittleEndian, &bytes)
	if e != nil {
		return e
	}
	return nil
}

func (m *Meta) Deserialize(r io.Reader) error {
	e := binary.Read(r, binary.LittleEndian, &m.StakeModifier)
	if e != nil {
		return e
	}
	e = binary.Read(r, binary.LittleEndian, &m.StakeModifierChecksum)
	if e != nil {
		return e
	}
	e = binary.Read(r, binary.LittleEndian, &m.Flags)
	if e != nil {
		return e
	}
	e = binary.Read(r, binary.LittleEndian, &m.HashProofOfStake)
	if e != nil {
		return e
	}

	var blen byte
	e = binary.Read(r, binary.LittleEndian, &blen)
	if e != nil {
		return e
	}
	var arr = make([]byte, blen)
	e = binary.Read(r, binary.LittleEndian, &arr)
	if e != nil {
		return e
	}
	m.ChainTrust.SetBytes(arr)
	return nil
}

func (b *Block) Meta() *Meta {
	if b.meta != nil {
		return b.meta
	}
	b.meta = new(Meta)
	return b.meta
}

func NewBlockFromBytesWithMeta(serializedBlock []byte) (*Block, error) {
	br := bytes.NewReader(serializedBlock)
	b, err := NewBlockFromReader(br)
	if err != nil {
		return nil, err
	}
	err = b.Meta().Deserialize(br)
	if err != nil {
		return nil, err
	}
	//b.serializedBlock = serializedBlock
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
	err = b.Meta().Serialize(&w)
	if err != nil {
		return nil, err
	}

	serializedBlockWithMeta := w.Bytes()

	// Cache the serialized bytes and return them.
	b.serializedBlock = serializedBlock

	return serializedBlockWithMeta, nil
}

// NewBlock returns a new instance of a bitcoin block given an underlying
// btcwire.MsgBlock.  See Block.
func NewBlockWithMetas(msgBlock *btcwire.MsgBlock, meta *Meta) *Block {
	return &Block{
		msgBlock:    msgBlock,
		blockHeight: BlockHeightUnknown,
		meta: meta,
	}
}

func Now() time.Time {
    return time.Now()
}

func TimeTrack(log btclog.Logger, start time.Time, name string) {
    elapsed := time.Since(start)
    log.Tracef("%s took %s", name, elapsed)
}

func Slice(args ...interface{}) []interface{} {
    return args
}