package transform

import (
	"testing"
)

func TestArnold(t *testing.T) {
	data, _ := New(2, 2)
	data.Set(1, 1, 1)

	Arnold(data)

	t.Log(data)
}
