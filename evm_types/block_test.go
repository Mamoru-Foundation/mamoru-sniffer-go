package evm_types

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func BenchmarkBlockDataCtxBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()

		builder.AddData(NewBlockData([]Block{
			{
				BlockIndex:                 0,
				Hash:                       "",
				ParentHash:                 "",
				StateRoot:                  "",
				Nonce:                      0,
				Status:                     "",
				Timestamp:                  0,
				BlockReward:                nil,
				FeeRecipient:               "",
				TotalDifficulty:            0,
				Size:                       0,
				GasUsed:                    0,
				GasLimit:                   0,
				BurntFees:                  nil,
				PosProposedOnTime:          0,
				PosSlot:                    0,
				PosEpoch:                   0,
				PosProposerIndex:           0,
				PosSlotRootHash:            nil,
				PosBeaconChainDepositCount: 0,
				PosSlotGraffiti:            nil,
				PosBlockRandomness:         nil,
				PosRandomReveal:            nil,
			},
		}))

		_ = builder.Finish("some-id", "some-hash", time.Now())
	}
}

func TestBlockDataCtxBuilder(t *testing.T) {
	builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()
	data := NewBlockData([]Block{
		{
			BlockIndex:                 0,
			Hash:                       "",
			ParentHash:                 "",
			StateRoot:                  "",
			Nonce:                      0,
			Status:                     "",
			Timestamp:                  0,
			BlockReward:                nil,
			FeeRecipient:               "",
			TotalDifficulty:            0,
			Size:                       0,
			GasUsed:                    0,
			GasLimit:                   0,
			BurntFees:                  nil,
			PosProposedOnTime:          0,
			PosSlot:                    0,
			PosEpoch:                   0,
			PosProposerIndex:           0,
			PosSlotRootHash:            nil,
			PosBeaconChainDepositCount: 0,
			PosSlotGraffiti:            nil,
			PosBlockRandomness:         nil,
			PosRandomReveal:            nil,
		},
	})

	assert.NotNil(t, data)

	builder.AddData(data)

	_ = builder.Finish("some-id", "some-hash", time.Now())
}
