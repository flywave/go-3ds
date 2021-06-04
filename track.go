package lib3ds

// #include <stdlib.h>
// #include <string.h>
// #include "lib3ds.h"
// #cgo linux LDFLAGS: -lstdc++ -lm
import "C"
import (
	"reflect"
	"unsafe"
)

type Key struct {
	Frame    int32
	Flags    uint32
	Tens     float32
	Cont     float32
	Bias     float32
	EaseTo   float32
	EaseFrom float32
	Value    [4]float32
}

type Track struct {
	Flags uint
	Type  TrackType
	Keys  []Key
}

func setTrack(m *C.struct_Lib3dsTrack, t Track) {
}

func getTrack(m *C.struct_Lib3dsTrack) Track {
	ret := Track{}
	ret.Flags = uint(m.flags)
	ret.Type = TrackType(m._type)
	ret.Keys = make([]Key, int(m.nkeys))

	var keysSlice []Key
	facesHeader := (*reflect.SliceHeader)((unsafe.Pointer(&keysSlice)))
	facesHeader.Cap = int(m.nkeys)
	facesHeader.Len = int(m.nkeys)
	facesHeader.Data = uintptr(unsafe.Pointer(m.keys))

	copy(ret.Keys, keysSlice)
	return ret
}
