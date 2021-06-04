package lib3ds

// #include <stdlib.h>
// #include <string.h>
// #include "lib3ds.h"
// #cgo linux LDFLAGS: -lstdc++ -lm
import "C"
import "unsafe"

type NodeNative struct {
	m *C.struct_Lib3dsNode
}

func setNode(m *C.struct_Lib3dsNode, n Node) {
	var ret *BaseNode
	switch n.GetType() {
	case NODE_AMBIENT_COLOR:
		node := n.(*AmbientColorNode)
		ret = &node.BaseNode
		cnode := (*C.struct_Lib3dsAmbientColorNode)(unsafe.Pointer(m))
		cnode.color = *((*[3]C.float)(unsafe.Pointer(&node.Color[0])))
		setTrack(&cnode.color_track, node.ColorTrack)
	case NODE_MESH_INSTANCE:
		node := n.(*MeshInstanceNode)
		ret = &node.BaseNode
		cnode := (*C.struct_Lib3dsMeshInstanceNode)(unsafe.Pointer(m))
		cnode.pivot = *((*[3]C.float)(unsafe.Pointer(&node.Pivot[0])))
		cnode.bbox_min = *((*[3]C.float)(unsafe.Pointer(&node.BBoxMin[0])))
		cnode.bbox_max = *((*[3]C.float)(unsafe.Pointer(&node.BBoxMax[0])))
		cnode.hide = C.int(node.Hide)
		cnode.pos = *((*[3]C.float)(unsafe.Pointer(&node.Pos[0])))
		cnode.rot = *((*[4]C.float)(unsafe.Pointer(&node.Rot[0])))
		cnode.scl = *((*[3]C.float)(unsafe.Pointer(&node.Scl[0])))
		cnode.morph_smooth = C.float(node.MorphSmooth)
		setTrack(&cnode.pos_track, node.PosTrack)
		setTrack(&cnode.rot_track, node.RotTrack)
		setTrack(&cnode.scl_track, node.SclTrack)
		setTrack(&cnode.hide_track, node.HideTrack)
		cname := *((*[64]byte)(unsafe.Pointer(&cnode.instance_name)))
		copy(cname[:], []byte(node.InstanceName))
		morphname := *((*[64]byte)(unsafe.Pointer(&cnode.morph)))
		copy(morphname[:], []byte(node.Morph))
	case NODE_CAMERA:
		node := n.(*CameraNode)
		ret = &node.BaseNode
		cnode := (*C.struct_Lib3dsCameraNode)(unsafe.Pointer(m))
		cnode.pos = *((*[3]C.float)(unsafe.Pointer(&node.Pos[0])))
		cnode.fov = C.float(node.Fov)
		cnode.roll = C.float(node.Roll)
		setTrack(&cnode.pos_track, node.PosTrack)
		setTrack(&cnode.fov_track, node.FovTrack)
		setTrack(&cnode.roll_track, node.RollTrack)
	case NODE_OMNILIGHT:
		node := n.(*OmnilightNode)
		ret = &node.BaseNode
		cnode := (*C.struct_Lib3dsOmnilightNode)(unsafe.Pointer(m))
		cnode.pos = *((*[3]C.float)(unsafe.Pointer(&node.Pos[0])))
		setTrack(&cnode.pos_track, node.PosTrack)
		cnode.color = *((*[3]C.float)(unsafe.Pointer(&node.Color[0])))
		setTrack(&cnode.color_track, node.ColorTrack)
	case NODE_SPOTLIGHT:
		node := n.(*SpotlightNode)
		ret = &node.BaseNode
		cnode := (*C.struct_Lib3dsSpotlightNode)(unsafe.Pointer(m))
		cnode.pos = *((*[3]C.float)(unsafe.Pointer(&node.Pos[0])))
		setTrack(&cnode.pos_track, node.PosTrack)
		cnode.color = *((*[3]C.float)(unsafe.Pointer(&node.Color[0])))
		setTrack(&cnode.color_track, node.ColorTrack)
		cnode.hotspot = C.float(node.Hotspot)
		setTrack(&cnode.hotspot_track, node.HotspotTrack)
		cnode.falloff = C.float(node.Falloff)
		setTrack(&cnode.falloff_track, node.FalloffTrack)
		cnode.roll = C.float(node.Roll)
		setTrack(&cnode.roll_track, node.RollTrack)
	case NODE_CAMERA_TARGET:
	case NODE_SPOTLIGHT_TARGET:
		node := n.(*TargetNode)
		ret = &node.BaseNode
		cnode := (*C.struct_Lib3dsTargetNode)(unsafe.Pointer(m))
		cnode.pos = *((*[3]C.float)(unsafe.Pointer(&node.Pos[0])))
		setTrack(&cnode.pos_track, node.PosTrack)
	}
	m.user_id = C.uint(ret.UserID)
	m.user_ptr = unsafe.Pointer(ret.UserPtr)
	m._type = C.Lib3dsNodeType(ret.Type)
	m.node_id = C.ushort(ret.NodeId)
	cname := *((*[64]byte)(unsafe.Pointer(&m.name)))
	copy(cname[:], []byte(ret.Name))
	m.flags = C.uint(ret.Flags)
	m.matrix = *((*[4][4]C.float)(unsafe.Pointer(&ret.Matrix[0])))
}

func getNode(m *C.struct_Lib3dsNode) Node {
	var node Node
	var ret *BaseNode
	switch NodeType(m._type) {
	case NODE_AMBIENT_COLOR:
		node := &AmbientColorNode{}
		ret = &node.BaseNode
		cnode := (*C.struct_Lib3dsAmbientColorNode)(unsafe.Pointer(m))
		node.Color = *((*[3]float32)(unsafe.Pointer(&cnode.color)))
		node.ColorTrack = getTrack(&cnode.color_track)
	case NODE_MESH_INSTANCE:
		node := &MeshInstanceNode{}
		ret = &node.BaseNode
		cnode := (*C.struct_Lib3dsMeshInstanceNode)(unsafe.Pointer(m))
		node.Pivot = *((*[3]float32)(unsafe.Pointer(&cnode.pivot)))
		node.BBoxMin = *((*[3]float32)(unsafe.Pointer(&cnode.bbox_min)))
		node.BBoxMax = *((*[3]float32)(unsafe.Pointer(&cnode.bbox_max)))
		node.Hide = int32(cnode.hide)
		node.Pos = *((*[3]float32)(unsafe.Pointer(&cnode.pos)))
		node.Rot = *((*[4]float32)(unsafe.Pointer(&cnode.rot)))
		node.Scl = *((*[3]float32)(unsafe.Pointer(&cnode.scl)))
		node.MorphSmooth = float32(cnode.morph_smooth)
		node.PosTrack = getTrack(&cnode.pos_track)
		node.RotTrack = getTrack(&cnode.rot_track)
		node.SclTrack = getTrack(&cnode.scl_track)
		node.HideTrack = getTrack(&cnode.hide_track)
		cname := *((*[64]byte)(unsafe.Pointer(&cnode.instance_name)))
		node.InstanceName = string(cname[:])
		morphname := *((*[64]byte)(unsafe.Pointer(&cnode.morph)))
		node.Morph = string(morphname[:])
	case NODE_CAMERA:
		node := &CameraNode{}
		ret = &node.BaseNode
		cnode := (*C.struct_Lib3dsCameraNode)(unsafe.Pointer(m))
		node.Pos = *((*[3]float32)(unsafe.Pointer(&cnode.pos)))
		node.Fov = float32(cnode.fov)
		node.Roll = float32(cnode.roll)
		node.PosTrack = getTrack(&cnode.pos_track)
		node.FovTrack = getTrack(&cnode.fov_track)
		node.RollTrack = getTrack(&cnode.roll_track)
	case NODE_OMNILIGHT:
		node := &OmnilightNode{}
		ret = &node.BaseNode
		cnode := (*C.struct_Lib3dsOmnilightNode)(unsafe.Pointer(m))
		node.Pos = *((*[3]float32)(unsafe.Pointer(&cnode.pos)))
		node.PosTrack = getTrack(&cnode.pos_track)
		node.Color = *((*[3]float32)(unsafe.Pointer(&cnode.color)))
		node.ColorTrack = getTrack(&cnode.color_track)
	case NODE_SPOTLIGHT:
		node := &SpotlightNode{}
		ret = &node.BaseNode
		cnode := (*C.struct_Lib3dsSpotlightNode)(unsafe.Pointer(m))
		node.Pos = *((*[3]float32)(unsafe.Pointer(&cnode.pos)))
		node.PosTrack = getTrack(&cnode.pos_track)
		node.Color = *((*[3]float32)(unsafe.Pointer(&cnode.color)))
		node.ColorTrack = getTrack(&cnode.color_track)
		node.Hotspot = float32(cnode.hotspot)
		node.HotspotTrack = getTrack(&cnode.hotspot_track)
		node.Falloff = float32(cnode.falloff)
		node.FalloffTrack = getTrack(&cnode.falloff_track)
		node.Roll = float32(cnode.roll)
		node.RollTrack = getTrack(&cnode.roll_track)
	case NODE_CAMERA_TARGET:
	case NODE_SPOTLIGHT_TARGET:
		node := &TargetNode{}
		ret = &node.BaseNode
		cnode := (*C.struct_Lib3dsTargetNode)(unsafe.Pointer(m))
		node.Pos = *((*[3]float32)(unsafe.Pointer(&cnode.pos)))
		node.PosTrack = getTrack(&cnode.pos_track)
	}
	ret.UserID = uint32(m.user_id)
	ret.UserPtr = uintptr(m.user_ptr)
	if m.next != nil {
		ret.Next = &NodeNative{m: m.next}
	}
	if m.childs != nil {
		ret.Childs = &NodeNative{m: m.childs}
	}
	if m.parent != nil {
		ret.Parent = &NodeNative{m: m.parent}
	}
	ret.Type = NodeType(m._type)
	ret.NodeId = uint16(m.node_id)
	cname := *((*[64]byte)(unsafe.Pointer(&m.name)))
	ret.Name = string(cname[:])
	ret.Flags = uint32(m.flags)
	ret.Matrix = *((*[4][4]float32)(unsafe.Pointer(&m.matrix)))
	return node
}

func (n *NodeNative) GetNode() Node {
	return getNode(n.m)
}

func (n *NodeNative) SetNode(node Node) {
	setNode(n.m, node)
}

func newNodeNativeByType(tp NodeType) *NodeNative {
	return &NodeNative{m: C.lib3ds_node_new(C.Lib3dsNodeType(tp))}
}

func newNodeNativeByNode(n Node) *NodeNative {
	tp := n.GetType()
	ret := &NodeNative{m: C.lib3ds_node_new(C.Lib3dsNodeType(tp))}
	setNode(ret.m, n)
	return ret
}

type Node interface {
	GetName() string
	GetNodeId() uint16
	GetType() NodeType
	GetNext() Node
	GetChilds() Node
	GetParent() Node
	GetMatrix() *[4][4]float32
}

type BaseNode struct {
	Node
	UserID  uint32
	UserPtr uintptr
	Next    *NodeNative
	Childs  *NodeNative
	Parent  *NodeNative
	Type    NodeType
	NodeId  uint16
	Name    string
	Flags   uint32
	Matrix  [4][4]float32
}

func (n *BaseNode) GetType() NodeType {
	return n.Type
}

func (n *BaseNode) GetNext() Node {
	return n.Next.GetNode()
}

func (n *BaseNode) GetChilds() Node {
	return n.Childs.GetNode()
}

func (n *BaseNode) GetParent() Node {
	return n.Parent.GetNode()
}

func (n *BaseNode) GetNodeId() uint16 {
	return n.NodeId
}

func (n *BaseNode) GetName() string {
	return n.Name
}

func (n *BaseNode) GetMatrix() *[4][4]float32 {
	return &n.Matrix
}

type AmbientColorNode struct {
	BaseNode
	Color      [3]float32
	ColorTrack Track
}

type MeshInstanceNode struct {
	BaseNode
	Pivot        [3]float32
	InstanceName string
	BBoxMin      [3]float32
	BBoxMax      [3]float32
	Hide         int32
	Pos          [3]float32
	Rot          [4]float32
	Scl          [3]float32
	MorphSmooth  float32
	Morph        string
	PosTrack     Track
	RotTrack     Track
	SclTrack     Track
	HideTrack    Track
}

type CameraNode struct {
	BaseNode
	Pos       [3]float32
	Fov       float32
	Roll      float32
	PosTrack  Track
	FovTrack  Track
	RollTrack Track
}

type TargetNode struct {
	BaseNode
	Pos      [3]float32
	PosTrack Track
}

type OmnilightNode struct {
	BaseNode
	Pos        [3]float32
	Color      [3]float32
	PosTrack   Track
	ColorTrack Track
}

type SpotlightNode struct {
	BaseNode
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
