package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
)

type Block struct {
	BlockIndex      uint64
	Hash            string
	ParentHash      string
	StateRoot       string
	Nonce           uint64
	Status          string
	Timestamp       uint64
	BlockReward     []byte
	FeeRecipient    string
	TotalDifficulty uint64
	Size            float64
	GasUsed         uint64
	GasLimit        uint64
}

func NewBlockData(txs []Block) mamoru_sniffer.BlockchainData {
	batch := generated_bindings.NewBlockBatch()

	for _, tx := range txs {
		blockReward := sliceToFfi(tx.BlockReward)
		generated_bindings.BlockBatchAppend(
			batch,
			tx.BlockIndex,
			tx.Hash,
			tx.ParentHash,
			tx.StateRoot,
			tx.Nonce,
			tx.Status,
			tx.Timestamp,
			blockReward,
			tx.FeeRecipient,
			tx.TotalDifficulty,
			tx.Size,
			tx.GasUsed,
			tx.GasLimit,
		)
		freeFfiSlice(blockReward)
	}

	data := generated_bindings.BlockBatchFinish(batch)

	return mamoru_sniffer.NewBlockchainData(data)
}
