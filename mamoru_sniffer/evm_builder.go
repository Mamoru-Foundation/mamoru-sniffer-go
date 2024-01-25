package mamoru_sniffer

import "github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"

type EvmCtxBuilder struct {
	*generated_bindings.FfiEvmBlockchainDataBuilderT
}

type EvmCtx struct {
	*generated_bindings.FfiEvmBlockchainDataCtxT
}

func NewEvmCtxBuilder() EvmCtxBuilder {
	builder := generated_bindings.NewEvmBlockchainDataBuilder()

	return EvmCtxBuilder{builder}
}

func (b EvmCtxBuilder) SetTxData(txId string, txHash string) {
	generated_bindings.EvmBlockchainDataBuilderSetTx(b.FfiEvmBlockchainDataBuilderT, txId, txHash)
}

func (b EvmCtxBuilder) SetBlockData(blockId string, blockHash string) {
	generated_bindings.EvmBlockchainDataBuilderSetBlock(b.FfiEvmBlockchainDataBuilderT, blockId, blockHash)
}

func (b EvmCtxBuilder) SetMempoolSource() {
	generated_bindings.EvmBlockchainDataBuilderSetMempoolSource(b.FfiEvmBlockchainDataBuilderT)
}

func (b EvmCtxBuilder) SetStatistics(blocks, txs, events, callTraces uint64) {
	generated_bindings.EvmBlockchainDataBuilderSetStatistics(b.FfiEvmBlockchainDataBuilderT, blocks, txs, events, callTraces)
}

func (b EvmCtxBuilder) Finish() EvmCtx {
	ctx := generated_bindings.EvmBlockchainDataBuilderFinish(b.FfiEvmBlockchainDataBuilderT)

	return EvmCtx{ctx}
}

func (b EvmCtxBuilder) SetBlock(block Block) {
	blockReward := SliceToFfi(block.BlockReward)

	generated_bindings.EvmBlockSet(
		b.FfiEvmBlockchainDataBuilderT,
		block.BlockIndex,
		block.Hash,
		block.ParentHash,
		block.StateRoot,
		block.Nonce,
		block.Status,
		block.Timestamp,
		blockReward,
		block.FeeRecipient,
		block.TotalDifficulty,
		block.Size,
		block.GasUsed,
		block.GasLimit,
	)

	FreeFfiSlice(blockReward)
}

func (b EvmCtxBuilder) AppendTxs(txs []Transaction) {
	for _, tx := range txs {
		input := SliceToFfi(tx.Input)
		generated_bindings.EvmTransactionAppend(
			b.FfiEvmBlockchainDataBuilderT,
			tx.TxIndex,
			tx.TxHash,
			tx.BlockIndex,
			tx.Type,
			tx.Nonce,
			tx.Status,
			tx.From,
			tx.To,
			tx.Value,
			tx.Fee,
			tx.GasPrice,
			tx.GasLimit,
			tx.GasUsed,
			input,
			tx.Size,
		)
		FreeFfiSlice(input)
	}
}

func (b EvmCtxBuilder) AppendEvents(events []Event) {
	for _, event := range events {
		topic0 := SliceToFfi(event.Topic0)
		topic1 := SliceToFfi(event.Topic1)
		topic2 := SliceToFfi(event.Topic2)
		topic3 := SliceToFfi(event.Topic3)
		topic4 := SliceToFfi(event.Topic4)
		data := SliceToFfi(event.Data)
		generated_bindings.EvmEventAppend(
			b.FfiEvmBlockchainDataBuilderT,
			event.Index,
			event.Address,
			event.BlockNumber,
			event.TxHash,
			event.TxIndex,
			event.BlockHash,
			topic0,
			topic1,
			topic2,
			topic3,
			topic4,
			data,
		)
		FreeFfiSlice(topic0)
		FreeFfiSlice(topic1)
		FreeFfiSlice(topic2)
		FreeFfiSlice(topic3)
		FreeFfiSlice(topic4)
		FreeFfiSlice(data)
	}
}

func (b EvmCtxBuilder) AppendCallTraces(callTraces []CallTrace) {
	for _, trace := range callTraces {
		input := SliceToFfi(trace.Input)
		generated_bindings.EvmCallTraceAppend(
			b.FfiEvmBlockchainDataBuilderT,
			trace.Seq,
			trace.Depth,
			trace.TxIndex,
			trace.BlockIndex,
			trace.Type,
			trace.From,
			trace.To,
			trace.Value,
			trace.GasLimit,
			trace.GasUsed,
			input,
		)
		FreeFfiSlice(input)
	}
}
