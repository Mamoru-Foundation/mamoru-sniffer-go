package tests

import (
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
		ctx := createEvmCtx(true)

		sniffer.ObserveEvmData(ctx)
	}
}

func TestSnifferSmoke(t *testing.T) {
	_ = os.Setenv("RUST_LOG", "info,mamoru_core=debug")
	once.Do(func() {
		mamoru_sniffer.InitLogger(func(entry mamoru_sniffer.LogEntry) {
			t.Log(entry)
		})
	})

	sniffer, err := testSniffer()

	require.Nil(t, err)
	require.NotNil(t, sniffer)

	ctx := createEvmCtx(true)

	sniffer.ObserveEvmData(ctx)

	ctxBlock := createEvmCtx(false)

	sniffer.ObserveEvmData(ctxBlock)

	time.Sleep(5 * time.Second)
}

var mutex sync.Mutex
var sniffer *mamoru_sniffer.Sniffer

func testSniffer() (*mamoru_sniffer.Sniffer, error) {
	_ = os.Setenv("MAMORU_CHAIN_TYPE", "ETH_TESTNET")
	_ = os.Setenv("MAMORU_ENDPOINT", "http://localhost:9090")
	_ = os.Setenv("MAMORU_PRIVATE_KEY", "KEzDqLinan69Ybx5HAQ7J/mxC3eH+IZJ5inoUdTIs70=")
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

func createEvmCtx(mempool bool) mamoru_sniffer.EvmCtx {
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
	if mempool {
		builder.SetMempoolSource()
	}

	builder.SetStatistics(1, 20, 30, 40)
	builder.SetBlockData("1", "0xtest_block_hash")
	builder.SetTxData("1", "0xtest_block_hash")

	return builder.Finish()
}

// Public this agent for testing Incidents report to MamoruCain(Validation Chain) ETH_TESTNET
// must send 2 incidents for mempool and block
/*
import { EvmCtx } from "@mamoru-ai/mamoru-evm-sdk-as/assembly"
import {IncidentSeverity, report} from "@mamoru-ai/mamoru-sdk-as/assembly"

export function main(): void {
  const ctx = EvmCtx.load();
  const block = ctx.block;
  const txs = ctx.txs;
  const events = ctx.events;
  const calltraces = ctx.callTraces;

  if (block == null) {
    report(
        "dummy-debug-evm",
        IncidentSeverity.Info,
        `Tx len: ${txs.length}; Env len: ${events.length}; Calltrace len: ${calltraces.length}`,
    );
    return;
  }
  report(
      "dummy-debug-evm",
      IncidentSeverity.Info,
      `Block number: ${block.blockIndex}; Block time: ${block.timestamp}`,
  );
}
*/
