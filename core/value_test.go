package core

import (
	"testing"
)

func TestValue(t *testing.T) {
	v := int(1)
	value := NewValue(1, v)
	t.Log(value)
}
