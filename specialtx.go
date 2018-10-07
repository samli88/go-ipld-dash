package iplddash

import (
	"bufio"
	"bytes"
	"fmt"
)

// SpecialTxType are the registerd DIP2 special transaction types.
// https://github.com/dashpay/dips/blob/master/dip-0002-special-transactions.md
type SpecialTxType uint16

// Define the registered Special TX types.
const (
	ClassicTx   SpecialTxType = 0 // always reserved to distinguish classical TXes
	ProRegTx    SpecialTxType = 1
	ProUpServTx SpecialTxType = 2
	ProUpRegTx  SpecialTxType = 3
	ProUpRevTx  SpecialTxType = 4
	CoinbaseTx  SpecialTxType = 5
)

//type ProRegTxPayload struct {
//	version         uint16
//	mnType          uint16
//	mnMode          uint16
//	collateralIndex uint32
//	ipAddress       []byte
//	port            uint16
//	ownerKeyID      []byte
//	operatorKeyID   []byte
//	voterKeyID      []byte
//	operatorReward  uint16
//	scriptPayout    []byte
//	inputsHash      []byte
//	payloadSig      []byte
//}

func readSpecialTxPayload(r *bufio.Reader, txType SpecialTxType) ([]byte, error) {
	var payload []byte

	switch txType {
	case ProRegTx:
		payload, err := readProRegTx(r)
		if err != nil {
			return payload, err
		}
	case ProUpServTx:
		payload, err := readProUpServTx(r)
		if err != nil {
			return payload, err
		}
	case ProUpRegTx:
		payload, err := readProUpRegTx(r)
		if err != nil {
			return payload, err
		}
	case ProUpRevTx:
		payload, err := readProUpRevTx(r)
		if err != nil {
			return payload, err
		}
	case CoinbaseTx:
		payload, err := readCoinbaseTx(r)
		if err != nil {
			return payload, err
		}
	default:
		return payload, fmt.Errorf("invalid special transaction type: %v", txType)
	}

	return payload, nil
}

func readProRegTx(r *bufio.Reader) ([]byte, error) {
	var buf bytes.Buffer

	data, err := readFixedSlice(r, 90)
	if err != nil {
		return nil, fmt.Errorf("data: %s", err)
	}
	buf.Write(data)

	scriptPayout, err := readVarSlice(r)
	if err != nil {
		return nil, fmt.Errorf("scriptPayout: %s", err)
	}
	buf.Write(scriptPayout)

	inputsHash, err := readFixedSlice(r, 32)
	if err != nil {
		return nil, fmt.Errorf("inputsHash: %s", err)
	}
	buf.Write(inputsHash)

	payloadSig, err := readVarSlice(r)
	if err != nil {
		return nil, fmt.Errorf("payloadSig: %s", err)
	}
	buf.Write(payloadSig)

	return buf.Bytes(), nil
}

func readProUpServTx(r *bufio.Reader) ([]byte, error) {
	var buf bytes.Buffer

	version, err := readFixedSlice(r, 2)
	if err != nil {
		return nil, fmt.Errorf("version: %s", err)
	}
	buf.Write(version)

	proTxHash, err := readFixedSlice(r, 32)
	if err != nil {
		return nil, fmt.Errorf("proTxHash: %s", err)
	}
	buf.Write(proTxHash)

	ipAddress, err := readFixedSlice(r, 16)
	if err != nil {
		return nil, fmt.Errorf("ipAddress: %s", err)
	}
	buf.Write(ipAddress)

	port, err := readFixedSlice(r, 2)
	if err != nil {
		return nil, fmt.Errorf("port: %s", err)
	}
	buf.Write(port)

	scriptPayout, err := readVarSlice(r)
	if err != nil {
		return nil, fmt.Errorf("scriptPayout: %s", err)
	}
	buf.Write(scriptPayout)

	inputsHash, err := readFixedSlice(r, 32)
	if err != nil {
		return nil, fmt.Errorf("inputsHash: %s", err)
	}
	buf.Write(inputsHash)

	payloadSig, err := readVarSlice(r)
	if err != nil {
		return nil, fmt.Errorf("payloadSig: %s", err)
	}
	buf.Write(payloadSig)

	return buf.Bytes(), nil
}

func readProUpRegTx(r *bufio.Reader) ([]byte, error) {
	var buf bytes.Buffer

	version, err := readFixedSlice(r, 2)
	if err != nil {
		return nil, fmt.Errorf("version: %s", err)
	}
	buf.Write(version)

	proTxHash, err := readFixedSlice(r, 32)
	if err != nil {
		return nil, fmt.Errorf("proTxHash: %s", err)
	}
	buf.Write(proTxHash)

	mode, err := readFixedSlice(r, 2)
	if err != nil {
		return nil, fmt.Errorf("mode: %s", err)
	}
	buf.Write(mode)

	operatorKeyID, err := readFixedSlice(r, 20)
	if err != nil {
		return nil, fmt.Errorf("operatorKeyID: %s", err)
	}
	buf.Write(operatorKeyID)

	voterKeyID, err := readFixedSlice(r, 20)
	if err != nil {
		return nil, fmt.Errorf("voterKeyID: %s", err)
	}
	buf.Write(voterKeyID)

	scriptPayout, err := readVarSlice(r)
	if err != nil {
		return nil, fmt.Errorf("scriptPayout: %s", err)
	}
	buf.Write(scriptPayout)

	payloadSig, err := readVarSlice(r)
	if err != nil {
		return nil, fmt.Errorf("payloadSig: %s", err)
	}
	buf.Write(payloadSig)

	return buf.Bytes(), nil
}

func readProUpRevTx(r *bufio.Reader) ([]byte, error) {
	var buf bytes.Buffer

	version, err := readFixedSlice(r, 2)
	if err != nil {
		return nil, fmt.Errorf("version: %s", err)
	}
	buf.Write(version)

	proTxHash, err := readFixedSlice(r, 32)
	if err != nil {
		return nil, fmt.Errorf("proTxHash: %s", err)
	}
	buf.Write(proTxHash)

	reason, err := readFixedSlice(r, 2)
	if err != nil {
		return nil, fmt.Errorf("reason: %s", err)
	}
	buf.Write(reason)

	inputsHash, err := readFixedSlice(r, 32)
	if err != nil {
		return nil, fmt.Errorf("inputsHash: %s", err)
	}
	buf.Write(inputsHash)

	payloadSig, err := readVarSlice(r)
	if err != nil {
		return nil, fmt.Errorf("payloadSig: %s", err)
	}
	buf.Write(payloadSig)

	return buf.Bytes(), nil
}

func readCoinbaseTx(r *bufio.Reader) ([]byte, error) {
	var buf bytes.Buffer

	version, err := readFixedSlice(r, 2)
	if err != nil {
		return nil, fmt.Errorf("version: %s", err)
	}
	buf.Write(version)

	height, err := readFixedSlice(r, 4)
	if err != nil {
		return nil, fmt.Errorf("height: %s", err)
	}
	buf.Write(height)

	merkleRootMNList, err := readFixedSlice(r, 32)
	if err != nil {
		return nil, fmt.Errorf("merkleRootMNList: %s", err)
	}
	buf.Write(merkleRootMNList)

	return buf.Bytes(), nil
}
