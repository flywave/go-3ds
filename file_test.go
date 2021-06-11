package lib3ds

import "testing"

func TestReader(t *testing.T) {
	f := OpenFile("./testdata/BieShu.3DS")
	meshs := f.GetMeshs()
	if len(meshs) == 0 {
		t.FailNow()
	}
}
