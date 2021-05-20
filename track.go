package lib3ds

type Key struct {
	Frame    int32
	Flags    uint32
	Tens     float32
	Cont     float32
	Bias     float32
	EaseTo   float32
	EaseFrom float32
	Value    [4]float32
}

type Track struct {
	Flags uint
	Type  TrackType
	NKeys int32
	Keys  *Key
}
