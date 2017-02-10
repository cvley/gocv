package transform

import (
	"testing"
)

func TestDct(t *testing.T) {
	m, err := New(4, 4)
	if err != nil {
		t.Fatal(err)
	}

	m.SetRow(2, []float64{1, 2, 3, 4})
	t.Log(m)

	f := Dct2D(m)
	t.Log(f)

	b := IDct2D(f)
	t.Log(b)

	if !m.Equal(b) {
		t.Fatal("dct and idct fail")
	}
}
