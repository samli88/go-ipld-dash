package iplddash

import (
	"encoding/hex"
	"io/ioutil"
	"testing"
)

func TestBlockMessageDecoding(t *testing.T) {
	hexdata, err := ioutil.ReadFile("fixtures/block.hex")
	if err != nil {
		t.Fatal(err)
	}

	data, err := hex.DecodeString(string(hexdata[:len(hexdata)-1]))
	if err != nil {
		t.Fatal(err)
	}

	nodes, err := DecodeBlockMessage(data)
	if err != nil {
		t.Fatal(err)
	}

	expblk := "000000000000003fd5ab15f68eaf0e16a304dd15079638ae270b1f06aad1af5f"
	if nodes[0].(*Block).HexHash() != expblk {
		t.Fatal("parsed incorrectly")
	}

	blk, _, err := nodes[0].ResolveLink([]string{"tx"})
	if err != nil {
		t.Fatal(err)
	}

	if !blk.Cid.Equals(nodes[len(nodes)-1].Cid()) {
		t.Fatal("merkle root looks wrong")
	}
}

func TestDecodingNoTxs(t *testing.T) {
	hexdata := "020000003f3572d8d15aca3f37279470b505f73806cc8e1ef7c23bf0b7f942d4160f00006345fa110fa94b3bde54bfea8f14b0473033869bcc606eda15f3388e201fce1cad4cdb52f0ff0f1e6c1500000101000000010000000000000000000000000000000000000000000000000000000000000000ffffffff0a5a0105062f503253482fffffffff0100743ba40b000000232102360b177e5d3402b7e7daf439cbaed6e686749cbebb47423ec01547ca7aa734f5ac00000000"

	data, err := hex.DecodeString(hexdata)
	if err != nil {
		t.Fatal(err)
	}

	nodes, err := DecodeBlockMessage(data)
	if err != nil {
		t.Fatal(err)
	}

	expblk := "00000a7da4d35c6663d68c04f9884dd01baaee9a2fdee8c02ef9d2f13f8b90ac"

	blk := nodes[0].(*Block)
	if nodes[0].(*Block).HexHash() != expblk {
		t.Fatal("parsed incorrectly")
	}

	tx := nodes[1].(*Tx)
	t.Log(blk.MerkleRoot)
	t.Log(tx.Cid())
}

func TestTxDecoding(t *testing.T) {
	txdata := "0200000001337c2e4244931eeb4cdd9f6e604780d897c118fb3bf9790cf6808290eb45398d010000006b48304502210094d6271f6879b50e07f49d6b0d0df1541355a2fb518c805ced8a38408dd4d3d002207a56839d8cb16ae9c8eee96e5948bf035d943ab7d7247e2e20bcbb4cc15d942201210293405b9d6db1ae85b5f2a7006ad68581915667810a07da73f09c78050c9f8414feffffff0240a91b00000000001976a9144e2b02348708c5badcc771856cd683bb6efb9ced88ac1c7bf81e000000001976a914e9ea5a6a38e2401c3cce4928117675e8d74c7a8788ac7f7c0e00"

	data, err := hex.DecodeString(txdata)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := DecodeTx(data)
	if err != nil {
		t.Fatal(err)
	}

	if tx.LockTime != 949375 {
		t.Fatal("lock time incorrect")
	}

	if tx.Inputs[0].SeqNo != 4294967294 {
		t.Fatal("seqno not right")
	}
}
