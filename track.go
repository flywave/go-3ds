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
	if m == nil {
		return
	}
	if int(m.nkeys) != len(t.Keys) {
		C.lib3ds_track_resize(m, C.int(len(t.Keys)))
	}
	m.flags = C.uint(t.Flags)
	m._type = C.Lib3dsTrackType(t.Type)
	m.nkeys = C.int(len(t.Keys))

	var keysSlice []Key
	facesHeader := (*reflect.SliceHeader)((unsafe.Pointer(&keysSlice)))
	facesHeader.Cap = int(m.nkeys)
	facesHeader.Len = int(m.nkeys)
	facesHeader.Data = uintptr(unsafe.Pointer(m.keys))

	copy(keysSlice, t.Keys)
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
