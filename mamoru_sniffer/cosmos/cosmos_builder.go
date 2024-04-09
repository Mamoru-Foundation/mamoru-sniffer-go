package cosmos

import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
)

type CosmosCtxBuilder struct {
	*generated_bindings.FfiCosmosBlockchainDataBuilderT
}

type CosmosCtx struct {
	*generated_bindings.FfiCosmosBlockchainDataCtxT
}

func NewCosmosCtxBuilder() CosmosCtxBuilder {
	builder := generated_bindings.NewCosmosBlockchainDataBuilder()

	return CosmosCtxBuilder{builder}
}

func (b CosmosCtxBuilder) SetTxData(txId string, txHash string) {
	generated_bindings.CosmosBlockchainDataBuilderSetTx(b.FfiCosmosBlockchainDataBuilderT, txId, txHash)
}

func (b CosmosCtxBuilder) SetBlockData(blockId string, blockHash string) {
	generated_bindings.CosmosBlockchainDataBuilderSetBlock(b.FfiCosmosBlockchainDataBuilderT, blockId, blockHash)
}

func (b CosmosCtxBuilder) SetMempoolSource() {
	generated_bindings.CosmosBlockchainDataBuilderSetMempoolSource(b.FfiCosmosBlockchainDataBuilderT)
}

func (b CosmosCtxBuilder) SetStatistics(blocks, txs, events, callTraces uint64) {
	generated_bindings.CosmosBlockchainDataBuilderSetStatistics(b.FfiCosmosBlockchainDataBuilderT, blocks, txs, events, callTraces)
}

func (b CosmosCtxBuilder) Finish() CosmosCtx {
	ctx := generated_bindings.CosmosBlockchainDataBuilderFinish(b.FfiCosmosBlockchainDataBuilderT)

	return CosmosCtx{ctx}
}

func (b CosmosCtxBuilder) SetBlock(block Block) {
	generated_bindings.CosmosBlockSet(
		b.FfiCosmosBlockchainDataBuilderT,
		block.Seq,
		block.Height,
		block.Hash,
		block.VersionBlock,
		block.VersionApp,
		block.ChainId,
		block.Time,
		block.LastBlockIdHash,
		block.LastBlockIdPartSetHeaderTotal,
		block.LastBlockIdPartSetHeaderHash,
		block.LastCommitHash,
		block.DataHash,
		block.ValidatorsHash,
		block.NextValidatorsHash,
		block.ConsensusHash,
		block.AppHash,
		block.LastResultsHash,
		block.EvidenceHash,
		block.ProposerAddress,
		block.LastCommitInfoRound,
		block.ConsensusParamUpdatesBlockMaxBytes,
		block.ConsensusParamUpdatesBlockMaxGas,
		block.ConsensusParamUpdatesEvidenceMaxAgeNumBlocks,
		block.ConsensusParamUpdatesEvidenceMaxAgeDuration,
		block.ConsensusParamUpdatesEvidenceMaxBytes,
		block.ConsensusParamUpdatesValidatorPubKeyTypes,
		block.ConsensusParamUpdatesVersionApp,
	)
}

func (b CosmosCtxBuilder) AppendTxs(txs []Transaction) {
	for _, tx := range txs {
		data := mamoru_sniffer.SliceToFfi(tx.Data)
		txData := mamoru_sniffer.SliceToFfi(tx.Tx)
		generated_bindings.CosmosTransactionAppend(
			b.FfiCosmosBlockchainDataBuilderT,
			txData,
			tx.TxHash,
			tx.TxIndex,
			tx.Code,
			data,
			tx.Log,
			tx.Info,
			tx.GasUsed,
			tx.GasWanted,
			tx.Codespace,
		)
		mamoru_sniffer.FreeFfiSlice(data)
		mamoru_sniffer.FreeFfiSlice(txData)
	}
}

func (b CosmosCtxBuilder) AppendEvents(events []Event) {
	for _, event := range events {
		generated_bindings.CosmosEventAppend(
			b.FfiCosmosBlockchainDataBuilderT,
			event.Seq,
			event.EventType,
		)
	}
}

func (b CosmosCtxBuilder) AppendEventAttributes(eventAttributes []EventAttribute) {
	for _, ea := range eventAttributes {
		generated_bindings.CosmosEventAttributeAppend(
			b.FfiCosmosBlockchainDataBuilderT,
			ea.Seq,
			ea.EventSeq,
			ea.Key,
			ea.Value,
			ea.Index,
		)
	}
}

func (b CosmosCtxBuilder) AppendValidatorUpdates(validatorUpdates []ValidatorUpdate) {
	for _, vu := range validatorUpdates {
		pubKey := mamoru_sniffer.SliceToFfi(vu.PubKey)

		generated_bindings.CosmosValidatorUpdateAppend(
			b.FfiCosmosBlockchainDataBuilderT,
			vu.Seq,
			pubKey,
			vu.Power,
		)
		mamoru_sniffer.FreeFfiSlice(pubKey)
	}
}

func (b CosmosCtxBuilder) AppendMisbehaviors(misbehaviors []Misbehavior) {
	for _, m := range misbehaviors {
		generated_bindings.CosmosMisbehaviorAppend(
			b.FfiCosmosBlockchainDataBuilderT,
			m.BlockSeq,
			m.Typ,
			m.ValidatorPower,
			m.ValidatorAddress,
			m.Height,
			m.Time,
			m.TotalVotingPower,
		)
	}
}

func (b CosmosCtxBuilder) AppendVoteInfos(voteInfos []VoteInfo) {
	for _, vi := range voteInfos {
		generated_bindings.CosmosVoteInfosAppend(
			b.FfiCosmosBlockchainDataBuilderT,
			vi.BlockSeq,
			vi.ValidatorAddress,
			vi.ValidatorPower,
			vi.SignedLastBlock,
		)
	}
}

func (b CosmosCtxBuilder) AppendEvmCallTraces(evmCallTraces []EvmCallTrace) {
	for _, ect := range evmCallTraces {
		input := mamoru_sniffer.SliceToFfi(ect.Input)
		generated_bindings.CosmosEvmCallTraceAppend(
			b.FfiCosmosBlockchainDataBuilderT,
			ect.TxHash,
			ect.Depth,
			ect.TxIndex,
			ect.BlockIndex,
			ect.Type,
			ect.From,
			ect.To,
			ect.Value,
			ect.GasLimit,
			ect.GasUsed,
			input,
			ect.Output,
			ect.Error,
			ect.RevertReason,
		)
		mamoru_sniffer.FreeFfiSlice(input)
	}
}
