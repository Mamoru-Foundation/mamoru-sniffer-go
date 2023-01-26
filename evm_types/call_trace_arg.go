package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
)

type CallTraceArg struct {
	CallTraceSeq uint32
	Depth        uint32
	TxIndex      uint32
	BlockIndex   uint64
	Arg          string
}

func NewCallTraceArgData(txs []CallTraceArg) mamoru_sniffer.BlockchainData {
	batch := generated_bindings.NewCallTraceArgBatch()

	for _, tx := range txs {

		generated_bindings.CallTraceArgBatchAppend(
			batch,
			tx.CallTraceSeq,
			tx.TxIndex,
			tx.BlockIndex,
			tx.Arg,
		)
	}

	data := generated_bindings.CallTraceArgBatchFinish(batch)

	return mamoru_sniffer.NewBlockchainData(data)
}
