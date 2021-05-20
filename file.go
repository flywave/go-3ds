package lib3ds

// #include <stdlib.h>
// #include "lib3ds.h"
// #cgo linux LDFLAGS: -lstdc++ -lm
import "C"
import "unsafe"

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
	UserID            uint32
	UserPtr           uintptr
	MeshVersion       uint32
	KeyfRevision      uint32
	Name              [12 + 1]byte
	MasterScale       float32
	ConstructionPlane [3]float32
	Ambient           [3]float32
	Shadow            Shadow
	Background        Background
	Atmosphere        Atmosphere
	Viewport          Viewport
	ViewportKeyf      Viewport
	Frames            int32
	SegmentFrom       int32
	SegmentTo         int32
	CurrentFrame      int32
	MaterialsSize     int32
	NMaterials        int32
	Materials         **Material
	CamerasSize       int32
	NCameras          int32
	Cameras           **Camera
	LightsSize        int32
	NLights           int32
	Lights            **Light
	MeshesSize        int32
	NMeshes           int32
	Meshes            **Mesh
	Nodes             *Node
}

func (f *File) GetMaterials() []Material {
	return nil
}

func (f *File) GetCameras() []Camera {
	return nil
}

func (f *File) GetLights() []Light {
	return nil
}

func (f *File) GetMeshs() []Mesh {
	return nil
}

func (f *File) ReserveMaterials() {

}

func (f *File) InsertMaterial() {

}

func (f *File) RemoveMaterial() {}

func (f *File) MaterialByName() {}

func (f *File) ReserveCameras() {}

func (f *File) InsertCamera() {}

func (f *File) RemoveCamera() {}

func (f *File) CameraByName() {}

func (f *File) ReserveLights() {}

func (f *File) InsertLight() {}

func (f *File) RemoveLight() {

}

func (f *File) LightByName() {

}

func (f *File) ReserveMeshes() {

}

func (f *File) InsertMesh() {

}

func (f *File) RemoveMesh() {

}

func (f *File) MeshByName() {}

func (f *File) MeshForNode() {}

func (f *File) NodeByName() {}

func (f *File) NodeById() {}

func (f *File) AppendNode() {}

func (f *File) InsertBode() {}

func (f *File) RemoveNode() {}

func (f *File) MinMaxNodeId() {}

func (f *File) CreateNodesForMeshes() {}

func (f *File) BoundingBoxOfObjects() {}

type FileAdapter struct {
	m *C.struct_Lib3dsFile
}

func NewFileAdapter() *FileAdapter {
	return &FileAdapter{m: C.lib3ds_file_new()}
}

func (f *FileAdapter) Open(path string) {
	cpath := C.CString(path)
	defer C.free((unsafe.Pointer)(cpath))
	f.m = C.lib3ds_file_open(cpath)
}

func (f *FileAdapter) Save(path string) {
	cpath := C.CString(path)
	defer C.free((unsafe.Pointer)(cpath))
	C.lib3ds_file_save(f.m, cpath)
}

func (f *FileAdapter) Eval(t float32) {
	C.lib3ds_file_eval(f.m, C.float(t))
}

func (f *FileAdapter) UnSafeGet() *File {
	return (*File)(unsafe.Pointer(&f.m))
}
