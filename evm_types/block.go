package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
)

type Block struct {
	BlockIndex                 uint64
	Hash                       string
	ParentHash                 string
	StateRoot                  string
	Nonce                      uint64
	Status                     string
	Timestamp                  uint64
	BlockReward                []byte
	FeeRecipient               string
	TotalDifficulty            uint64
	Size                       float64
	GasUsed                    uint64
	GasLimit                   uint64
	BurntFees                  []byte
	PosProposedOnTime          uint32
	PosSlot                    uint32
	PosEpoch                   uint32
	PosProposerIndex           uint32
	PosSlotRootHash            []byte
	PosBeaconChainDepositCount uint32
	PosSlotGraffiti            []byte
	PosBlockRandomness         []byte
	PosRandomReveal            []byte
}

func NewBlockData(txs []Block) mamoru_sniffer.BlockchainData {
	batch := generated_bindings.NewBlockBatch()

	for _, tx := range txs {
		blockReward := sliceToFfi(tx.BlockReward)
		burntFees := sliceToFfi(tx.BurntFees)
		posSlotRootHash := sliceToFfi(tx.PosSlotRootHash)
		posSlotGraffiti := sliceToFfi(tx.PosSlotGraffiti)
		posBlockRandomness := sliceToFfi(tx.PosBlockRandomness)
		posRandomReveal := sliceToFfi(tx.PosRandomReveal)

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
			burntFees,
			tx.PosProposedOnTime,
			tx.PosSlot,
			tx.PosEpoch,
			tx.PosProposerIndex,
			posSlotRootHash,
			tx.PosBeaconChainDepositCount,
			posSlotGraffiti,
			posBlockRandomness,
			posRandomReveal,
		)

		freeFfiSlice(blockReward)
		freeFfiSlice(burntFees)
		freeFfiSlice(posSlotRootHash)
		freeFfiSlice(posSlotGraffiti)
		freeFfiSlice(posBlockRandomness)
		freeFfiSlice(posRandomReveal)

	}

	data := generated_bindings.BlockBatchFinish(batch)

	return mamoru_sniffer.NewBlockchainData(data)
}
