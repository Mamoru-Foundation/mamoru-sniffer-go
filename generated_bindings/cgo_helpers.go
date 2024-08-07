// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Fri, 26 Jul 2024 13:47:09 WEST.
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
func (x *FfiCosmosBlockchainDataBuilderT) Ref() *C.FfiCosmosBlockchainDataBuilder_t {
	if x == nil {
		return nil
	}
	return (*C.FfiCosmosBlockchainDataBuilder_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiCosmosBlockchainDataBuilderT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiCosmosBlockchainDataBuilderTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiCosmosBlockchainDataBuilderTRef(ref unsafe.Pointer) *FfiCosmosBlockchainDataBuilderT {
	return (*FfiCosmosBlockchainDataBuilderT)(ref)
}

// NewFfiCosmosBlockchainDataBuilderT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiCosmosBlockchainDataBuilderT() *FfiCosmosBlockchainDataBuilderT {
	return (*FfiCosmosBlockchainDataBuilderT)(allocFfiCosmosBlockchainDataBuilderTMemory(1))
}

// allocFfiCosmosBlockchainDataBuilderTMemory allocates memory for type C.FfiCosmosBlockchainDataBuilder_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiCosmosBlockchainDataBuilderTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiCosmosBlockchainDataBuilderTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiCosmosBlockchainDataBuilderTValue = unsafe.Sizeof([1]C.FfiCosmosBlockchainDataBuilder_t{})

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
func (x *FfiCosmosBlockchainDataBuilderT) PassRef() *C.FfiCosmosBlockchainDataBuilder_t {
	if x == nil {
		x = (*FfiCosmosBlockchainDataBuilderT)(allocFfiCosmosBlockchainDataBuilderTMemory(1))
	}
	return (*C.FfiCosmosBlockchainDataBuilder_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FfiCosmosBlockchainDataCtxT) Ref() *C.FfiCosmosBlockchainDataCtx_t {
	if x == nil {
		return nil
	}
	return (*C.FfiCosmosBlockchainDataCtx_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiCosmosBlockchainDataCtxT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiCosmosBlockchainDataCtxTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiCosmosBlockchainDataCtxTRef(ref unsafe.Pointer) *FfiCosmosBlockchainDataCtxT {
	return (*FfiCosmosBlockchainDataCtxT)(ref)
}

// NewFfiCosmosBlockchainDataCtxT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiCosmosBlockchainDataCtxT() *FfiCosmosBlockchainDataCtxT {
	return (*FfiCosmosBlockchainDataCtxT)(allocFfiCosmosBlockchainDataCtxTMemory(1))
}

// allocFfiCosmosBlockchainDataCtxTMemory allocates memory for type C.FfiCosmosBlockchainDataCtx_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiCosmosBlockchainDataCtxTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiCosmosBlockchainDataCtxTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiCosmosBlockchainDataCtxTValue = unsafe.Sizeof([1]C.FfiCosmosBlockchainDataCtx_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiCosmosBlockchainDataCtxT) PassRef() *C.FfiCosmosBlockchainDataCtx_t {
	if x == nil {
		x = (*FfiCosmosBlockchainDataCtxT)(allocFfiCosmosBlockchainDataCtxTMemory(1))
	}
	return (*C.FfiCosmosBlockchainDataCtx_t)(unsafe.Pointer(x))
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

// Ref returns a reference to C object as it is.
func (x *FfiEvmBlockchainDataBuilderT) Ref() *C.FfiEvmBlockchainDataBuilder_t {
	if x == nil {
		return nil
	}
	return (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiEvmBlockchainDataBuilderT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiEvmBlockchainDataBuilderTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiEvmBlockchainDataBuilderTRef(ref unsafe.Pointer) *FfiEvmBlockchainDataBuilderT {
	return (*FfiEvmBlockchainDataBuilderT)(ref)
}

// NewFfiEvmBlockchainDataBuilderT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiEvmBlockchainDataBuilderT() *FfiEvmBlockchainDataBuilderT {
	return (*FfiEvmBlockchainDataBuilderT)(allocFfiEvmBlockchainDataBuilderTMemory(1))
}

// allocFfiEvmBlockchainDataBuilderTMemory allocates memory for type C.FfiEvmBlockchainDataBuilder_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiEvmBlockchainDataBuilderTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiEvmBlockchainDataBuilderTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiEvmBlockchainDataBuilderTValue = unsafe.Sizeof([1]C.FfiEvmBlockchainDataBuilder_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiEvmBlockchainDataBuilderT) PassRef() *C.FfiEvmBlockchainDataBuilder_t {
	if x == nil {
		x = (*FfiEvmBlockchainDataBuilderT)(allocFfiEvmBlockchainDataBuilderTMemory(1))
	}
	return (*C.FfiEvmBlockchainDataBuilder_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *FfiEvmBlockchainDataCtxT) Ref() *C.FfiEvmBlockchainDataCtx_t {
	if x == nil {
		return nil
	}
	return (*C.FfiEvmBlockchainDataCtx_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *FfiEvmBlockchainDataCtxT) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewFfiEvmBlockchainDataCtxTRef converts the C object reference into a raw struct reference without wrapping.
func NewFfiEvmBlockchainDataCtxTRef(ref unsafe.Pointer) *FfiEvmBlockchainDataCtxT {
	return (*FfiEvmBlockchainDataCtxT)(ref)
}

// NewFfiEvmBlockchainDataCtxT allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewFfiEvmBlockchainDataCtxT() *FfiEvmBlockchainDataCtxT {
	return (*FfiEvmBlockchainDataCtxT)(allocFfiEvmBlockchainDataCtxTMemory(1))
}

// allocFfiEvmBlockchainDataCtxTMemory allocates memory for type C.FfiEvmBlockchainDataCtx_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFfiEvmBlockchainDataCtxTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFfiEvmBlockchainDataCtxTValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfFfiEvmBlockchainDataCtxTValue = unsafe.Sizeof([1]C.FfiEvmBlockchainDataCtx_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *FfiEvmBlockchainDataCtxT) PassRef() *C.FfiEvmBlockchainDataCtx_t {
	if x == nil {
		x = (*FfiEvmBlockchainDataCtxT)(allocFfiEvmBlockchainDataCtxTMemory(1))
	}
	return (*C.FfiEvmBlockchainDataCtx_t)(unsafe.Pointer(x))
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
