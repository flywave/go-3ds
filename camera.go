package lib3ds

type Camera struct {
	UserID      uint32
	UserPtr     uintptr
	Name        [64]byte
	ObjectFlags uint32
	Position    [3]float32
	Target      [3]float32
	Roll        float32
	Fov         float32
	SeeCone     int32
	NearRange   float32
	FarRange    float32
}
