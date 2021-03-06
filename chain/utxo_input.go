package chain

import (
	"bytes"
	"github.com/michain/dotcoin/wallet"
	"github.com/michain/dotcoin/config/chainhash"
	"strconv"
)

// OutPoint defines a dotcoin data type that is used to track previous
// transaction outputs.
type OutPoint struct {
	Hash  chainhash.Hash
	Index int
}

func (o OutPoint) StringHash() string {
	return o.Hash.String()
}

// NewOutPoint returns a new dotcoin transaction outpoint point with the
// provided hash and index.
func NewOutPoint(hash *chainhash.Hash, index int) *OutPoint {
	return &OutPoint{
		Hash:  *hash,
		Index: index,
	}
}

// String returns the OutPoint in the human-readable form "hash:index".
func (o OutPoint) String() string {
	buf := make([]byte, 2*chainhash.HashSize+1, 2*chainhash.HashSize+1+10)
	copy(buf, o.Hash.String())
	buf[2*chainhash.HashSize] = ':'
	buf = strconv.AppendUint(buf, uint64(o.Index), 10)
	return string(buf)
}


// TXInput represents a transaction input
type TXInput struct {
	PreviousOutPoint OutPoint
	Signature []byte
	PubKey    []byte
}



// UsesKey checks whether the address initiated the transaction
func (in *TXInput) UnLock(pubKeyHash []byte) bool {
	lockingHash := wallet.HashPublicKey(in.PubKey)
	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

// NewTXInput create a new TXInput
func NewTXInput(prevOut *OutPoint, sign, pubKey []byte) *TXInput{
	return &TXInput{*prevOut, sign, pubKey}
}



