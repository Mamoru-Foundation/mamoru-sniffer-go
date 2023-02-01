package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
)

type CallTrace struct {
	Seq        uint32
	Depth      uint32
	TxIndex    uint32
	BlockIndex uint64
	Type       string
	From       string
	To         string
	Value      uint64
	GasLimit   uint64
	GasUsed    uint64
	Input      []byte
}

func NewCallTraceData(txs []CallTrace) mamoru_sniffer.BlockchainData {
	batch := generated_bindings.NewCallTraceBatch()

	for _, tx := range txs {
		input := sliceToFfi(tx.Input)
		generated_bindings.CallTraceBatchAppend(
			batch,
			tx.Seq,
			tx.Depth,
			tx.TxIndex,
			tx.BlockIndex,
			tx.Type,
			tx.From,
			tx.To,
			tx.Value,
			tx.GasLimit,
			tx.GasUsed,
			input,
		)
		freeFfiSlice(input)
	}

	data := generated_bindings.CallTraceBatchFinish(batch)

	return mamoru_sniffer.NewBlockchainData(data)
}
