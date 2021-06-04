package lib3ds

// #include <stdlib.h>
// #include <string.h>
// #include "lib3ds.h"
// #cgo linux LDFLAGS: -lstdc++ -lm
import "C"
import (
	"errors"
	"reflect"
	"runtime"
	"unsafe"
)

type Shadow struct {
	MapSize uint16
	LowBias float32
	HiBias  float32
	Filter  float32
	RayBias float32
}

type Background struct {
	UseBitmap       int32
	BitmapName      [64]byte
	UseSolid        int32
	SolidColor      [3]float32
	UseGradient     int32
	GradientPercent float32
	GradientTop     [3]float32
	GradientMiddle  [3]float32
	GradientBottom  [3]float32
}

type Atmosphere struct {
	UseFog             int32
	FogColor           [3]float32
	FogBackground      int32
	FogNearPlane       float32
	FogNearDensity     float32
	FogFarPlane        float32
	FogFarDensity      float32
	UseLayerFog        int32
	LayerFogFlags      uint32
	LayerFogColor      [3]float32
	LayerFogNearY      float32
	LayerFogFarY       float32
	LayerFogDensity    float32
	UseDistCue         int32
	DistCueBackground  int32
	DistCueNearPlane   float32
	DistCueNearDimming float32
	DistCueFarPlane    float32
	DistCueFarDimming  float32
}

type View struct {
	Type       int32
	AxisLock   uint32
	Position   [2]int16
	Size       [2]int16
	Zoom       float32
	Center     [3]float32
	HorizAngle float32
	VertAngle  float32
	Camera     [11]byte
}

type Viewport struct {
	LayoutStyle       int32
	LayoutActive      int32
	LayoutSwap        int32
	LayoutSwapPrior   int32
	LayoutSwapView    int32
	LayoutPosition    [2]uint16
	LayoutSize        [2]uint16
	LayoutNViews      int32
	LayoutViews       [32]View
	DefaultType       int32
	DefaultPosition   [3]float32
	DefaultWidth      float32
	DefaultHorizAngle float32
	DefaultVertAngle  float32
	DefaultRollAngle  float32
	DefaultCamera     [64]byte
}

type File struct {
	m *C.struct_Lib3dsFile
}

func OpenFile(path string) *File {
	ret := new(File)
	ret.open(path)
	runtime.SetFinalizer(ret, (*File).free)
	return ret
}

func NewFile(path string) *File {
	ret := new(File)
	ret.m = C.lib3ds_file_new()
	runtime.SetFinalizer(ret, (*File).free)
	return ret
}

func (f *File) free() {
	C.lib3ds_file_free(f.m)
}

func (f *File) open(path string) {
	cpath := C.CString(path)
	defer C.free((unsafe.Pointer)(cpath))
	f.m = C.lib3ds_file_open(cpath)
}

func (f *File) Save(path string) {
	cpath := C.CString(path)
	defer C.free((unsafe.Pointer)(cpath))
	C.lib3ds_file_save(f.m, cpath)
}

func (f *File) Eval(t float32) {
	C.lib3ds_file_eval(f.m, C.float(t))
}

func (f *File) GetUserID() uint32 {
	return uint32(f.m.user_id)
}

func (f *File) GetUserPtr() uintptr {
	return uintptr(f.m.user_ptr)
}

func (f *File) GetMeshVersion() uint32 {
	return uint32(f.m.mesh_version)
}

func (f *File) GetKeyfRevision() uint32 {
	return uint32(f.m.keyf_revision)
}

func (f *File) GetName() string {
	cname := *((*[13]byte)(unsafe.Pointer(&f.m.name)))
	return string(cname[:])
}

func (f *File) GetMasterScale() float32 {
	return float32(f.m.master_scale)
}

func (f *File) GetConstructionPlane() [3]float32 {
	return *(*[3]float32)(unsafe.Pointer(&f.m.construction_plane))
}

func (f *File) GetAmbient() [3]float32 {
	return *(*[3]float32)(unsafe.Pointer(&f.m.ambient))
}

func (f *File) GetShadow() Shadow {
	return *(*Shadow)(unsafe.Pointer(&f.m.shadow))
}

func (f *File) GetBackground() Background {
	return *(*Background)(unsafe.Pointer(&f.m.background))
}

func (f *File) GetAtmosphere() Atmosphere {
	return *(*Atmosphere)(unsafe.Pointer(&f.m.atmosphere))
}

func (f *File) GetViewport() Viewport {
	return *(*Viewport)(unsafe.Pointer(&f.m.viewport))
}

func (f *File) GetViewportKeyf() Viewport {
	return *(*Viewport)(unsafe.Pointer(&f.m.viewport_keyf))
}

func (f *File) GetFrames() int32 {
	return int32(f.m.frames)
}

func (f *File) GetSegmentFrom() int32 {
	return int32(f.m.segment_from)
}

func (f *File) GetSegmentTo() int32 {
	return int32(f.m.segment_to)
}

func (f *File) GetCurrentFrame() int32 {
	return int32(f.m.current_frame)
}

func (f *File) GetMaterialsSize() int32 {
	return int32(f.m.materials_size)
}

func (f *File) GetMaterialsCount() int32 {
	return int32(f.m.nmaterials)
}

func (f *File) GetMaterials() []Material {
	ret := make([]Material, int(f.m.nmaterials))

	var mtlsSlice []*C.struct_Lib3dsMaterial
	mtlsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&mtlsSlice)))
	mtlsHeader.Cap = int(f.m.nmaterials)
	mtlsHeader.Len = int(f.m.nmaterials)
	mtlsHeader.Data = uintptr(unsafe.Pointer(f.m.materials))

	for i := range ret {
		ret[i] = *(*Material)(unsafe.Pointer(mtlsSlice[i]))
	}
	return ret
}

func (f *File) GetCamerasSize() int32 {
	return int32(f.m.cameras_size)
}

func (f *File) GetCamerasCount() int32 {
	return int32(f.m.ncameras)
}

func (f *File) GetCameras() []Camera {
	ret := make([]Camera, int(f.m.ncameras))

	var camsSlice []*C.struct_Lib3dsCamera
	camsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&camsSlice)))
	camsHeader.Cap = int(f.m.ncameras)
	camsHeader.Len = int(f.m.ncameras)
	camsHeader.Data = uintptr(unsafe.Pointer(f.m.materials))

	for i := range ret {
		ret[i] = *(*Camera)(unsafe.Pointer(camsSlice[i]))
	}
	return ret
}

func (f *File) GetLightsSize() int32 {
	return int32(f.m.lights_size)
}

func (f *File) GetLightsCount() int32 {
	return int32(f.m.nlights)
}

func (f *File) GetLights() []Light {
	ret := make([]Light, int(f.m.nlights))

	var lightsSlice []*C.struct_Lib3dsCamera
	lightsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&lightsSlice)))
	lightsHeader.Cap = int(f.m.nlights)
	lightsHeader.Len = int(f.m.nlights)
	lightsHeader.Data = uintptr(unsafe.Pointer(f.m.materials))

	for i := range ret {
		ret[i] = *(*Light)(unsafe.Pointer(lightsSlice[i]))
	}
	return ret
}

func (f *File) GetMeshsSize() int32 {
	return int32(f.m.meshes_size)
}

func (f *File) GetMeshsCount() int32 {
	return int32(f.m.nmeshes)
}

func (f *File) GetMeshs() []Mesh {
	ret := make([]Mesh, int(f.m.nmeshes))

	var meshSlice []*C.struct_Lib3dsMesh
	meshHeader := (*reflect.SliceHeader)((unsafe.Pointer(&meshSlice)))
	meshHeader.Cap = int(f.m.nmeshes)
	meshHeader.Len = int(f.m.nmeshes)
	meshHeader.Data = uintptr(unsafe.Pointer(f.m.meshes))

	for i := range ret {
		ret[i].setMesh(meshSlice[i])
	}
	return ret
}

func (f *File) ReserveMaterials(si int, force bool) {
	if force {
		C.lib3ds_file_reserve_materials(f.m, C.int(si), C.int(1))
	} else {
		C.lib3ds_file_reserve_materials(f.m, C.int(si), C.int(0))
	}
}

func (f *File) InsertMaterial(i int, mtl Material) error {
	if i >= int(f.m.nmaterials) {
		return errors.New("index error")
	}
	cname := C.CString(string(mtl.Name[:]))
	defer C.free(unsafe.Pointer(cname))

	cmtl := C.lib3ds_material_new(cname)
	*cmtl = *(*C.struct_Lib3dsMaterial)(unsafe.Pointer(&mtl))
	C.lib3ds_file_insert_material(f.m, cmtl, C.int(i))
	return nil
}

func (f *File) RemoveMaterial(i int) {
	C.lib3ds_file_remove_material(f.m, C.int(i))
}

func (f *File) GetMaterialByName(name string) Material {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	i := int(C.lib3ds_file_material_by_name(f.m, cname))
	return f.GetMaterial(i)
}

func (f *File) GetMaterial(i int) Material {
	var mtlsSlice []*C.struct_Lib3dsMaterial
	mtlsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&mtlsSlice)))
	mtlsHeader.Cap = int(f.m.nmaterials)
	mtlsHeader.Len = int(f.m.nmaterials)
	mtlsHeader.Data = uintptr(unsafe.Pointer(f.m.materials))

	return *(*Material)(unsafe.Pointer(mtlsSlice[i]))
}

func (f *File) ReserveCameras(si int, force bool) {
	if force {
		C.lib3ds_file_reserve_cameras(f.m, C.int(si), C.int(1))
	} else {
		C.lib3ds_file_reserve_cameras(f.m, C.int(si), C.int(0))
	}
}

func (f *File) InsertCamera(i int, cam Camera) error {
	if i >= int(f.m.ncameras) {
		return errors.New("index error")
	}
	cname := C.CString(string(cam.Name[:]))
	defer C.free(unsafe.Pointer(cname))
	ccam := C.lib3ds_camera_new(cname)
	*ccam = *(*C.struct_Lib3dsCamera)(unsafe.Pointer(&cam))
	C.lib3ds_file_insert_camera(f.m, ccam, C.int(i))
	return nil
}

func (f *File) RemoveCamera(i int) {
	C.lib3ds_file_remove_camera(f.m, C.int(i))
}

func (f *File) GetCameraByName(name string) Camera {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	i := int(C.lib3ds_file_camera_by_name(f.m, cname))
	return f.GetCamera(i)
}

func (f *File) GetCamera(i int) Camera {
	var camsSlice []*C.Lib3dsCamera
	camsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&camsSlice)))
	camsHeader.Cap = int(f.m.ncameras)
	camsHeader.Len = int(f.m.ncameras)
	camsHeader.Data = uintptr(unsafe.Pointer(f.m.cameras))

	return *(*Camera)(unsafe.Pointer(camsSlice[i]))
}

func (f *File) ReserveLights(si int, force bool) {
	if force {
		C.lib3ds_file_reserve_lights(f.m, C.int(si), C.int(1))
	} else {
		C.lib3ds_file_reserve_lights(f.m, C.int(si), C.int(0))
	}
}

func (f *File) InsertLight(i int, light Light) error {
	if i >= int(f.m.nlights) {
		return errors.New("index error")
	}
	cname := C.CString(string(light.Name[:]))
	defer C.free(unsafe.Pointer(cname))
	clight := C.lib3ds_light_new(cname)
	*clight = *(*C.struct_Lib3dsLight)(unsafe.Pointer(&light))
	C.lib3ds_file_insert_light(f.m, clight, C.int(i))
	return nil
}

func (f *File) RemoveLight(i int) {
	C.lib3ds_file_remove_light(f.m, C.int(i))
}

func (f *File) GetLightByName(name string) Light {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	i := int(C.lib3ds_file_light_by_name(f.m, cname))
	return f.GetLight(i)
}

func (f *File) GetLight(i int) Light {
	var lightsSlice []*C.Lib3dsLight
	lightsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&lightsSlice)))
	lightsHeader.Cap = int(f.m.nlights)
	lightsHeader.Len = int(f.m.nlights)
	lightsHeader.Data = uintptr(unsafe.Pointer(f.m.lights))

	return *(*Light)(unsafe.Pointer(lightsSlice[i]))
}

func (f *File) ReserveMeshes(si int, force bool) {
	if force {
		C.lib3ds_file_reserve_meshes(f.m, C.int(si), C.int(1))
	} else {
		C.lib3ds_file_reserve_meshes(f.m, C.int(si), C.int(0))
	}
}

func (f *File) InsertMesh(i int, mesh Mesh) error {
	if i >= int(f.m.nmeshes) {
		return errors.New("index error")
	}
	cname := C.CString(mesh.Name)
	defer C.free(unsafe.Pointer(cname))
	cmesh := C.lib3ds_mesh_new(cname)
	mesh.getMesh(cmesh)
	C.lib3ds_file_insert_mesh(f.m, cmesh, C.int(i))
	return nil
}

func (f *File) RemoveMesh(i int) {
	C.lib3ds_file_remove_mesh(f.m, C.int(i))
}

func (f *File) GetMeshByName(name string) Mesh {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	i := int(C.lib3ds_file_mesh_by_name(f.m, cname))
	return f.GetMesh(i)
}

func (f *File) GetMesh(i int) Mesh {
	var meshsSlice []*C.Lib3dsMesh
	meshsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&meshsSlice)))
	meshsHeader.Cap = int(f.m.nmeshes)
	meshsHeader.Len = int(f.m.nmeshes)
	meshsHeader.Data = uintptr(unsafe.Pointer(f.m.meshes))

	m := Mesh{}
	m.setMesh(meshsSlice[i])
	return m
}

func (f *File) GetMeshForNode(n *NodeNative) Mesh {
	cm := C.lib3ds_file_mesh_for_node(f.m, n.m)

	m := Mesh{}
	m.setMesh(cm)
	return m
}

func (f *File) GetNodeByName(name string, tp NodeType) *NodeNative {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cm := C.lib3ds_file_node_by_name(f.m, cname, C.Lib3dsNodeType(tp))

	return &NodeNative{m: cm}
}

func (f *File) GetNodeById(node_id uint16) *NodeNative {
	cm := C.lib3ds_file_node_by_id(f.m, C.ushort(node_id))

	return &NodeNative{m: cm}
}

func (f *File) AppendNode(n *NodeNative, p *NodeNative) {
	C.lib3ds_file_append_node(f.m, n.m, p.m)
}

func (f *File) InsertNode(n *NodeNative, at *NodeNative) {
	C.lib3ds_file_insert_node(f.m, n.m, at.m)
}

func (f *File) RemoveNode(n *NodeNative) {
	C.lib3ds_file_remove_node(f.m, n.m)
}

func (f *File) GetMinMaxNodeId() (min, max uint16) {
	var cmin, cmax C.ushort
	C.lib3ds_file_minmax_node_id(f.m, &cmin, &cmax)
	min = uint16(cmin)
	max = uint16(cmax)
	return
}

func (f *File) CreateNodesForMeshes() {
	C.lib3ds_file_create_nodes_for_meshes(f.m)
}
