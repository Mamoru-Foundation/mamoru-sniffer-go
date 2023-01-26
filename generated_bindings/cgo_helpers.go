// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 26 Jan 2023 16:39:17 CET.
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
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

// Ref returns a reference to C object as it is.
func (x *FfiTransactionBatchT) Ref() *C.FfiTransactionBatch_t {
	if x == nil {
		return nil
	}
	return (*C.FfiTransactionBatch_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiTransactionBatchT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiTransactionBatchTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiTransactionBatchTRef(ref unsafe.Pointer) *FfiTransactionBatchT {
	return (*FfiTransactionBatchT)(ref)
}

// NewFfiTransactionBatchT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiTransactionBatchT() *FfiTransactionBatchT {
	return (*FfiTransactionBatchT)(allocFfiTransactionBatchTMemory(1))
}

// allocFfiTransactionBatchTMemory allocates memory for type C.FfiTransactionBatch_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiTransactionBatchTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiTransactionBatchTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiTransactionBatchTValue = unsafe.Sizeof([1]C.FfiTransactionBatch_t{})

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiTransactionBatchT) PassRef() *C.FfiTransactionBatch_t {
	if x == nil {
		x = (*FfiTransactionBatchT)(allocFfiTransactionBatchTMemory(1))
	}
	return (*C.FfiTransactionBatch_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FfiBlockchainDataT) Ref() *C.FfiBlockchainData_t {
	if x == nil {
		return nil
	}
	return (*C.FfiBlockchainData_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiBlockchainDataT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiBlockchainDataTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiBlockchainDataTRef(ref unsafe.Pointer) *FfiBlockchainDataT {
	return (*FfiBlockchainDataT)(ref)
}

// NewFfiBlockchainDataT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiBlockchainDataT() *FfiBlockchainDataT {
	return (*FfiBlockchainDataT)(allocFfiBlockchainDataTMemory(1))
}

// allocFfiBlockchainDataTMemory allocates memory for type C.FfiBlockchainData_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiBlockchainDataTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiBlockchainDataTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiBlockchainDataTValue = unsafe.Sizeof([1]C.FfiBlockchainData_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiBlockchainDataT) PassRef() *C.FfiBlockchainData_t {
	if x == nil {
		x = (*FfiBlockchainDataT)(allocFfiBlockchainDataTMemory(1))
	}
	return (*C.FfiBlockchainData_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FfiTransactionArgBatchT) Ref() *C.FfiTransactionArgBatch_t {
	if x == nil {
		return nil
	}
	return (*C.FfiTransactionArgBatch_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiTransactionArgBatchT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiTransactionArgBatchTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiTransactionArgBatchTRef(ref unsafe.Pointer) *FfiTransactionArgBatchT {
	return (*FfiTransactionArgBatchT)(ref)
}

// NewFfiTransactionArgBatchT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiTransactionArgBatchT() *FfiTransactionArgBatchT {
	return (*FfiTransactionArgBatchT)(allocFfiTransactionArgBatchTMemory(1))
}

// allocFfiTransactionArgBatchTMemory allocates memory for type C.FfiTransactionArgBatch_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiTransactionArgBatchTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiTransactionArgBatchTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiTransactionArgBatchTValue = unsafe.Sizeof([1]C.FfiTransactionArgBatch_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiTransactionArgBatchT) PassRef() *C.FfiTransactionArgBatch_t {
	if x == nil {
		x = (*FfiTransactionArgBatchT)(allocFfiTransactionArgBatchTMemory(1))
	}
	return (*C.FfiTransactionArgBatch_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FfiCallTraceBatchT) Ref() *C.FfiCallTraceBatch_t {
	if x == nil {
		return nil
	}
	return (*C.FfiCallTraceBatch_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiCallTraceBatchT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiCallTraceBatchTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiCallTraceBatchTRef(ref unsafe.Pointer) *FfiCallTraceBatchT {
	return (*FfiCallTraceBatchT)(ref)
}

// NewFfiCallTraceBatchT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiCallTraceBatchT() *FfiCallTraceBatchT {
	return (*FfiCallTraceBatchT)(allocFfiCallTraceBatchTMemory(1))
}

// allocFfiCallTraceBatchTMemory allocates memory for type C.FfiCallTraceBatch_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiCallTraceBatchTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiCallTraceBatchTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiCallTraceBatchTValue = unsafe.Sizeof([1]C.FfiCallTraceBatch_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiCallTraceBatchT) PassRef() *C.FfiCallTraceBatch_t {
	if x == nil {
		x = (*FfiCallTraceBatchT)(allocFfiCallTraceBatchTMemory(1))
	}
	return (*C.FfiCallTraceBatch_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FfiCallTraceArgBatchT) Ref() *C.FfiCallTraceArgBatch_t {
	if x == nil {
		return nil
	}
	return (*C.FfiCallTraceArgBatch_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiCallTraceArgBatchT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiCallTraceArgBatchTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiCallTraceArgBatchTRef(ref unsafe.Pointer) *FfiCallTraceArgBatchT {
	return (*FfiCallTraceArgBatchT)(ref)
}

// NewFfiCallTraceArgBatchT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiCallTraceArgBatchT() *FfiCallTraceArgBatchT {
	return (*FfiCallTraceArgBatchT)(allocFfiCallTraceArgBatchTMemory(1))
}

// allocFfiCallTraceArgBatchTMemory allocates memory for type C.FfiCallTraceArgBatch_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiCallTraceArgBatchTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiCallTraceArgBatchTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiCallTraceArgBatchTValue = unsafe.Sizeof([1]C.FfiCallTraceArgBatch_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiCallTraceArgBatchT) PassRef() *C.FfiCallTraceArgBatch_t {
	if x == nil {
		x = (*FfiCallTraceArgBatchT)(allocFfiCallTraceArgBatchTMemory(1))
	}
	return (*C.FfiCallTraceArgBatch_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FfiBlockBatchT) Ref() *C.FfiBlockBatch_t {
	if x == nil {
		return nil
	}
	return (*C.FfiBlockBatch_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiBlockBatchT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiBlockBatchTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiBlockBatchTRef(ref unsafe.Pointer) *FfiBlockBatchT {
	return (*FfiBlockBatchT)(ref)
}

// NewFfiBlockBatchT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiBlockBatchT() *FfiBlockBatchT {
	return (*FfiBlockBatchT)(allocFfiBlockBatchTMemory(1))
}

// allocFfiBlockBatchTMemory allocates memory for type C.FfiBlockBatch_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiBlockBatchTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiBlockBatchTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiBlockBatchTValue = unsafe.Sizeof([1]C.FfiBlockBatch_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiBlockBatchT) PassRef() *C.FfiBlockBatch_t {
	if x == nil {
		x = (*FfiBlockBatchT)(allocFfiBlockBatchTMemory(1))
	}
	return (*C.FfiBlockBatch_t)(unsafe.Pointer(x))
}

// allocSliceRefUint8TMemory allocates memory for type C.slice_ref_uint8_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSliceRefUint8TMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSliceRefUint8TValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfSliceRefUint8TValue = unsafe.Sizeof([1]C.slice_ref_uint8_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *SliceRefUint8T) Ref() *C.slice_ref_uint8_t {
	if x == nil {
		return nil
	}
	return x.refd8594c66
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *SliceRefUint8T) Free() {
	if x != nil && x.allocsd8594c66 != nil {
		x.allocsd8594c66.(*cgoAllocMap).Free()
		x.refd8594c66 = nil
	}
}

// NewSliceRefUint8TRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewSliceRefUint8TRef(ref unsafe.Pointer) *SliceRefUint8T {
	if ref == nil {
		return nil
	}
	obj := new(SliceRefUint8T)
	obj.refd8594c66 = (*C.slice_ref_uint8_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *SliceRefUint8T) PassRef() (*C.slice_ref_uint8_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd8594c66 != nil {
		return x.refd8594c66, nil
	}
	memd8594c66 := allocSliceRefUint8TMemory(1)
	refd8594c66 := (*C.slice_ref_uint8_t)(memd8594c66)
	allocsd8594c66 := new(cgoAllocMap)
	allocsd8594c66.Add(memd8594c66)

	var cptr_allocs *cgoAllocMap
	refd8594c66.ptr, cptr_allocs = *(**C.uint8_t)(unsafe.Pointer(&x.Ptr)), cgoAllocsUnknown
	allocsd8594c66.Borrow(cptr_allocs)

	var clen_allocs *cgoAllocMap
	refd8594c66.len, clen_allocs = (C.size_t)(x.Len), cgoAllocsUnknown
	allocsd8594c66.Borrow(clen_allocs)

	x.refd8594c66 = refd8594c66
	x.allocsd8594c66 = allocsd8594c66
	return refd8594c66, allocsd8594c66

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x SliceRefUint8T) PassValue() (C.slice_ref_uint8_t, *cgoAllocMap) {
	if x.refd8594c66 != nil {
		return *x.refd8594c66, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *SliceRefUint8T) Deref() {
	if x.refd8594c66 == nil {
		return
	}
	x.Ptr = (*byte)(unsafe.Pointer(x.refd8594c66.ptr))
	x.Len = (uint64)(x.refd8594c66.len)
}

// Ref returns a reference to C object as it is.
func (x *FfiEventBatchT) Ref() *C.FfiEventBatch_t {
	if x == nil {
		return nil
	}
	return (*C.FfiEventBatch_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiEventBatchT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiEventBatchTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiEventBatchTRef(ref unsafe.Pointer) *FfiEventBatchT {
	return (*FfiEventBatchT)(ref)
}

// NewFfiEventBatchT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiEventBatchT() *FfiEventBatchT {
	return (*FfiEventBatchT)(allocFfiEventBatchTMemory(1))
}

// allocFfiEventBatchTMemory allocates memory for type C.FfiEventBatch_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiEventBatchTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiEventBatchTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiEventBatchTValue = unsafe.Sizeof([1]C.FfiEventBatch_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiEventBatchT) PassRef() *C.FfiEventBatch_t {
	if x == nil {
		x = (*FfiEventBatchT)(allocFfiEventBatchTMemory(1))
	}
	return (*C.FfiEventBatch_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FfiEventTopicBatchT) Ref() *C.FfiEventTopicBatch_t {
	if x == nil {
		return nil
	}
	return (*C.FfiEventTopicBatch_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiEventTopicBatchT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiEventTopicBatchTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiEventTopicBatchTRef(ref unsafe.Pointer) *FfiEventTopicBatchT {
	return (*FfiEventTopicBatchT)(ref)
}

// NewFfiEventTopicBatchT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiEventTopicBatchT() *FfiEventTopicBatchT {
	return (*FfiEventTopicBatchT)(allocFfiEventTopicBatchTMemory(1))
}

// allocFfiEventTopicBatchTMemory allocates memory for type C.FfiEventTopicBatch_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiEventTopicBatchTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiEventTopicBatchTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiEventTopicBatchTValue = unsafe.Sizeof([1]C.FfiEventTopicBatch_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiEventTopicBatchT) PassRef() *C.FfiEventTopicBatch_t {
	if x == nil {
		x = (*FfiEventTopicBatchT)(allocFfiEventTopicBatchTMemory(1))
	}
	return (*C.FfiEventTopicBatch_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FfiSnifferResultT) Ref() *C.FfiSnifferResult_t {
	if x == nil {
		return nil
	}
	return (*C.FfiSnifferResult_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiSnifferResultT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiSnifferResultTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiSnifferResultTRef(ref unsafe.Pointer) *FfiSnifferResultT {
	return (*FfiSnifferResultT)(ref)
}

// NewFfiSnifferResultT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiSnifferResultT() *FfiSnifferResultT {
	return (*FfiSnifferResultT)(allocFfiSnifferResultTMemory(1))
}

// allocFfiSnifferResultTMemory allocates memory for type C.FfiSnifferResult_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiSnifferResultTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiSnifferResultTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiSnifferResultTValue = unsafe.Sizeof([1]C.FfiSnifferResult_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiSnifferResultT) PassRef() *C.FfiSnifferResult_t {
	if x == nil {
		x = (*FfiSnifferResultT)(allocFfiSnifferResultTMemory(1))
	}
	return (*C.FfiSnifferResult_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FfiSnifferT) Ref() *C.FfiSniffer_t {
	if x == nil {
		return nil
	}
	return (*C.FfiSniffer_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiSnifferT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiSnifferTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiSnifferTRef(ref unsafe.Pointer) *FfiSnifferT {
	return (*FfiSnifferT)(ref)
}

// NewFfiSnifferT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiSnifferT() *FfiSnifferT {
	return (*FfiSnifferT)(allocFfiSnifferTMemory(1))
}

// allocFfiSnifferTMemory allocates memory for type C.FfiSniffer_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiSnifferTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiSnifferTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiSnifferTValue = unsafe.Sizeof([1]C.FfiSniffer_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiSnifferT) PassRef() *C.FfiSniffer_t {
	if x == nil {
		x = (*FfiSnifferT)(allocFfiSnifferTMemory(1))
	}
	return (*C.FfiSniffer_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FfiBlockchainDataCtxT) Ref() *C.FfiBlockchainDataCtx_t {
	if x == nil {
		return nil
	}
	return (*C.FfiBlockchainDataCtx_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiBlockchainDataCtxT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiBlockchainDataCtxTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiBlockchainDataCtxTRef(ref unsafe.Pointer) *FfiBlockchainDataCtxT {
	return (*FfiBlockchainDataCtxT)(ref)
}

// NewFfiBlockchainDataCtxT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiBlockchainDataCtxT() *FfiBlockchainDataCtxT {
	return (*FfiBlockchainDataCtxT)(allocFfiBlockchainDataCtxTMemory(1))
}

// allocFfiBlockchainDataCtxTMemory allocates memory for type C.FfiBlockchainDataCtx_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiBlockchainDataCtxTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiBlockchainDataCtxTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiBlockchainDataCtxTValue = unsafe.Sizeof([1]C.FfiBlockchainDataCtx_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiBlockchainDataCtxT) PassRef() *C.FfiBlockchainDataCtx_t {
	if x == nil {
		x = (*FfiBlockchainDataCtxT)(allocFfiBlockchainDataCtxTMemory(1))
	}
	return (*C.FfiBlockchainDataCtx_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FfiBlockchainDataCtxBuilderT) Ref() *C.FfiBlockchainDataCtxBuilder_t {
	if x == nil {
		return nil
	}
	return (*C.FfiBlockchainDataCtxBuilder_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiBlockchainDataCtxBuilderT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiBlockchainDataCtxBuilderTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiBlockchainDataCtxBuilderTRef(ref unsafe.Pointer) *FfiBlockchainDataCtxBuilderT {
	return (*FfiBlockchainDataCtxBuilderT)(ref)
}

// NewFfiBlockchainDataCtxBuilderT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiBlockchainDataCtxBuilderT() *FfiBlockchainDataCtxBuilderT {
	return (*FfiBlockchainDataCtxBuilderT)(allocFfiBlockchainDataCtxBuilderTMemory(1))
}

// allocFfiBlockchainDataCtxBuilderTMemory allocates memory for type C.FfiBlockchainDataCtxBuilder_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiBlockchainDataCtxBuilderTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiBlockchainDataCtxBuilderTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiBlockchainDataCtxBuilderTValue = unsafe.Sizeof([1]C.FfiBlockchainDataCtxBuilder_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiBlockchainDataCtxBuilderT) PassRef() *C.FfiBlockchainDataCtxBuilder_t {
	if x == nil {
		x = (*FfiBlockchainDataCtxBuilderT)(allocFfiBlockchainDataCtxBuilderTMemory(1))
	}
	return (*C.FfiBlockchainDataCtxBuilder_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FfiValueT) Ref() *C.FfiValue_t {
	if x == nil {
		return nil
	}
	return (*C.FfiValue_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiValueT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiValueTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiValueTRef(ref unsafe.Pointer) *FfiValueT {
	return (*FfiValueT)(ref)
}

// NewFfiValueT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiValueT() *FfiValueT {
	return (*FfiValueT)(allocFfiValueTMemory(1))
}

// allocFfiValueTMemory allocates memory for type C.FfiValue_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiValueTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiValueTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiValueTValue = unsafe.Sizeof([1]C.FfiValue_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiValueT) PassRef() *C.FfiValue_t {
	if x == nil {
		x = (*FfiValueT)(allocFfiValueTMemory(1))
	}
	return (*C.FfiValue_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FfiValueDataT) Ref() *C.FfiValueData_t {
	if x == nil {
		return nil
	}
	return (*C.FfiValueData_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiValueDataT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiValueDataTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiValueDataTRef(ref unsafe.Pointer) *FfiValueDataT {
	return (*FfiValueDataT)(ref)
}

// NewFfiValueDataT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiValueDataT() *FfiValueDataT {
	return (*FfiValueDataT)(allocFfiValueDataTMemory(1))
}

// allocFfiValueDataTMemory allocates memory for type C.FfiValueData_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiValueDataTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiValueDataTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiValueDataTValue = unsafe.Sizeof([1]C.FfiValueData_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiValueDataT) PassRef() *C.FfiValueData_t {
	if x == nil {
		x = (*FfiValueDataT)(allocFfiValueDataTMemory(1))
	}
	return (*C.FfiValueData_t)(unsafe.Pointer(x))
}

// safeString ensures that the string is NULL-terminated, a NULL-terminated copy is created otherwise.
func safeString(str string) string {
	if len(str) > 0 && str[len(str)-1] != '\x00' {
		str = str + "\x00"
	} else if len(str) == 0 {
		str = "\x00"
	}
	return str
}

// unpackPCharString copies the data from Go string as *C.char.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	allocs := new(cgoAllocMap)
	defer runtime.SetFinalizer(allocs, func(a *cgoAllocMap) {
		go a.Free()
	})

	str = safeString(str)
	mem0 := unsafe.Pointer(C.CString(str))
	allocs.Add(mem0)
	return (*C.char)(mem0), allocs
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}
