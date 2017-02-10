package transform

import (
	"testing"
)

func TestMatrix(t *testing.T) {
	m, err := New(5, 5)
	if err != nil {
		t.Fatal(err)
	}
	if len(m.data) != 25 {
		t.Fatal("New matrix fail")
	}

	m.SetRow(3, []float64{1, 1, 1, 1, 1})
	m.SetCol(2, []float64{2, 2, 2, 2, 2})
	t.Logf("%s", m)
	if m.At(3, 2) != 2 {
		t.Fatal("mismatch value")
	}

	cm := m.Copy()
	if cm.At(3, 2) != 2 {
		t.Fatal("copy matrix fail")
	}
}
