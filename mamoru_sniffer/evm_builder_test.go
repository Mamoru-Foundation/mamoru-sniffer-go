package mamoru_sniffer

import (
	"testing"
)

func BenchmarkEvmCtxBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = createEvmCtx()
	}
}

func TestEvmCtxBuilder(t *testing.T) {
	_ = createEvmCtx()
}

func createEvmCtx() EvmCtx {
	builder := NewEvmCtxBuilder()

	builder.SetBlock(Block{
		BlockIndex:      0,
		Hash:            "",
		ParentHash:      "",
		StateRoot:       "",
		Nonce:           0,
		Status:          "",
		Timestamp:       0,
		BlockReward:     nil,
		FeeRecipient:    "",
		TotalDifficulty: 0,
		Size:            0,
		GasUsed:         0,
		GasLimit:        0,
	})

	builder.AppendTxs([]Transaction{
		{
			TxIndex:    0,
			TxHash:     "",
			Type:       0,
			Nonce:      0,
			Status:     0,
			BlockIndex: 0,
			From:       "",
			To:         "",
			Value:      0,
			Fee:        0,
			GasPrice:   0,
			GasLimit:   0,
			GasUsed:    0,
			Size:       123.1123,
			Input:      []byte{1, 2, 3, 4},
		},
		{
			TxIndex:    0,
			TxHash:     "hash",
			Type:       1,
			Nonce:      29876,
			Status:     0,
			BlockIndex: 0,
			From:       "hash",
			To:         "to_hash",
			Value:      12345,
			Fee:        123,
			GasPrice:   0,
			GasLimit:   0,
			GasUsed:    0,
			Size:       123.1123,
			Input:      []byte{1, 2, 3, 4},
		},
	})

	builder.AppendEvents([]Event{
		{
			Index:       0,
			TxIndex:     0,
			Address:     "",
			BlockNumber: 0,
			TxHash:      "",
			BlockHash:   "",
			Topic0:      nil,
			Topic1:      nil,
			Topic2:      nil,
			Topic3:      nil,
			Topic4:      nil,
			Data:        nil,
		},
		{
			Index:       0,
			TxIndex:     0,
			Address:     "",
			BlockNumber: 0,
			TxHash:      "",
			BlockHash:   "",
			Topic0:      []byte{1, 2, 3, 4, 5},
			Topic1:      []byte{1, 2, 3, 4, 5},
			Topic2:      []byte{1, 2, 3, 4, 5},
			Topic3:      []byte{1, 2, 3, 4, 5},
			Topic4:      []byte{1, 2, 3, 4, 5},
			Data:        []byte{1, 2, 3, 4, 5},
		},
	})

	builder.AppendCallTraces([]CallTrace{
		{
			Depth:      1,
			TxIndex:    0,
			BlockIndex: 0,
			Type:       "",
			From:       "",
			To:         "",
			Value:      0,
			GasLimit:   0,
			GasUsed:    0,
			Input:      []byte{1, 2, 3, 4, 5},
		},
		{
			Depth:      2,
			TxIndex:    1,
			BlockIndex: 2,
			Type:       "string",
			From:       "string",
			To:         "string",
			Value:      123,
			GasLimit:   123,
			GasUsed:    321,
			Input:      []byte{1, 2, 3, 4, 5},
		},
	})

	return builder.Finish("some-id", "some-hash")
}
