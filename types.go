package lib3ds

type TextureMapFlags int32

const (
	TEXTURE_DECALE       TextureMapFlags = 0x0001
	TEXTURE_MIRROR       TextureMapFlags = 0x0002
	TEXTURE_NEGATE       TextureMapFlags = 0x0008
	TEXTURE_NO_TILE      TextureMapFlags = 0x0010
	TEXTURE_SUMMED_AREA  TextureMapFlags = 0x002
	TEXTURE_ALPHA_SOURCE TextureMapFlags = 0x0040
	TEXTURE_TINT         TextureMapFlags = 0x0080
	TEXTURE_IGNORE_ALPHA TextureMapFlags = 0x0100
	TEXTURE_RGB_TINT     TextureMapFlags = 0x0200
)

type AutoReflMapFlags int32

const (
	AUTOREFL_USE                   AutoReflMapFlags = 0x0001
	AUTOREFL_READ_FIRST_FRAME_ONLY AutoReflMapFlags = 0x0004
	AUTOREFL_FLAT_MIRROR           AutoReflMapFlags = 0x0008
)

type Shading int32

const (
	SHADING_WIRE_FRAME Shading = 0
	SHADING_FLAT       Shading = 1
	SHADING_GOURAUD    Shading = 2
	SHADING_PHONG      Shading = 3
	SHADING_METAL      Shading = 4
)

type ObjectFlags int32

const (
	OBJECT_HIDDEN         ObjectFlags = 0x01
	OBJECT_VIS_LOFTER     ObjectFlags = 0x02
	OBJECT_DOESNT_CAST    ObjectFlags = 0x04
	OBJECT_MATTE          ObjectFlags = 0x08
	OBJECT_DONT_RCVSHADOW ObjectFlags = 0x10
	OBJECT_FAST           ObjectFlags = 0x20
	OBJECT_FROZEN         ObjectFlags = 0x40
)

type MapType int32

const (
	MAP_NONE        MapType = -1
	MAP_PLANAR      MapType = 0
	MAP_CYLINDRICAL MapType = 1
	MAP_SPHERICAL   MapType = 2
)

type FaceFlags int32

const (
	FACE_VIS_AC   FaceFlags = 0x01
	FACE_VIS_BC   FaceFlags = 0x02
	FACE_VIS_AB   FaceFlags = 0x04
	FACE_WRAP_U   FaceFlags = 0x08
	FACE_WRAP_V   FaceFlags = 0x10
	FACE_SELECT_3 FaceFlags = (1 << 13)
	FACE_SELECT_2 FaceFlags = (1 << 14)
	FACE_SELECT_1 FaceFlags = (1 << 15)
)

type NodeType int32

const (
	NODE_AMBIENT_COLOR    NodeType = 0
	NODE_MESH_INSTANCE    NodeType = 1
	NODE_CAMERA           NodeType = 2
	NODE_CAMERA_TARGET    NodeType = 3
	NODE_OMNILIGHT        NodeType = 4
	NODE_SPOTLIGHT        NodeType = 5
	NODE_SPOTLIGHT_TARGET NodeType = 6
)

type NodeFlags int32

const (
	NODE_HIDDEN          NodeFlags = 0x000800
	NODE_SHOW_PATH       NodeFlags = 0x010000
	NODE_SMOOTHING       NodeFlags = 0x020000
	NODE_MOTION_BLUR     NodeFlags = 0x100000
	NODE_MORPH_MATERIALS NodeFlags = 0x400000
)

type KeyFlags int32

const (
	KEY_USE_TENS      KeyFlags = 0x0
	KEY_USE_CONT      KeyFlags = 0x02
	KEY_USE_BIAS      KeyFlags = 0x04
	KEY_USE_EASE_TO   KeyFlags = 0x08
	KEY_USE_EASE_FROM KeyFlags = 0x10
)

type TrackType int32

const (
	TRACK_BOOL   TrackType = 0
	TRACK_FLOAT  TrackType = 1
	TRACK_VECTOR TrackType = 3
	TRACK_QUAT   TrackType = 4
)

type TrackFlags int32

const (
	TRACK_REPEAT   TrackFlags = 0x0001
	TRACK_SMOOTH   TrackFlags = 0x0002
	TRACK_LOCK_X   TrackFlags = 0x0008
	TRACK_LOCK_Y   TrackFlags = 0x0010
	TRACK_LOCK_Z   TrackFlags = 0x0020
	TRACK_UNLINK_X TrackFlags = 0x0100
	TRACK_UNLINK_Y TrackFlags = 0x0200
	TRACK_UNLINK_Z TrackFlags = 0x0400
)
