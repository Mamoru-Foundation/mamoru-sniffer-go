package tests

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/stretchr/testify/require"
	"os"
	"sync"
	"testing"
)

func BenchmarkSnifferSmoke(b *testing.B) {
	sniffer, err := testSniffer()

	require.Nil(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ctx := createEvmCtx()

		sniffer.ObserveEvmData(ctx)
	}
}

func TestSnifferSmoke(t *testing.T) {
	sniffer, err := testSniffer()

	require.Nil(t, err)
	require.NotNil(t, sniffer)

	ctx := createEvmCtx()

	sniffer.ObserveEvmData(ctx)
}

var mutex sync.Mutex
var sniffer *mamoru_sniffer.Sniffer

func testSniffer() (*mamoru_sniffer.Sniffer, error) {
	_ = os.Setenv("MAMORU_CHAIN_TYPE", "ETH_TESTNET")
	_ = os.Setenv("MAMORU_ENDPOINT", "http://localhost:9090")
	_ = os.Setenv("MAMORU_PRIVATE_KEY", "4gHo0h3hL4MOhC4e+fCPSREUsy+DLOKsBgo4MxAyMbU=")
	_ = os.Setenv("MAMORU_CHAIN_ID", "validationchain")
	_ = os.Setenv("MAMORU_STATISTICS_SEND_INTERVAL_SECS", "1")

	mutex.Lock()

	var err error
	if sniffer == nil {
		sniffer, err = mamoru_sniffer.Connect()
	}

	mutex.Unlock()

	return sniffer, err
}

func createEvmCtx() mamoru_sniffer.EvmCtx {
	builder := mamoru_sniffer.NewEvmCtxBuilder()

	builder.SetBlock(mamoru_sniffer.Block{
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

	builder.AppendTxs([]mamoru_sniffer.Transaction{
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

	builder.AppendEvents([]mamoru_sniffer.Event{
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

	builder.AppendCallTraces([]mamoru_sniffer.CallTrace{
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
	builder.SetMempoolSource()

	builder.SetStatistics(1, 20, 30, 40)
	builder.SetBlockData("1", "0xtest_block_hash")
	builder.SetTxData("1", "0xtest_block_hash")

	return builder.Finish()
}
