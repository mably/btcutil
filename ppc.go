// Copyright (c) 2014-2014 PPCD developers.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcutil

import (
	"bytes"
	"github.com/conformal/btclog"
	"github.com/mably/btcwire"
	"time"
)

func (b *Block) Meta() *btcwire.Meta {
	if b.meta != nil {
		return b.meta
	}
	b.meta = new(btcwire.Meta)
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
func NewBlockWithMetas(msgBlock *btcwire.MsgBlock, meta *btcwire.Meta) *Block {
	return &Block{
		msgBlock:    msgBlock,
		blockHeight: BlockHeightUnknown,
		meta: meta,
	}
}

// https://github.com/ppcoin/ppcoin/blob/v0.4.0ppc/src/main.h#L962
// ppcoin: two types of block: proof-of-work or proof-of-stake
func (block *Block) IsProofOfStake() bool {
	return block.msgBlock.IsProofOfStake()
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