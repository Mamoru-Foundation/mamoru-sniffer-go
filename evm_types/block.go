package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
)

/*
	pub struct Block {
	    pub block_index: u64,
	    pub hash: String,
	    pub parent_hash: String,
	    pub state_root: String,
	    pub nonce: u64,
	    pub status: String,
	    pub timestamp: u64,
	    pub block_reward: Vec<u8>,
	    pub fee_recipient: String,
	    pub total_difficulty: u64,
	    pub size: f64,
	    pub gas_used: u64,
	    pub gas_limit: u64,

}
*/
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
