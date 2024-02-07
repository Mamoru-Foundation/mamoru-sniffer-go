// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 07 Feb 2024 14:03:48 WET.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package generated_bindings

/*
#cgo CFLAGS: -I${SRCDIR}/../packaged/include
#cgo LDFLAGS: -lmamoru_sniffer_go
#cgo darwin,arm64 LDFLAGS: -framework Security -L${SRCDIR}/../packaged/lib/darwin-arm64
#cgo linux,amd64 LDFLAGS: -Wl,--no-as-needed -ldl -lm -L${SRCDIR}/../packaged/lib/linux-amd64
#include <libmamoru_sniffer_go.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// InitLogger function as declared in include/libmamoru_sniffer_go.h:22
func InitLogger() {
	C.init_logger()
}

// NewEvmBlockchainDataBuilder function as declared in include/libmamoru_sniffer_go.h:26
func NewEvmBlockchainDataBuilder() *FfiEvmBlockchainDataBuilderT {
	__ret := C.new_evm_blockchain_data_builder()
	__v := *(**FfiEvmBlockchainDataBuilderT)(unsafe.Pointer(&__ret))
	return __v
}

// EvmBlockchainDataBuilderSetTx function as declared in include/libmamoru_sniffer_go.h:28
func EvmBlockchainDataBuilderSetTx(Builder *FfiEvmBlockchainDataBuilderT, TxId string, TxHash string) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	TxId = safeString(TxId)
	cTxId, cTxIdAllocMap := unpackPCharString(TxId)
	TxHash = safeString(TxHash)
	cTxHash, cTxHashAllocMap := unpackPCharString(TxHash)
	C.evm_blockchain_data_builder_set_tx(cBuilder, cTxId, cTxHash)
	runtime.KeepAlive(TxHash)
	runtime.KeepAlive(cTxHashAllocMap)
	runtime.KeepAlive(TxId)
	runtime.KeepAlive(cTxIdAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// EvmBlockchainDataBuilderSetBlock function as declared in include/libmamoru_sniffer_go.h:33
func EvmBlockchainDataBuilderSetBlock(Builder *FfiEvmBlockchainDataBuilderT, BlockId string, BlockHash string) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	BlockId = safeString(BlockId)
	cBlockId, cBlockIdAllocMap := unpackPCharString(BlockId)
	BlockHash = safeString(BlockHash)
	cBlockHash, cBlockHashAllocMap := unpackPCharString(BlockHash)
	C.evm_blockchain_data_builder_set_block(cBuilder, cBlockId, cBlockHash)
	runtime.KeepAlive(BlockHash)
	runtime.KeepAlive(cBlockHashAllocMap)
	runtime.KeepAlive(BlockId)
	runtime.KeepAlive(cBlockIdAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// EvmBlockchainDataBuilderSetMempoolSource function as declared in include/libmamoru_sniffer_go.h:38
func EvmBlockchainDataBuilderSetMempoolSource(Builder *FfiEvmBlockchainDataBuilderT) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	C.evm_blockchain_data_builder_set_mempool_source(cBuilder)
	runtime.KeepAlive(cBuilderAllocMap)
}

// EvmBlockchainDataBuilderSetStatistics function as declared in include/libmamoru_sniffer_go.h:45
func EvmBlockchainDataBuilderSetStatistics(Builder *FfiEvmBlockchainDataBuilderT, Blocks uint64, Transactions uint64, Events uint64, CallTraces uint64) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cBlocks, cBlocksAllocMap := (C.uint64_t)(Blocks), cgoAllocsUnknown
	cTransactions, cTransactionsAllocMap := (C.uint64_t)(Transactions), cgoAllocsUnknown
	cEvents, cEventsAllocMap := (C.uint64_t)(Events), cgoAllocsUnknown
	cCallTraces, cCallTracesAllocMap := (C.uint64_t)(CallTraces), cgoAllocsUnknown
	C.evm_blockchain_data_builder_set_statistics(cBuilder, cBlocks, cTransactions, cEvents, cCallTraces)
	runtime.KeepAlive(cCallTracesAllocMap)
	runtime.KeepAlive(cEventsAllocMap)
	runtime.KeepAlive(cTransactionsAllocMap)
	runtime.KeepAlive(cBlocksAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// EvmBlockchainDataBuilderFinish function as declared in include/libmamoru_sniffer_go.h:57
func EvmBlockchainDataBuilderFinish(Builder *FfiEvmBlockchainDataBuilderT) *FfiEvmBlockchainDataCtxT {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	__ret := C.evm_blockchain_data_builder_finish(cBuilder)
	runtime.KeepAlive(cBuilderAllocMap)
	__v := *(**FfiEvmBlockchainDataCtxT)(unsafe.Pointer(&__ret))
	return __v
}

// EvmSnifferObserveData function as declared in include/libmamoru_sniffer_go.h:65
func EvmSnifferObserveData(Sniffer *FfiSnifferT, Data *FfiEvmBlockchainDataCtxT) {
	cSniffer, cSnifferAllocMap := (*C.FfiSniffer_t)(unsafe.Pointer(Sniffer)), cgoAllocsUnknown
	cData, cDataAllocMap := (*C.FfiEvmBlockchainDataCtx_t)(unsafe.Pointer(Data)), cgoAllocsUnknown
	C.evm_sniffer_observe_data(cSniffer, cData)
	runtime.KeepAlive(cDataAllocMap)
	runtime.KeepAlive(cSnifferAllocMap)
}

// EvmTransactionAppend function as declared in include/libmamoru_sniffer_go.h:96
func EvmTransactionAppend(Builder *FfiEvmBlockchainDataBuilderT, TxIndex uint32, TxHash string, BlockIndex uint64, Typ byte, Nonce uint64, Status uint64, From string, To string, Value uint64, Fee uint64, GasPrice uint64, GasLimit uint64, GasUsed uint64, Input SliceRefUint8T, Size float64) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cTxIndex, cTxIndexAllocMap := (C.uint32_t)(TxIndex), cgoAllocsUnknown
	TxHash = safeString(TxHash)
	cTxHash, cTxHashAllocMap := unpackPCharString(TxHash)
	cBlockIndex, cBlockIndexAllocMap := (C.uint64_t)(BlockIndex), cgoAllocsUnknown
	cTyp, cTypAllocMap := (C.uint8_t)(Typ), cgoAllocsUnknown
	cNonce, cNonceAllocMap := (C.uint64_t)(Nonce), cgoAllocsUnknown
	cStatus, cStatusAllocMap := (C.uint64_t)(Status), cgoAllocsUnknown
	From = safeString(From)
	cFrom, cFromAllocMap := unpackPCharString(From)
	To = safeString(To)
	cTo, cToAllocMap := unpackPCharString(To)
	cValue, cValueAllocMap := (C.uint64_t)(Value), cgoAllocsUnknown
	cFee, cFeeAllocMap := (C.uint64_t)(Fee), cgoAllocsUnknown
	cGasPrice, cGasPriceAllocMap := (C.uint64_t)(GasPrice), cgoAllocsUnknown
	cGasLimit, cGasLimitAllocMap := (C.uint64_t)(GasLimit), cgoAllocsUnknown
	cGasUsed, cGasUsedAllocMap := (C.uint64_t)(GasUsed), cgoAllocsUnknown
	cInput, cInputAllocMap := Input.PassValue()
	cSize, cSizeAllocMap := (C.double)(Size), cgoAllocsUnknown
	C.evm_transaction_append(cBuilder, cTxIndex, cTxHash, cBlockIndex, cTyp, cNonce, cStatus, cFrom, cTo, cValue, cFee, cGasPrice, cGasLimit, cGasUsed, cInput, cSize)
	runtime.KeepAlive(cSizeAllocMap)
	runtime.KeepAlive(cInputAllocMap)
	runtime.KeepAlive(cGasUsedAllocMap)
	runtime.KeepAlive(cGasLimitAllocMap)
	runtime.KeepAlive(cGasPriceAllocMap)
	runtime.KeepAlive(cFeeAllocMap)
	runtime.KeepAlive(cValueAllocMap)
	runtime.KeepAlive(To)
	runtime.KeepAlive(cToAllocMap)
	runtime.KeepAlive(From)
	runtime.KeepAlive(cFromAllocMap)
	runtime.KeepAlive(cStatusAllocMap)
	runtime.KeepAlive(cNonceAllocMap)
	runtime.KeepAlive(cTypAllocMap)
	runtime.KeepAlive(cBlockIndexAllocMap)
	runtime.KeepAlive(TxHash)
	runtime.KeepAlive(cTxHashAllocMap)
	runtime.KeepAlive(cTxIndexAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// EvmCallTraceAppend function as declared in include/libmamoru_sniffer_go.h:114
func EvmCallTraceAppend(Builder *FfiEvmBlockchainDataBuilderT, Seq uint32, Depth uint32, TxIndex uint32, BlockIndex uint64, Typ string, From string, To string, Value uint64, GasLimit uint64, GasUsed uint64, Input SliceRefUint8T) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cSeq, cSeqAllocMap := (C.uint32_t)(Seq), cgoAllocsUnknown
	cDepth, cDepthAllocMap := (C.uint32_t)(Depth), cgoAllocsUnknown
	cTxIndex, cTxIndexAllocMap := (C.uint32_t)(TxIndex), cgoAllocsUnknown
	cBlockIndex, cBlockIndexAllocMap := (C.uint64_t)(BlockIndex), cgoAllocsUnknown
	Typ = safeString(Typ)
	cTyp, cTypAllocMap := unpackPCharString(Typ)
	From = safeString(From)
	cFrom, cFromAllocMap := unpackPCharString(From)
	To = safeString(To)
	cTo, cToAllocMap := unpackPCharString(To)
	cValue, cValueAllocMap := (C.uint64_t)(Value), cgoAllocsUnknown
	cGasLimit, cGasLimitAllocMap := (C.uint64_t)(GasLimit), cgoAllocsUnknown
	cGasUsed, cGasUsedAllocMap := (C.uint64_t)(GasUsed), cgoAllocsUnknown
	cInput, cInputAllocMap := Input.PassValue()
	C.evm_call_trace_append(cBuilder, cSeq, cDepth, cTxIndex, cBlockIndex, cTyp, cFrom, cTo, cValue, cGasLimit, cGasUsed, cInput)
	runtime.KeepAlive(cInputAllocMap)
	runtime.KeepAlive(cGasUsedAllocMap)
	runtime.KeepAlive(cGasLimitAllocMap)
	runtime.KeepAlive(cValueAllocMap)
	runtime.KeepAlive(To)
	runtime.KeepAlive(cToAllocMap)
	runtime.KeepAlive(From)
	runtime.KeepAlive(cFromAllocMap)
	runtime.KeepAlive(Typ)
	runtime.KeepAlive(cTypAllocMap)
	runtime.KeepAlive(cBlockIndexAllocMap)
	runtime.KeepAlive(cTxIndexAllocMap)
	runtime.KeepAlive(cDepthAllocMap)
	runtime.KeepAlive(cSeqAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// EvmBlockSet function as declared in include/libmamoru_sniffer_go.h:128
func EvmBlockSet(Builder *FfiEvmBlockchainDataBuilderT, BlockIndex uint64, Hash string, ParentHash string, StateRoot string, Nonce uint64, Status string, Timestamp uint64, BlockReward SliceRefUint8T, FeeRecipient string, TotalDifficulty uint64, Size float64, GasUsed uint64, GasLimit uint64) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cBlockIndex, cBlockIndexAllocMap := (C.uint64_t)(BlockIndex), cgoAllocsUnknown
	Hash = safeString(Hash)
	cHash, cHashAllocMap := unpackPCharString(Hash)
	ParentHash = safeString(ParentHash)
	cParentHash, cParentHashAllocMap := unpackPCharString(ParentHash)
	StateRoot = safeString(StateRoot)
	cStateRoot, cStateRootAllocMap := unpackPCharString(StateRoot)
	cNonce, cNonceAllocMap := (C.uint64_t)(Nonce), cgoAllocsUnknown
	Status = safeString(Status)
	cStatus, cStatusAllocMap := unpackPCharString(Status)
	cTimestamp, cTimestampAllocMap := (C.uint64_t)(Timestamp), cgoAllocsUnknown
	cBlockReward, cBlockRewardAllocMap := BlockReward.PassValue()
	FeeRecipient = safeString(FeeRecipient)
	cFeeRecipient, cFeeRecipientAllocMap := unpackPCharString(FeeRecipient)
	cTotalDifficulty, cTotalDifficultyAllocMap := (C.uint64_t)(TotalDifficulty), cgoAllocsUnknown
	cSize, cSizeAllocMap := (C.double)(Size), cgoAllocsUnknown
	cGasUsed, cGasUsedAllocMap := (C.uint64_t)(GasUsed), cgoAllocsUnknown
	cGasLimit, cGasLimitAllocMap := (C.uint64_t)(GasLimit), cgoAllocsUnknown
	C.evm_block_set(cBuilder, cBlockIndex, cHash, cParentHash, cStateRoot, cNonce, cStatus, cTimestamp, cBlockReward, cFeeRecipient, cTotalDifficulty, cSize, cGasUsed, cGasLimit)
	runtime.KeepAlive(cGasLimitAllocMap)
	runtime.KeepAlive(cGasUsedAllocMap)
	runtime.KeepAlive(cSizeAllocMap)
	runtime.KeepAlive(cTotalDifficultyAllocMap)
	runtime.KeepAlive(FeeRecipient)
	runtime.KeepAlive(cFeeRecipientAllocMap)
	runtime.KeepAlive(cBlockRewardAllocMap)
	runtime.KeepAlive(cTimestampAllocMap)
	runtime.KeepAlive(Status)
	runtime.KeepAlive(cStatusAllocMap)
	runtime.KeepAlive(cNonceAllocMap)
	runtime.KeepAlive(StateRoot)
	runtime.KeepAlive(cStateRootAllocMap)
	runtime.KeepAlive(ParentHash)
	runtime.KeepAlive(cParentHashAllocMap)
	runtime.KeepAlive(Hash)
	runtime.KeepAlive(cHashAllocMap)
	runtime.KeepAlive(cBlockIndexAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// EvmEventAppend function as declared in include/libmamoru_sniffer_go.h:144
func EvmEventAppend(Builder *FfiEvmBlockchainDataBuilderT, Index uint32, Address string, BlockNumber uint64, TxHash string, TxIndex uint32, BlockHash string, Topic0 SliceRefUint8T, Topic1 SliceRefUint8T, Topic2 SliceRefUint8T, Topic3 SliceRefUint8T, Topic4 SliceRefUint8T, Data SliceRefUint8T) {
	cBuilder, cBuilderAllocMap := (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cIndex, cIndexAllocMap := (C.uint32_t)(Index), cgoAllocsUnknown
	Address = safeString(Address)
	cAddress, cAddressAllocMap := unpackPCharString(Address)
	cBlockNumber, cBlockNumberAllocMap := (C.uint64_t)(BlockNumber), cgoAllocsUnknown
	TxHash = safeString(TxHash)
	cTxHash, cTxHashAllocMap := unpackPCharString(TxHash)
	cTxIndex, cTxIndexAllocMap := (C.uint32_t)(TxIndex), cgoAllocsUnknown
	BlockHash = safeString(BlockHash)
	cBlockHash, cBlockHashAllocMap := unpackPCharString(BlockHash)
	cTopic0, cTopic0AllocMap := Topic0.PassValue()
	cTopic1, cTopic1AllocMap := Topic1.PassValue()
	cTopic2, cTopic2AllocMap := Topic2.PassValue()
	cTopic3, cTopic3AllocMap := Topic3.PassValue()
	cTopic4, cTopic4AllocMap := Topic4.PassValue()
	cData, cDataAllocMap := Data.PassValue()
	C.evm_event_append(cBuilder, cIndex, cAddress, cBlockNumber, cTxHash, cTxIndex, cBlockHash, cTopic0, cTopic1, cTopic2, cTopic3, cTopic4, cData)
	runtime.KeepAlive(cDataAllocMap)
	runtime.KeepAlive(cTopic4AllocMap)
	runtime.KeepAlive(cTopic3AllocMap)
	runtime.KeepAlive(cTopic2AllocMap)
	runtime.KeepAlive(cTopic1AllocMap)
	runtime.KeepAlive(cTopic0AllocMap)
	runtime.KeepAlive(BlockHash)
	runtime.KeepAlive(cBlockHashAllocMap)
	runtime.KeepAlive(cTxIndexAllocMap)
	runtime.KeepAlive(TxHash)
	runtime.KeepAlive(cTxHashAllocMap)
	runtime.KeepAlive(cBlockNumberAllocMap)
	runtime.KeepAlive(Address)
	runtime.KeepAlive(cAddressAllocMap)
	runtime.KeepAlive(cIndexAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// NewCosmosBlockchainDataBuilder function as declared in include/libmamoru_sniffer_go.h:161
func NewCosmosBlockchainDataBuilder() *FfiCosmosBlockchainDataBuilderT {
	__ret := C.new_cosmos_blockchain_data_builder()
	__v := *(**FfiCosmosBlockchainDataBuilderT)(unsafe.Pointer(&__ret))
	return __v
}

// CosmosBlockchainDataBuilderSetTx function as declared in include/libmamoru_sniffer_go.h:163
func CosmosBlockchainDataBuilderSetTx(Builder *FfiCosmosBlockchainDataBuilderT, TxId string, TxHash string) {
	cBuilder, cBuilderAllocMap := (*C.FfiCosmosBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	TxId = safeString(TxId)
	cTxId, cTxIdAllocMap := unpackPCharString(TxId)
	TxHash = safeString(TxHash)
	cTxHash, cTxHashAllocMap := unpackPCharString(TxHash)
	C.cosmos_blockchain_data_builder_set_tx(cBuilder, cTxId, cTxHash)
	runtime.KeepAlive(TxHash)
	runtime.KeepAlive(cTxHashAllocMap)
	runtime.KeepAlive(TxId)
	runtime.KeepAlive(cTxIdAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// CosmosBlockchainDataBuilderSetBlock function as declared in include/libmamoru_sniffer_go.h:168
func CosmosBlockchainDataBuilderSetBlock(Builder *FfiCosmosBlockchainDataBuilderT, BlockId string, BlockHash string) {
	cBuilder, cBuilderAllocMap := (*C.FfiCosmosBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	BlockId = safeString(BlockId)
	cBlockId, cBlockIdAllocMap := unpackPCharString(BlockId)
	BlockHash = safeString(BlockHash)
	cBlockHash, cBlockHashAllocMap := unpackPCharString(BlockHash)
	C.cosmos_blockchain_data_builder_set_block(cBuilder, cBlockId, cBlockHash)
	runtime.KeepAlive(BlockHash)
	runtime.KeepAlive(cBlockHashAllocMap)
	runtime.KeepAlive(BlockId)
	runtime.KeepAlive(cBlockIdAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// CosmosBlockchainDataBuilderSetMempoolSource function as declared in include/libmamoru_sniffer_go.h:173
func CosmosBlockchainDataBuilderSetMempoolSource(Builder *FfiCosmosBlockchainDataBuilderT) {
	cBuilder, cBuilderAllocMap := (*C.FfiCosmosBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	C.cosmos_blockchain_data_builder_set_mempool_source(cBuilder)
	runtime.KeepAlive(cBuilderAllocMap)
}

// CosmosBlockchainDataBuilderSetStatistics function as declared in include/libmamoru_sniffer_go.h:176
func CosmosBlockchainDataBuilderSetStatistics(Builder *FfiCosmosBlockchainDataBuilderT, Blocks uint64, Transactions uint64, Events uint64, CallTraces uint64) {
	cBuilder, cBuilderAllocMap := (*C.FfiCosmosBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cBlocks, cBlocksAllocMap := (C.uint64_t)(Blocks), cgoAllocsUnknown
	cTransactions, cTransactionsAllocMap := (C.uint64_t)(Transactions), cgoAllocsUnknown
	cEvents, cEventsAllocMap := (C.uint64_t)(Events), cgoAllocsUnknown
	cCallTraces, cCallTracesAllocMap := (C.uint64_t)(CallTraces), cgoAllocsUnknown
	C.cosmos_blockchain_data_builder_set_statistics(cBuilder, cBlocks, cTransactions, cEvents, cCallTraces)
	runtime.KeepAlive(cCallTracesAllocMap)
	runtime.KeepAlive(cEventsAllocMap)
	runtime.KeepAlive(cTransactionsAllocMap)
	runtime.KeepAlive(cBlocksAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// CosmosBlockchainDataBuilderFinish function as declared in include/libmamoru_sniffer_go.h:188
func CosmosBlockchainDataBuilderFinish(Builder *FfiCosmosBlockchainDataBuilderT) *FfiCosmosBlockchainDataCtxT {
	cBuilder, cBuilderAllocMap := (*C.FfiCosmosBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	__ret := C.cosmos_blockchain_data_builder_finish(cBuilder)
	runtime.KeepAlive(cBuilderAllocMap)
	__v := *(**FfiCosmosBlockchainDataCtxT)(unsafe.Pointer(&__ret))
	return __v
}

// CosmosSnifferObserveData function as declared in include/libmamoru_sniffer_go.h:194
func CosmosSnifferObserveData(Sniffer *FfiSnifferT, Data *FfiCosmosBlockchainDataCtxT) {
	cSniffer, cSnifferAllocMap := (*C.FfiSniffer_t)(unsafe.Pointer(Sniffer)), cgoAllocsUnknown
	cData, cDataAllocMap := (*C.FfiCosmosBlockchainDataCtx_t)(unsafe.Pointer(Data)), cgoAllocsUnknown
	C.cosmos_sniffer_observe_data(cSniffer, cData)
	runtime.KeepAlive(cDataAllocMap)
	runtime.KeepAlive(cSnifferAllocMap)
}

// CosmosTransactionAppend function as declared in include/libmamoru_sniffer_go.h:198
func CosmosTransactionAppend(Builder *FfiCosmosBlockchainDataBuilderT, TxData SliceRefUint8T, Code uint32, Data SliceRefUint8T, Log string, Info string, GasWanted int64, GasUsed int64, Codespace string) {
	cBuilder, cBuilderAllocMap := (*C.FfiCosmosBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cTxData, cTxDataAllocMap := TxData.PassValue()
	cCode, cCodeAllocMap := (C.uint32_t)(Code), cgoAllocsUnknown
	cData, cDataAllocMap := Data.PassValue()
	Log = safeString(Log)
	cLog, cLogAllocMap := unpackPCharString(Log)
	Info = safeString(Info)
	cInfo, cInfoAllocMap := unpackPCharString(Info)
	cGasWanted, cGasWantedAllocMap := (C.int64_t)(GasWanted), cgoAllocsUnknown
	cGasUsed, cGasUsedAllocMap := (C.int64_t)(GasUsed), cgoAllocsUnknown
	Codespace = safeString(Codespace)
	cCodespace, cCodespaceAllocMap := unpackPCharString(Codespace)
	C.cosmos_transaction_append(cBuilder, cTxData, cCode, cData, cLog, cInfo, cGasWanted, cGasUsed, cCodespace)
	runtime.KeepAlive(Codespace)
	runtime.KeepAlive(cCodespaceAllocMap)
	runtime.KeepAlive(cGasUsedAllocMap)
	runtime.KeepAlive(cGasWantedAllocMap)
	runtime.KeepAlive(Info)
	runtime.KeepAlive(cInfoAllocMap)
	runtime.KeepAlive(Log)
	runtime.KeepAlive(cLogAllocMap)
	runtime.KeepAlive(cDataAllocMap)
	runtime.KeepAlive(cCodeAllocMap)
	runtime.KeepAlive(cTxDataAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// CosmosBlockSet function as declared in include/libmamoru_sniffer_go.h:209
func CosmosBlockSet(Builder *FfiCosmosBlockchainDataBuilderT, Seq uint64, Height int64, Hash SliceRefUint8T, VersionBlock uint64, VersionApp uint64, ChainId string, Time int64, LastBlockIdHash SliceRefUint8T, LastBlockIdPartSetHeaderTotal uint32, LastBlockIdPartSetHeaderHash SliceRefUint8T, LastCommitHash SliceRefUint8T, DataHash SliceRefUint8T, ValidatorsHash SliceRefUint8T, NextValidatorsHash SliceRefUint8T, ConsensusHash SliceRefUint8T, AppHash SliceRefUint8T, LastResultsHash SliceRefUint8T, EvidenceHash SliceRefUint8T, ProposerAddress SliceRefUint8T, LastCommitInfoRound int32, ConsensusParamUpdatesBlockMaxBytes int64, ConsensusParamUpdatesBlockMaxGas int64, ConsensusParamUpdatesEvidenceMaxAgeNumBlocks int64, ConsensusParamUpdatesEvidenceMaxAgeDuration int64, ConsensusParamUpdatesEvidenceMaxBytes int64, ConsensusParamUpdatesValidatorPubKeyTypes string, ConsensusParamUpdatesVersionApp uint64) {
	cBuilder, cBuilderAllocMap := (*C.FfiCosmosBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cSeq, cSeqAllocMap := (C.uint64_t)(Seq), cgoAllocsUnknown
	cHeight, cHeightAllocMap := (C.int64_t)(Height), cgoAllocsUnknown
	cHash, cHashAllocMap := Hash.PassValue()
	cVersionBlock, cVersionBlockAllocMap := (C.uint64_t)(VersionBlock), cgoAllocsUnknown
	cVersionApp, cVersionAppAllocMap := (C.uint64_t)(VersionApp), cgoAllocsUnknown
	ChainId = safeString(ChainId)
	cChainId, cChainIdAllocMap := unpackPCharString(ChainId)
	cTime, cTimeAllocMap := (C.int64_t)(Time), cgoAllocsUnknown
	cLastBlockIdHash, cLastBlockIdHashAllocMap := LastBlockIdHash.PassValue()
	cLastBlockIdPartSetHeaderTotal, cLastBlockIdPartSetHeaderTotalAllocMap := (C.uint32_t)(LastBlockIdPartSetHeaderTotal), cgoAllocsUnknown
	cLastBlockIdPartSetHeaderHash, cLastBlockIdPartSetHeaderHashAllocMap := LastBlockIdPartSetHeaderHash.PassValue()
	cLastCommitHash, cLastCommitHashAllocMap := LastCommitHash.PassValue()
	cDataHash, cDataHashAllocMap := DataHash.PassValue()
	cValidatorsHash, cValidatorsHashAllocMap := ValidatorsHash.PassValue()
	cNextValidatorsHash, cNextValidatorsHashAllocMap := NextValidatorsHash.PassValue()
	cConsensusHash, cConsensusHashAllocMap := ConsensusHash.PassValue()
	cAppHash, cAppHashAllocMap := AppHash.PassValue()
	cLastResultsHash, cLastResultsHashAllocMap := LastResultsHash.PassValue()
	cEvidenceHash, cEvidenceHashAllocMap := EvidenceHash.PassValue()
	cProposerAddress, cProposerAddressAllocMap := ProposerAddress.PassValue()
	cLastCommitInfoRound, cLastCommitInfoRoundAllocMap := (C.int32_t)(LastCommitInfoRound), cgoAllocsUnknown
	cConsensusParamUpdatesBlockMaxBytes, cConsensusParamUpdatesBlockMaxBytesAllocMap := (C.int64_t)(ConsensusParamUpdatesBlockMaxBytes), cgoAllocsUnknown
	cConsensusParamUpdatesBlockMaxGas, cConsensusParamUpdatesBlockMaxGasAllocMap := (C.int64_t)(ConsensusParamUpdatesBlockMaxGas), cgoAllocsUnknown
	cConsensusParamUpdatesEvidenceMaxAgeNumBlocks, cConsensusParamUpdatesEvidenceMaxAgeNumBlocksAllocMap := (C.int64_t)(ConsensusParamUpdatesEvidenceMaxAgeNumBlocks), cgoAllocsUnknown
	cConsensusParamUpdatesEvidenceMaxAgeDuration, cConsensusParamUpdatesEvidenceMaxAgeDurationAllocMap := (C.int64_t)(ConsensusParamUpdatesEvidenceMaxAgeDuration), cgoAllocsUnknown
	cConsensusParamUpdatesEvidenceMaxBytes, cConsensusParamUpdatesEvidenceMaxBytesAllocMap := (C.int64_t)(ConsensusParamUpdatesEvidenceMaxBytes), cgoAllocsUnknown
	ConsensusParamUpdatesValidatorPubKeyTypes = safeString(ConsensusParamUpdatesValidatorPubKeyTypes)
	cConsensusParamUpdatesValidatorPubKeyTypes, cConsensusParamUpdatesValidatorPubKeyTypesAllocMap := unpackPCharString(ConsensusParamUpdatesValidatorPubKeyTypes)
	cConsensusParamUpdatesVersionApp, cConsensusParamUpdatesVersionAppAllocMap := (C.uint64_t)(ConsensusParamUpdatesVersionApp), cgoAllocsUnknown
	C.cosmos_block_set(cBuilder, cSeq, cHeight, cHash, cVersionBlock, cVersionApp, cChainId, cTime, cLastBlockIdHash, cLastBlockIdPartSetHeaderTotal, cLastBlockIdPartSetHeaderHash, cLastCommitHash, cDataHash, cValidatorsHash, cNextValidatorsHash, cConsensusHash, cAppHash, cLastResultsHash, cEvidenceHash, cProposerAddress, cLastCommitInfoRound, cConsensusParamUpdatesBlockMaxBytes, cConsensusParamUpdatesBlockMaxGas, cConsensusParamUpdatesEvidenceMaxAgeNumBlocks, cConsensusParamUpdatesEvidenceMaxAgeDuration, cConsensusParamUpdatesEvidenceMaxBytes, cConsensusParamUpdatesValidatorPubKeyTypes, cConsensusParamUpdatesVersionApp)
	runtime.KeepAlive(cConsensusParamUpdatesVersionAppAllocMap)
	runtime.KeepAlive(ConsensusParamUpdatesValidatorPubKeyTypes)
	runtime.KeepAlive(cConsensusParamUpdatesValidatorPubKeyTypesAllocMap)
	runtime.KeepAlive(cConsensusParamUpdatesEvidenceMaxBytesAllocMap)
	runtime.KeepAlive(cConsensusParamUpdatesEvidenceMaxAgeDurationAllocMap)
	runtime.KeepAlive(cConsensusParamUpdatesEvidenceMaxAgeNumBlocksAllocMap)
	runtime.KeepAlive(cConsensusParamUpdatesBlockMaxGasAllocMap)
	runtime.KeepAlive(cConsensusParamUpdatesBlockMaxBytesAllocMap)
	runtime.KeepAlive(cLastCommitInfoRoundAllocMap)
	runtime.KeepAlive(cProposerAddressAllocMap)
	runtime.KeepAlive(cEvidenceHashAllocMap)
	runtime.KeepAlive(cLastResultsHashAllocMap)
	runtime.KeepAlive(cAppHashAllocMap)
	runtime.KeepAlive(cConsensusHashAllocMap)
	runtime.KeepAlive(cNextValidatorsHashAllocMap)
	runtime.KeepAlive(cValidatorsHashAllocMap)
	runtime.KeepAlive(cDataHashAllocMap)
	runtime.KeepAlive(cLastCommitHashAllocMap)
	runtime.KeepAlive(cLastBlockIdPartSetHeaderHashAllocMap)
	runtime.KeepAlive(cLastBlockIdPartSetHeaderTotalAllocMap)
	runtime.KeepAlive(cLastBlockIdHashAllocMap)
	runtime.KeepAlive(cTimeAllocMap)
	runtime.KeepAlive(ChainId)
	runtime.KeepAlive(cChainIdAllocMap)
	runtime.KeepAlive(cVersionAppAllocMap)
	runtime.KeepAlive(cVersionBlockAllocMap)
	runtime.KeepAlive(cHashAllocMap)
	runtime.KeepAlive(cHeightAllocMap)
	runtime.KeepAlive(cSeqAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// CosmosEventAppend function as declared in include/libmamoru_sniffer_go.h:239
func CosmosEventAppend(Builder *FfiCosmosBlockchainDataBuilderT, Seq uint64, EventType string) {
	cBuilder, cBuilderAllocMap := (*C.FfiCosmosBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cSeq, cSeqAllocMap := (C.uint64_t)(Seq), cgoAllocsUnknown
	EventType = safeString(EventType)
	cEventType, cEventTypeAllocMap := unpackPCharString(EventType)
	C.cosmos_event_append(cBuilder, cSeq, cEventType)
	runtime.KeepAlive(EventType)
	runtime.KeepAlive(cEventTypeAllocMap)
	runtime.KeepAlive(cSeqAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// CosmosEventAttributeAppend function as declared in include/libmamoru_sniffer_go.h:247
func CosmosEventAttributeAppend(Builder *FfiCosmosBlockchainDataBuilderT, Seq uint64, EventSeq uint64, Key string, Value string, Index bool) {
	cBuilder, cBuilderAllocMap := (*C.FfiCosmosBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cSeq, cSeqAllocMap := (C.uint64_t)(Seq), cgoAllocsUnknown
	cEventSeq, cEventSeqAllocMap := (C.uint64_t)(EventSeq), cgoAllocsUnknown
	Key = safeString(Key)
	cKey, cKeyAllocMap := unpackPCharString(Key)
	Value = safeString(Value)
	cValue, cValueAllocMap := unpackPCharString(Value)
	cIndex, cIndexAllocMap := (C._Bool)(Index), cgoAllocsUnknown
	C.cosmos_event_attribute_append(cBuilder, cSeq, cEventSeq, cKey, cValue, cIndex)
	runtime.KeepAlive(cIndexAllocMap)
	runtime.KeepAlive(Value)
	runtime.KeepAlive(cValueAllocMap)
	runtime.KeepAlive(Key)
	runtime.KeepAlive(cKeyAllocMap)
	runtime.KeepAlive(cEventSeqAllocMap)
	runtime.KeepAlive(cSeqAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// CosmosValidatorUpdateAppend function as declared in include/libmamoru_sniffer_go.h:255
func CosmosValidatorUpdateAppend(Builder *FfiCosmosBlockchainDataBuilderT, BlockSeq uint64, PubKey SliceRefUint8T, Power int64) {
	cBuilder, cBuilderAllocMap := (*C.FfiCosmosBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cBlockSeq, cBlockSeqAllocMap := (C.uint64_t)(BlockSeq), cgoAllocsUnknown
	cPubKey, cPubKeyAllocMap := PubKey.PassValue()
	cPower, cPowerAllocMap := (C.int64_t)(Power), cgoAllocsUnknown
	C.cosmos_validator_update_append(cBuilder, cBlockSeq, cPubKey, cPower)
	runtime.KeepAlive(cPowerAllocMap)
	runtime.KeepAlive(cPubKeyAllocMap)
	runtime.KeepAlive(cBlockSeqAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// CosmosMisbehaviorAppend function as declared in include/libmamoru_sniffer_go.h:261
func CosmosMisbehaviorAppend(Builder *FfiCosmosBlockchainDataBuilderT, BlockSeq uint64, Typ string, ValidatorPower int64, ValidatorAddress string, Height int64, Time int64, TotalVotingPower int64) {
	cBuilder, cBuilderAllocMap := (*C.FfiCosmosBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cBlockSeq, cBlockSeqAllocMap := (C.uint64_t)(BlockSeq), cgoAllocsUnknown
	Typ = safeString(Typ)
	cTyp, cTypAllocMap := unpackPCharString(Typ)
	cValidatorPower, cValidatorPowerAllocMap := (C.int64_t)(ValidatorPower), cgoAllocsUnknown
	ValidatorAddress = safeString(ValidatorAddress)
	cValidatorAddress, cValidatorAddressAllocMap := unpackPCharString(ValidatorAddress)
	cHeight, cHeightAllocMap := (C.int64_t)(Height), cgoAllocsUnknown
	cTime, cTimeAllocMap := (C.int64_t)(Time), cgoAllocsUnknown
	cTotalVotingPower, cTotalVotingPowerAllocMap := (C.int64_t)(TotalVotingPower), cgoAllocsUnknown
	C.cosmos_misbehavior_append(cBuilder, cBlockSeq, cTyp, cValidatorPower, cValidatorAddress, cHeight, cTime, cTotalVotingPower)
	runtime.KeepAlive(cTotalVotingPowerAllocMap)
	runtime.KeepAlive(cTimeAllocMap)
	runtime.KeepAlive(cHeightAllocMap)
	runtime.KeepAlive(ValidatorAddress)
	runtime.KeepAlive(cValidatorAddressAllocMap)
	runtime.KeepAlive(cValidatorPowerAllocMap)
	runtime.KeepAlive(Typ)
	runtime.KeepAlive(cTypAllocMap)
	runtime.KeepAlive(cBlockSeqAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// CosmosVoteInfosAppend function as declared in include/libmamoru_sniffer_go.h:271
func CosmosVoteInfosAppend(Builder *FfiCosmosBlockchainDataBuilderT, BlockSeq uint64, Validator SliceRefUint8T, SignedLastBlock bool) {
	cBuilder, cBuilderAllocMap := (*C.FfiCosmosBlockchainDataBuilder_t)(unsafe.Pointer(Builder)), cgoAllocsUnknown
	cBlockSeq, cBlockSeqAllocMap := (C.uint64_t)(BlockSeq), cgoAllocsUnknown
	cValidator, cValidatorAllocMap := Validator.PassValue()
	cSignedLastBlock, cSignedLastBlockAllocMap := (C._Bool)(SignedLastBlock), cgoAllocsUnknown
	C.cosmos_vote_infos_append(cBuilder, cBlockSeq, cValidator, cSignedLastBlock)
	runtime.KeepAlive(cSignedLastBlockAllocMap)
	runtime.KeepAlive(cValidatorAllocMap)
	runtime.KeepAlive(cBlockSeqAllocMap)
	runtime.KeepAlive(cBuilderAllocMap)
}

// NewSniffer function as declared in include/libmamoru_sniffer_go.h:279
func NewSniffer() *FfiSnifferResultT {
	__ret := C.new_sniffer()
	__v := *(**FfiSnifferResultT)(unsafe.Pointer(&__ret))
	return __v
}

// SnifferResultGetErrorMessage function as declared in include/libmamoru_sniffer_go.h:284
func SnifferResultGetErrorMessage(Result *FfiSnifferResultT) *byte {
	cResult, cResultAllocMap := (*C.FfiSnifferResult_t)(unsafe.Pointer(Result)), cgoAllocsUnknown
	__ret := C.sniffer_result_get_error_message(cResult)
	runtime.KeepAlive(cResultAllocMap)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// SnifferResultGetSniffer function as declared in include/libmamoru_sniffer_go.h:291
func SnifferResultGetSniffer(Result *FfiSnifferResultT) *FfiSnifferT {
	cResult, cResultAllocMap := (*C.FfiSnifferResult_t)(unsafe.Pointer(Result)), cgoAllocsUnknown
	__ret := C.sniffer_result_get_sniffer(cResult)
	runtime.KeepAlive(cResultAllocMap)
	__v := *(**FfiSnifferT)(unsafe.Pointer(&__ret))
	return __v
}

// NewValueData function as declared in include/libmamoru_sniffer_go.h:301
func NewValueData(Value *FfiValueT) *FfiValueDataT {
	cValue, cValueAllocMap := (*C.FfiValue_t)(unsafe.Pointer(Value)), cgoAllocsUnknown
	__ret := C.new_value_data(cValue)
	runtime.KeepAlive(cValueAllocMap)
	__v := *(**FfiValueDataT)(unsafe.Pointer(&__ret))
	return __v
}

// NewValueBool function as declared in include/libmamoru_sniffer_go.h:304
func NewValueBool(Data bool) *FfiValueT {
	cData, cDataAllocMap := (C._Bool)(Data), cgoAllocsUnknown
	__ret := C.new_value_bool(cData)
	runtime.KeepAlive(cDataAllocMap)
	__v := *(**FfiValueT)(unsafe.Pointer(&__ret))
	return __v
}

// NewValueU64 function as declared in include/libmamoru_sniffer_go.h:307
func NewValueU64(Data uint64) *FfiValueT {
	cData, cDataAllocMap := (C.uint64_t)(Data), cgoAllocsUnknown
	__ret := C.new_value_u64(cData)
	runtime.KeepAlive(cDataAllocMap)
	__v := *(**FfiValueT)(unsafe.Pointer(&__ret))
	return __v
}

// NewValueList function as declared in include/libmamoru_sniffer_go.h:310
func NewValueList() *FfiValueT {
	__ret := C.new_value_list()
	__v := *(**FfiValueT)(unsafe.Pointer(&__ret))
	return __v
}

// ValueListAppend function as declared in include/libmamoru_sniffer_go.h:316
func ValueListAppend(ValueList *FfiValueT, Data *FfiValueT) bool {
	cValueList, cValueListAllocMap := (*C.FfiValue_t)(unsafe.Pointer(ValueList)), cgoAllocsUnknown
	cData, cDataAllocMap := (*C.FfiValue_t)(unsafe.Pointer(Data)), cgoAllocsUnknown
	__ret := C.value_list_append(cValueList, cData)
	runtime.KeepAlive(cDataAllocMap)
	runtime.KeepAlive(cValueListAllocMap)
	__v := (bool)(__ret)
	return __v
}

// NewValueStruct function as declared in include/libmamoru_sniffer_go.h:320
func NewValueStruct(Ty string) *FfiValueT {
	Ty = safeString(Ty)
	cTy, cTyAllocMap := unpackPCharString(Ty)
	__ret := C.new_value_struct(cTy)
	runtime.KeepAlive(Ty)
	runtime.KeepAlive(cTyAllocMap)
	__v := *(**FfiValueT)(unsafe.Pointer(&__ret))
	return __v
}

// ValueStructAddField function as declared in include/libmamoru_sniffer_go.h:327
func ValueStructAddField(ValueStruct *FfiValueT, Key string, Data *FfiValueT) bool {
	cValueStruct, cValueStructAllocMap := (*C.FfiValue_t)(unsafe.Pointer(ValueStruct)), cgoAllocsUnknown
	Key = safeString(Key)
	cKey, cKeyAllocMap := unpackPCharString(Key)
	cData, cDataAllocMap := (*C.FfiValue_t)(unsafe.Pointer(Data)), cgoAllocsUnknown
	__ret := C.value_struct_add_field(cValueStruct, cKey, cData)
	runtime.KeepAlive(cDataAllocMap)
	runtime.KeepAlive(Key)
	runtime.KeepAlive(cKeyAllocMap)
	runtime.KeepAlive(cValueStructAllocMap)
	__v := (bool)(__ret)
	return __v
}
