package block

import (
	"fmt"
	"testing"
)

func Test_GenesisBlock(t *testing.T) {
	gb := InitGenesisBlock()
	if !IsGenesisBlock(gb) {
		t.Errorf("Should be genesis block")
	}
}

func Test_InitBlock(t *testing.T) {
	bk := InitBlock(12, "Hello World", []byte{1, 2, 3})
	if IsGenesisBlock(bk) {
		t.Errorf("Should not be genesis block")
	}
	if bk.Index != 12 {
		t.Errorf("Expected index of 12")
	}
	exp := "a5939c6b3f2c404798db8d78bc9de6d2839e8129a9c7ce5ea0c8abe4a35b2eee"
	got := fmt.Sprintf("%x", bk.ThisBlockHash)
	if exp != got {
		t.Errorf("Block hash incorrect\nexpected: ->%s<-\nactual:   ->%s<-\n", exp, got)
	}
}

//
func Test_SerializeBlock(t *testing.T) {
	bk := InitBlock(12, "Welcome to Blockchain Wallet", []byte{1, 2, 3, 4})
	data := SerializeBlock(bk)
	dataStr := fmt.Sprintf("%x", data)
	testDataStr := "57656c636f6d6520746f20426c6f636b636861696e2057616c6c657401020304"
	if dataStr != testDataStr {
		t.Errorf("Invalid data for block\nexpected: ->%s<-\nactual:   ->%s<-\n", testDataStr, dataStr)
	}
}

//
func Test_SerializeForSeal(t *testing.T) {
	bk := InitBlock(12, "Welcome to Blockchain Wallet", []byte{1, 2, 3, 4})
	data := SerializeForSeal(bk)
	dataStr := fmt.Sprintf("%x", data)
	testDataStr := "57656c636f6d6520746f20426c6f636b636861696e2057616c6c6574ca519928ad23782a57400cad460c4f8baf207ff6157269c297cf2a00238ad635010203040000000000000000"
	if dataStr != testDataStr {
		t.Errorf("Invalid data for block\nexpected: ->%s<-\nactual:   ->%s<-\n", testDataStr, dataStr)
	}
}
