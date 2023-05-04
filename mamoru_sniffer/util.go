package mamoru_sniffer

/*
#include <stdlib.h>
*/
import "C"
import (
	"github.com/Mamoru-Foundation/mamoru-sniffer-go/generated_bindings"
	"unsafe"
)

func sliceToFfi(bytes []byte) generated_bindings.SliceRefUint8T {
	ptr := C.CBytes(bytes)

	return generated_bindings.SliceRefUint8T{
		Ptr: (*byte)(ptr),
		Len: uint64(len(bytes)),
	}
}

func freeFfiSlice(slice generated_bindings.SliceRefUint8T) {
	C.free(unsafe.Pointer(slice.Ptr))
	slice.Free()
}
