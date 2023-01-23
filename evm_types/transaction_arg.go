package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
)

type TransactionArg struct {
	TxIndex uint32
	Arg     string
}

func NewTransactionArgData(txs []TransactionArg) mamoru_sniffer.BlockchainData {
	batch := generated_bindings.NewTransactionArgBatch()

	for _, tx := range txs {

		generated_bindings.TransactionArgBatchAppend(
			batch,
			tx.TxIndex,
			tx.Arg,
		)

	}

	data := generated_bindings.TransactionArgBatchFinish(batch)

	return mamoru_sniffer.NewBlockchainData(data)
}
