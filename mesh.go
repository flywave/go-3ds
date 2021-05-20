package lib3ds

type Face struct {
	Index          [3]uint16
	Flags          uint16
	Material       int32
	SmoothingGroup uint32
}

type Mesh struct {
	UserID            uint32
	UserPtr           uintptr
	Name              [64]byte
	ObjectFlags       uint32
	Color             int32
	Matrix            [4][4]float32
	NVertices         uint16
	Vertices          *[3]float32
	Texcos            *[2]float32
	VFlags            *uint16
	Nfaces            uint16
	Faces             *Face
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
