package transform

import (
	"testing"
)

func TestHaar(t *testing.T) {
	m, err := New(4, 4)
	if err != nil {
		t.Fatal(err)
	}

	m.SetRow(2, []float64{1, 2, 3, 4})
	t.Log(m)

	f := HaarForward2D(m, 1)
	t.Log(f)

	b := HaarBackward2D(f, 1)
	t.Log(b)
}
