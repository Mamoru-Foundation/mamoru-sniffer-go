/*! \file */
/*******************************************
 *                                         *
 *  File auto-generated by `::safer_ffi`.  *
 *                                         *
 *  Do not manually edit this file.        *
 *                                         *
 *******************************************/

#ifndef __RUST_MAMORU_SNIFFER_C__
#define __RUST_MAMORU_SNIFFER_C__

#ifdef __cplusplus
extern "C" {
#endif

typedef struct FfiTransactionBatch FfiTransactionBatch_t;

FfiTransactionBatch_t * new_transaction_batch (void);


#include <stddef.h>
#include <stdint.h>

/** \brief
 *  `&'lt [T]` but with a guaranteed `#[repr(C)]` layout.
 *
 *  # C layout (for some given type T)
 *
 *  ```c
 *  typedef struct {
 *      // Cannot be NULL
 *      T * ptr;
 *      size_t len;
 *  } slice_T;
 *  ```
 *
 *  # Nullable pointer?
 *
 *  If you want to support the above typedef, but where the `ptr` field is
 *  allowed to be `NULL` (with the contents of `len` then being undefined)
 *  use the `Option< slice_ptr<_> >` type.
 */
typedef struct slice_ref_uint8 {

    uint8_t const * ptr;

    size_t len;

} slice_ref_uint8_t;

void transaction_batch_append (
    FfiTransactionBatch_t * batch,
    uint32_t tx_index,
    char const * tx_hash,
    uint64_t block_index,
    uint8_t typ,
    uint64_t nonce,
    uint64_t status,
    char const * from,
    char const * to,
    uint64_t value,
    uint64_t fee,
    uint64_t gas_price,
    uint64_t gas_limit,
    uint64_t gas_used,
    slice_ref_uint8_t input,
    double size);

typedef struct FfiBlockchainData FfiBlockchainData_t;

/** \brief
 *  Frees `object` argument.
 */
FfiBlockchainData_t * transaction_batch_finish (
    FfiTransactionBatch_t * object);

typedef struct FfiCallTraceBatch FfiCallTraceBatch_t;

FfiCallTraceBatch_t * new_call_trace_batch (void);

void call_trace_batch_append (
    FfiCallTraceBatch_t * batch,
    uint32_t seq,
    uint32_t depth,
    uint32_t tx_index,
    uint64_t block_index,
    char const * typ,
    char const * from,
    char const * to,
    uint64_t value,
    uint64_t gas_limit,
    uint64_t gas_used,
    slice_ref_uint8_t input);

/** \brief
 *  Frees `object` argument.
 */
FfiBlockchainData_t * call_trace_batch_finish (
    FfiCallTraceBatch_t * object);

typedef struct FfiBlockBatch FfiBlockBatch_t;

FfiBlockBatch_t * new_block_batch (void);

void block_batch_append (
    FfiBlockBatch_t * batch,
    uint64_t block_index,
    char const * hash,
    char const * parent_hash,
    char const * state_root,
    uint64_t nonce,
    char const * status,
    uint64_t timestamp,
    slice_ref_uint8_t block_reward,
    char const * fee_recipient,
    uint64_t total_difficulty,
    double size,
    uint64_t gas_used,
    uint64_t gas_limit);

/** \brief
 *  Frees `object` argument.
 */
FfiBlockchainData_t * block_batch_finish (
    FfiBlockBatch_t * object);

typedef struct FfiEventBatch FfiEventBatch_t;

FfiEventBatch_t * new_event_batch (void);

void event_batch_append (
    FfiEventBatch_t * batch,
    uint32_t index,
    char const * address,
    uint64_t block_number,
    char const * tx_hash,
    uint32_t tx_index,
    char const * block_hash,
    slice_ref_uint8_t topic0,
    slice_ref_uint8_t topic1,
    slice_ref_uint8_t topic2,
    slice_ref_uint8_t topic3,
    slice_ref_uint8_t topic4,
    slice_ref_uint8_t data);

/** \brief
 *  Frees `object` argument.
 */
FfiBlockchainData_t * event_batch_finish (
    FfiEventBatch_t * object);

typedef struct FfiSnifferResult FfiSnifferResult_t;

FfiSnifferResult_t * new_sniffer (void);

/** \brief
 *  Returns a pointer to C-like string or NULL if there is no error.
 */
char * sniffer_result_get_error_message (
    FfiSnifferResult_t const * result);

typedef struct FfiSniffer FfiSniffer_t;

/** \brief
 *  Returns a pointer to `FfiSniffer` or NULL.
 *  Frees `result` arguments, so be sure to call `sniffer_result_get_error_message` to get an error message.
 */
FfiSniffer_t * sniffer_result_get_sniffer (
    FfiSnifferResult_t * result);

typedef struct FfiBlockchainDataCtx FfiBlockchainDataCtx_t;

/** \brief
 *  Frees `data` argument.
 */
void sniffer_observe_data (
    FfiSniffer_t const * sniffer,
    FfiBlockchainDataCtx_t * data);

typedef struct FfiBlockchainDataCtxBuilder FfiBlockchainDataCtxBuilder_t;

FfiBlockchainDataCtxBuilder_t * new_blockchain_data_ctx_builder (void);


#include <stdbool.h>

/** \brief
 *  Returns `true` is success.
 *  Frees `data` argument.
 */
bool blockchain_data_ctx_builder_add_data (
    FfiBlockchainDataCtxBuilder_t const * builder,
    FfiBlockchainData_t * data);

FfiBlockchainDataCtx_t * blockchain_data_ctx_builder_finish (
    FfiBlockchainDataCtxBuilder_t * builder,
    char const * tx_id,
    char const * tx_hash);

typedef struct FfiValue FfiValue_t;

typedef struct FfiValueData FfiValueData_t;

/** \brief
 *  Frees `value` argument.
 */
FfiValueData_t * new_value_data (
    FfiValue_t * value);

FfiValue_t * new_value_bool (
    bool data);

FfiValue_t * new_value_u64 (
    uint64_t data);

FfiValue_t * new_value_list (void);

/** \brief
 *  Returns `true` if success.
 *  Frees `data` argument.
 */
bool value_list_append (
    FfiValue_t * value_list,
    FfiValue_t * data);

FfiValue_t * new_value_struct (
    char const * ty);

/** \brief
 *  Returns `true` if success.
 *  Frees `data` argument.
 */
bool value_struct_add_field (
    FfiValue_t * value_struct,
    char const * key,
    FfiValue_t * data);


#ifdef __cplusplus
} /* extern "C" */
#endif

#endif /* __RUST_MAMORU_SNIFFER_C__ */
