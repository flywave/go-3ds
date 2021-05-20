package lib3ds

type Node struct {
	UserID  uint32
	UserPtr uintptr
	Next    *Node
	Childs  *Node
	Parent  *Node
	Type    NodeType
	NodeId  uint16
	Name    [64]byte
	Flags   uint32
	Matrix  [4][4]float32
}

type AmbientColorNode struct {
	Node
	Color      [3]float32
	ColorTrack Track
}

type MeshInstanceNode struct {
	Node
	Pivot        [3]float32
	InstanceName [64]byte
	BBoxMin      [3]float32
	BBoxMax      [3]float32
	Hide         int32
	Pos          [3]float32
	Rot          [4]float32
	Scl          [3]float32
	MorphSmooth  float32
	Morph        [64]byte
	PosTrack     Track
	RotTrack     Track
	SclTrack     Track
	HideTrack    Track
}

type CameraNode struct {
	Node
	Pos       [3]float32
	Fov       float32
	Roll      float32
	PosTrack  Track
	FovTrack  Track
	RollTrack Track
}

type TargetNode struct {
	Node
	Pos      [3]float32
	PosTrack Track
}

type OmnilightNode struct {
	Node
	Pos        [3]float32
	Color      [3]float32
	PosTrack   Track
	ColorTrack Track
}

type SpotlightNode struct {
	Node
	Pos          [3]float32
	Color        [3]float32
	Hotspot      float32
	Falloff      float32
	Roll         float32
	PosTrack     Track
	ColorTrack   Track
	HotspotTrack Track
	FalloffTrack Track
	RollTrack    Track
}
