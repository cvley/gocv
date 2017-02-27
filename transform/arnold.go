package transform

import (
	"github.com/cvley/gocv/core"
)

func Arnold(data *core.Mat) {
	if data.Rows() != data.Cols() {
		panic("Arnold fail: rows != cols")
	}

	N := data.Rows()
	result := core.NewMat(N, N)
	for h := 0; h < N; h++ {
		for w := 0; w < N; w++ {
			x := (h + w) % N
			y := (2*h + w) % N
			result.Set(y, x, data.At(h, w))
		}
	}

	for h := 0; h < N; h++ {
		for w := 0; w < N; w++ {
			data.Set(h, w, result.At(h, w))
		}
	}
}
