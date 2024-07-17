package tests

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer/cosmos"
	"github.com/stretchr/testify/require"
	"os"
	"sync"
	"testing"
)

var once sync.Once

func BenchmarkCosmosSnifferSmoke(b *testing.B) {
	sniffer, err := testCosmosSniffer()

	require.Nil(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ctx := createCosmosCtx()

		sniffer.ObserveCosmosData(ctx)
	}
}

func TestCosmosSnifferSmoke(t *testing.T) {
	_ = os.Setenv("RUST_LOG", "info,mamoru_core=debug")

	once.Do(func() {
		mamoru_sniffer.InitLogger(func(entry mamoru_sniffer.LogEntry) {
			t.Log(entry)
		})
	})

	sniffer, err := testCosmosSniffer()

	require.Nil(t, err)
	require.NotNil(t, sniffer)

	ctx := createCosmosCtx()

	sniffer.ObserveCosmosData(ctx)
}

var cosmosSniffer *cosmos.SnifferCosmos

func testCosmosSniffer() (*cosmos.SnifferCosmos, error) {
	const storageDir = "/tmp/mamoru-storage-api"
	var err error

	if stat, _ := os.Stat(storageDir); !stat.IsDir() {
		err = os.Mkdir(storageDir, 0755)
		if err != nil {
			return nil, err
		}
		defer func() {
			os.Remove(storageDir)
		}()
	}
	_ = os.Setenv("MAMORU_STORAGE_API_DIR", storageDir)
	_ = os.Setenv("MAMORU_CHAIN_TYPE", "PROVENANCE_MAINNET")
	_ = os.Setenv("MAMORU_ENDPOINT", "http://localhost:9090")
	_ = os.Setenv("MAMORU_PRIVATE_KEY", "KEzDqLinan69Ybx5HAQ7J/mxC3eH+IZJ5inoUdTIs70=")
	_ = os.Setenv("MAMORU_CHAIN_ID", "validationchain")
	_ = os.Setenv("MAMORU_STORAGE_TRACK_LAST_EXECUTED_BLOCK", "true")
	_ = os.Setenv("MAMORU_STATISTICS_SEND_INTERVAL_SECS", "1")

	mutex.Lock()

	if cosmosSniffer == nil {
		cosmosSniffer, err = cosmos.CosmosConnect()
	}

	mutex.Unlock()

	return cosmosSniffer, err
}

func createCosmosCtx() cosmos.CosmosCtx {
	builder := cosmos.NewCosmosCtxBuilder()

	builder.SetBlock(cosmos.Block{
		Seq:                                0,
		Height:                             0,
		Hash:                               "some_hash",
		VersionBlock:                       0,
		VersionApp:                         0,
		ChainId:                            "",
		Time:                               1,
		LastBlockIdHash:                    "some_hash",
		LastBlockIdPartSetHeaderTotal:      0,
		LastBlockIdPartSetHeaderHash:       "some_hash",
		LastCommitHash:                     "some_hash",
		DataHash:                           "some_hash",
		ValidatorsHash:                     "some_hash",
		NextValidatorsHash:                 "some_hash",
		ConsensusHash:                      "some_hash",
		AppHash:                            "some_hash",
		LastResultsHash:                    "some_hash",
		EvidenceHash:                       "some_hash",
		ProposerAddress:                    "some_hash",
		LastCommitInfoRound:                0,
		ConsensusParamUpdatesBlockMaxBytes: 0,
		ConsensusParamUpdatesBlockMaxGas:   0,
		ConsensusParamUpdatesEvidenceMaxAgeNumBlocks: 0,
		ConsensusParamUpdatesEvidenceMaxAgeDuration:  0,
		ConsensusParamUpdatesEvidenceMaxBytes:        0,
		ConsensusParamUpdatesValidatorPubKeyTypes:    "",
		ConsensusParamUpdatesVersionApp:              0,
	})

	builder.AppendValidatorUpdates([]cosmos.ValidatorUpdate{
		{
			Seq:    0,
			PubKey: nil,
			Power:  0,
		},
		{
			Seq:    0,
			PubKey: nil,
			Power:  0,
		},
	})

	builder.AppendMisbehaviors([]cosmos.Misbehavior{
		{
			Seq:              0,
			BlockSeq:         0,
			Typ:              "",
			ValidatorPower:   0,
			ValidatorAddress: "some_address",
			Height:           0,
			Time:             0,
			TotalVotingPower: 0,
		},
		{
			Seq:              0,
			BlockSeq:         0,
			Typ:              "",
			ValidatorPower:   0,
			ValidatorAddress: "some_address2",
			Height:           0,
			Time:             0,
			TotalVotingPower: 0,
		},
	})

	builder.AppendVoteInfos([]cosmos.VoteInfo{
		{
			Seq:              0,
			BlockSeq:         0,
			ValidatorAddress: "",
			ValidatorPower:   0,
			SignedLastBlock:  false,
		},
		{
			Seq:              0,
			BlockSeq:         0,
			ValidatorAddress: "",
			ValidatorPower:   0,
			SignedLastBlock:  false,
		},
	})

	builder.AppendTxs([]cosmos.Transaction{
		{
			Seq:       0,
			Tx:        nil,
			TxHash:    "",
			TxIndex:   0,
			Code:      0,
			Data:      nil,
			Log:       "",
			Info:      "",
			GasWanted: 0,
			GasUsed:   0,
			Codespace: "",
		},
		{
			Seq:       0,
			Tx:        nil,
			TxHash:    "",
			TxIndex:   0,
			Code:      0,
			Data:      nil,
			Log:       "",
			Info:      "",
			GasWanted: 0,
			GasUsed:   0,
			Codespace: "",
		},
	})

	builder.AppendEvents([]cosmos.Event{
		{
			Seq:       0,
			EventType: "",
		},
		{
			Seq:       0,
			EventType: "",
		},
	})

	builder.AppendEventAttributes([]cosmos.EventAttribute{
		{
			Seq:      0,
			EventSeq: 0,
			Key:      "",
			Value:    "",
			Index:    false,
		},
		{
			Seq:      0,
			EventSeq: 0,
			Key:      "",
			Value:    "",
			Index:    false,
		},
	})
	builder.SetMempoolSource()

	builder.SetStatistics(1, 20, 30, 40)
	builder.SetBlockData("1", "0xtest_block_hash")
	builder.SetTxData("1", "0xtest_block_hash")

	return builder.Finish()
}
