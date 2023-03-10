package tests

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/evm_types"
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
		builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()

		builder.AddData(evm_types.NewTransactionData([]evm_types.Transaction{
			{
				TxIndex:    2,
				TxHash:     "hash",
				Type:       0,
				Nonce:      2123,
				Status:     1,
				BlockIndex: 2,
				From:       "from",
				To:         "to",
				Value:      123123,
				Fee:        122,
				GasPrice:   12,
				GasLimit:   123,
				GasUsed:    123,
			},
		}))

		ctx := builder.Finish("one", "two")

		sniffer.ObserveData(ctx)
	}
}

func TestSnifferSmoke(t *testing.T) {
	sniffer, err := testSniffer()

	require.Nil(t, err)
	require.NotNil(t, sniffer)

	builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()

	builder.AddData(evm_types.NewTransactionData([]evm_types.Transaction{
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
		},
	}))

	ctx := builder.Finish("one", "two")

	sniffer.ObserveData(ctx)
}

var mutex sync.Mutex
var sniffer *mamoru_sniffer.Sniffer

func testSniffer() (*mamoru_sniffer.Sniffer, error) {
	_ = os.Setenv("MAMORU_CHAIN_TYPE", "SUI_DEVNET")
	_ = os.Setenv("MAMORU_ENDPOINT", "http://localhost:9090")
	_ = os.Setenv("MAMORU_PRIVATE_KEY", "XWDJ6pY2rSSoN0QvgXsZUTjHkww063IEV5K/ihblVDw=")
	_ = os.Setenv("MAMORU_CHAIN_ID", "validationchain")

	mutex.Lock()

	var err error
	if sniffer == nil {
		sniffer, err = mamoru_sniffer.Connect()
	}

	mutex.Unlock()

	return sniffer, err
}
