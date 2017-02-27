package transform

import (
	"testing"

	"github.com/cvley/gocv/core"
)

func TestArnold(t *testing.T) {
	data := core.NewMat(2, 2)
	data.Set(1, 1, core.NewValue(1, 1))

	Arnold(data)

	t.Log(data)
}
