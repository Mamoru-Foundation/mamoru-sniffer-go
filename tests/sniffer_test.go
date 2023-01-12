package tests

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/evm_types"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/stretchr/testify/require"
	"os"
	"sync"
	"testing"
	"time"
)

func BenchmarkSnifferSmoke(b *testing.B) {
	sniffer, err := testSniffer()

	require.Nil(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		builder := mamoru_sniffer.NewBlockchainDataCtxBuilder()

		builder.AddData(evm_types.NewTransactionData([]evm_types.Transaction{
			{
				Index:      1,
				BlockIndex: 2,
				Hash:       []byte{3, 4, 5},
				Type:       6,
				Nonce:      7,
				Status:     "completed",
				Timestamp:  42,
				From:       []byte{7, 8, 9},
				To:         []byte{7, 8, 9},
			},
			{
				Index:      1,
				BlockIndex: 2,
				Hash:       []byte{3, 4, 5},
				Type:       6,
				Nonce:      7,
				Status:     "completed",
				Timestamp:  42,
				From:       []byte{7, 8, 9},
				To:         []byte{7, 8, 9},
			},
		}))

		ctx := builder.Finish("one", "two", time.Now())

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
			Index:      1,
			BlockIndex: 2,
			Hash:       []byte{3, 4, 5},
			Type:       6,
			Nonce:      7,
			Status:     "completed",
			Timestamp:  42,
			From:       []byte{7, 8, 9},
			To:         []byte{7, 8, 9},
		},
		{
			Index:      1,
			BlockIndex: 2,
			Hash:       []byte{3, 4, 5},
			Type:       6,
			Nonce:      7,
			Status:     "completed",
			Timestamp:  42,
			From:       []byte{7, 8, 9},
			To:         []byte{7, 8, 9},
		},
	}))

	ctx := builder.Finish("one", "two", time.Now())

	sniffer.ObserveData(ctx)
}

var mutex sync.Mutex
var sniffer *mamoru_sniffer.Sniffer

func testSniffer() (*mamoru_sniffer.Sniffer, error) {
	_ = os.Setenv("MAMORU_CHAIN_TYPE", "SUI_DEVNET")
	_ = os.Setenv("MAMORU_ENDPOINT", "http://localhost:9090")
	_ = os.Setenv("MAMORU_PRIVATE_KEY", "7EiTMlvnm0zUIoozWIkx00+H+2U9MhRxwooF/yN/ofc=")
	_ = os.Setenv("MAMORU_CHAIN_ID", "validationchain")

	mutex.Lock()

	var err error
	if sniffer == nil {
		sniffer, err = mamoru_sniffer.Connect()
	}

	mutex.Unlock()

	return sniffer, err
}
