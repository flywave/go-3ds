package lib3ds

type Light struct {
	UserID          uint32
	UserPtr         uintptr
	Name            [64]byte
	ObjectFlags     uint32
	SpotLight       int32
	SeeCone         int32
	Color           [3]float32
	Position        [3]float32
	Target          [3]float32
	Roll            float32
	Off             int32
	OuterRange      float32
	InnerRange      float32
	Multiplier      float32
	Attenuation     float32
	RectangularSpot int32
	Shadowed        int32
	ShadowBias      float32
	ShadowFilter    float32
	ShadowSize      int32
	SpotAspect      float32
	UseProjector    int32
	Projector       [64]byte
	SpotOvershoot   int32
	RayShadows      int32
	RayBias         float32
	Hotspot         float32
	Falloff         float32
}
