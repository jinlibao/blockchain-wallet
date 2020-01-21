package merkle

import (
	"fmt"

	"github.com/jinlibao/blockchain-wallet/server/hash"
)

func MerkleHash(data [][]byte) []byte {
	hLeaf := make([][]byte, len(data))
	for i, block := range data {
		hLeaf[i] = hash.HashOf(block)
	}
	ln := len(hLeaf)/2 + 1
	hMid := make([][]byte, 0, ln)
	for ln >= 1 {
		for i := 0; i < len(hLeaf)/2; i++ {
			hMid = append(hMid, hash.Keccak256(hLeaf[2*i], hLeaf[2*i+1]))
		}
		if len(hLeaf)%2 == 1 {
			hMid = append(hMid, hash.Keccak256(hLeaf[len(hLeaf)-1]))
		}
		hLeaf = hMid
		ln = len(hLeaf) / 2
		hMid = make([][]byte, 0, ln)
	}
	return hLeaf[0]
}

func dumpSSB(x [][]byte) (s string) {
	s = "["
	com := ""
	for ii := range x {
		st := fmt.Sprintf("%x", x[ii])
		s += com + st
		com = ", "
	}
	s += "]"
	return
}
