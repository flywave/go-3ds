package lib3ds

type TextureMap struct {
	UserID   uint32
	UserPtr  uintptr
	Name     [64]byte
	Flags    uint32
	Percent  float32
	Blur     float32
	Scale    [2]float32
	Offset   [2]float32
	Rotation float32
	Tint_1   [3]float32
	Tint_2   [3]float32
	Tint_r   [3]float32
	Tint_g   [3]float32
	Tint_b   [3]float32
}

type Material struct {
	UserID               uint32
	UserPtr              uintptr
	Name                 [64]byte
	Ambient              [3]float32
	Diffuse              [3]float32
	Specular             [3]float32
	Shininess            float32
	ShinStrength         float32
	UseBlur              int32
	Blur                 float32
	Transparency         float32
	Falloff              float32
	IsAdditive           int32
	SelfIllumFlag        int32
	SelfIllum            float32
	UseFalloff           int32
	Shading              int32
	Soften               int32
	FaceMap              int32
	TwoSided             int32
	MapDecal             int32
	UseWire              int32
	UseWireAbs           int32
	WireSize             float32
	Texture1Map          TextureMap
	Texture1Mask         TextureMap
	Texture2Map          TextureMap
	Texture2Mask         TextureMap
	OpacityMap           TextureMap
	OpacityMask          TextureMap
	BumpMap              TextureMap
	BumpMask             TextureMap
	SpecularMap          TextureMap
	SpecularMask         TextureMap
	ShininessMap         TextureMap
	ShininessMask        TextureMap
	SelfIllumMap         TextureMap
	SelfIllumMask        TextureMap
	ReflectionMap        TextureMap
	ReflectionMask       TextureMap
	AutoreflMapFlags     uint32
	AutoreflMapAntiAlias int32
	AutoreflMapSize      int32
	AutoreflMapFrameStep int32
}
