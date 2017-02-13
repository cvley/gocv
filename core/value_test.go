package core

import (
	"testing"
)

func TestInt(t *testing.T) {
	a := Int(1)
	t.Log(a.Int64())
	t.Log(a.Float64())
}
