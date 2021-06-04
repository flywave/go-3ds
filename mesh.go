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

type Face struct {
	Index          [3]uint16
	Flags          uint16
	Material       int32
	SmoothingGroup uint32
}

type Mesh struct {
	UserID            uint32
	UserPtr           uintptr
	Name              string
	ObjectFlags       uint32
	Color             int32
	Matrix            [4][4]float32
	Vertices          [][3]float32
	Texcos            [][2]float32
	VFlags            []uint16
	Faces             []Face
	BoxFront          [64]byte
	BoxBack           [64]byte
	BoxLeft           [64]byte
	BoxRight          [64]byte
	BoxTop            [64]byte
	BoxBottom         [64]byte
	MapType           int32
	MapPos            [3]float32
	MapMatrix         [4][4]float32
	MapScale          float32
	MapTile           [2]float32
	MapPlanarSize     [2]float32
	MapCylinderHeight float32
}

func (m *Mesh) getMesh(cm *C.struct_Lib3dsMesh) {
	cm.user_id = C.uint(m.UserID)
	cm.user_ptr = unsafe.Pointer(m.UserPtr)
	namebytes := []byte(m.Name)
	cm.name = *((*[64]C.char)(unsafe.Pointer(&namebytes[0])))
	cm.object_flags = C.uint(m.ObjectFlags)
	cm.color = C.int(m.Color)
	cm.matrix = *((*[4][4]C.float)(unsafe.Pointer(&m.Matrix[0][0])))
	cm.box_front = *((*[64]C.char)(unsafe.Pointer(&m.BoxFront[0])))
	cm.box_back = *((*[64]C.char)(unsafe.Pointer(&m.BoxBack[0])))
	cm.box_left = *((*[64]C.char)(unsafe.Pointer(&m.BoxLeft[0])))
	cm.box_right = *((*[64]C.char)(unsafe.Pointer(&m.BoxRight[0])))
	cm.box_top = *((*[64]C.char)(unsafe.Pointer(&m.BoxTop[0])))
	cm.box_bottom = *((*[64]C.char)(unsafe.Pointer(&m.BoxFront[0])))
	cm.map_type = C.int(m.MapType)
	cm.map_pos = *((*[3]C.float)(unsafe.Pointer(&m.MapPos[0])))
	cm.map_matrix = *((*[4][4]C.float)(unsafe.Pointer(&m.MapMatrix[0][0])))
	cm.map_scale = C.float(m.MapScale)
	cm.map_tile = *((*[2]C.float)(unsafe.Pointer(&m.MapTile[0])))
	cm.map_planar_size = *((*[2]C.float)(unsafe.Pointer(&m.MapPlanarSize[0])))
	cm.map_cylinder_height = C.float(m.MapCylinderHeight)

	use_texcos := 0
	if len(m.Texcos) > 0 {
		use_texcos = 1
	}

	use_flags := 0
	if len(m.VFlags) > 0 {
		use_flags = 1
	}
	n := len(m.Vertices)
	C.lib3ds_mesh_resize_vertices(cm, C.int(n), C.int(use_texcos), C.int(use_flags))

	var vertsSlice [][3]float32
	vertsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&vertsSlice)))
	vertsHeader.Cap = n
	vertsHeader.Len = n
	vertsHeader.Data = uintptr(unsafe.Pointer(cm.vertices))

	copy(vertsSlice, m.Vertices)

	if use_texcos > 0 {
		var texsSlice [][2]float32
		texsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&texsSlice)))
		texsHeader.Cap = n
		texsHeader.Len = n
		texsHeader.Data = uintptr(unsafe.Pointer(cm.texcos))

		copy(texsSlice, m.Texcos)
	}

	if use_flags > 0 {
		var flagsSlice []uint16
		flagsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&flagsSlice)))
		flagsHeader.Cap = n
		flagsHeader.Len = n
		flagsHeader.Data = uintptr(unsafe.Pointer(cm.vflags))

		copy(flagsSlice, m.VFlags)
	}

	C.lib3ds_mesh_resize_faces(cm, C.int(len(m.Texcos)))

	nface := len(m.Faces)

	var facesSlice []Face
	facesHeader := (*reflect.SliceHeader)((unsafe.Pointer(&facesSlice)))
	facesHeader.Cap = nface
	facesHeader.Len = nface
	facesHeader.Data = uintptr(unsafe.Pointer(cm.faces))

	copy(facesSlice, m.Faces)
}

func (m *Mesh) setMesh(cm *C.struct_Lib3dsMesh) {
	m.UserID = uint32(cm.user_id)
	m.UserPtr = uintptr(cm.user_ptr)
	cname := *((*[64]byte)(unsafe.Pointer(&cm.name)))
	m.Name = string(cname[:])
	m.ObjectFlags = uint32(cm.object_flags)
	m.Color = int32(cm.color)
	m.Matrix = *((*[4][4]float32)(unsafe.Pointer(&cm.matrix)))
	m.BoxFront = *((*[64]byte)(unsafe.Pointer(&cm.box_front)))
	m.BoxBack = *((*[64]byte)(unsafe.Pointer(&cm.box_back)))
	m.BoxLeft = *((*[64]byte)(unsafe.Pointer(&cm.box_left)))
	m.BoxRight = *((*[64]byte)(unsafe.Pointer(&cm.box_right)))
	m.BoxTop = *((*[64]byte)(unsafe.Pointer(&cm.box_top)))
	m.BoxBottom = *((*[64]byte)(unsafe.Pointer(&cm.box_bottom)))
	m.MapType = int32(cm.map_type)
	m.MapPos = *((*[3]float32)(unsafe.Pointer(&cm.map_pos)))
	m.MapMatrix = *((*[4][4]float32)(unsafe.Pointer(&cm.map_matrix)))
	m.MapScale = float32(cm.map_scale)
	m.MapTile = *((*[2]float32)(unsafe.Pointer(&cm.map_tile)))
	m.MapPlanarSize = *((*[2]float32)(unsafe.Pointer(&cm.map_planar_size)))
	m.MapCylinderHeight = float32(cm.map_cylinder_height)

	m.setVertices(cm.vertices, int(cm.nvertices))
	m.setTexCoords(cm.texcos, int(cm.nvertices))
	m.setVFlags(cm.vflags, int(cm.nvertices))

	m.setFaces(cm.faces, int(cm.nfaces))
}

func (m *Mesh) setVertices(verts *[3]C.float, n int) {
	if verts == nil {
		return
	}
	var vertsSlice [][3]float32
	vertsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&vertsSlice)))
	vertsHeader.Cap = n
	vertsHeader.Len = n
	vertsHeader.Data = uintptr(unsafe.Pointer(verts))

	m.Vertices = make([][3]float32, n)
	copy(m.Vertices, vertsSlice)
}

func (m *Mesh) setTexCoords(texs *[2]C.float, n int) {
	if texs == nil {
		return
	}
	var texsSlice [][2]float32
	texsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&texsSlice)))
	texsHeader.Cap = n
	texsHeader.Len = n
	texsHeader.Data = uintptr(unsafe.Pointer(texs))

	m.Texcos = make([][2]float32, n)
	copy(m.Texcos, texsSlice)
}

func (m *Mesh) setFaces(faces *C.struct_Lib3dsFace, n int) {
	if faces == nil {
		return
	}
	var facesSlice []Face
	facesHeader := (*reflect.SliceHeader)((unsafe.Pointer(&facesSlice)))
	facesHeader.Cap = n
	facesHeader.Len = n
	facesHeader.Data = uintptr(unsafe.Pointer(faces))

	m.Faces = make([]Face, n)
	copy(m.Faces, facesSlice)
}

func (m *Mesh) setVFlags(vflags *C.ushort, n int) {
	if vflags == nil {
		return
	}
	var flagsSlice []uint16
	flagsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&flagsSlice)))
	flagsHeader.Cap = n
	flagsHeader.Len = n
	flagsHeader.Data = uintptr(unsafe.Pointer(vflags))

	m.VFlags = make([]uint16, n)
	copy(m.VFlags, flagsSlice)
}
