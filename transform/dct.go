package transform

import (
	_ "log"
	"math"
)

func Dct1D(data []float64) []float64 {
	result := make([]float64, len(data))
	c := math.Pi / (2.0 * float64(len(data)))
	scale := math.Sqrt(2.0 / float64(len(data)))

	for k := 0; k < len(data); k++ {
		var sum float64
		for n := 0; n < len(data); n++ {
			sum += data[n] * math.Cos((2.0*float64(n)+1.0)*float64(k)*c)
		}
		result[k] = scale * sum
	}

	result[0] = result[0] / math.Sqrt2
	return result
}

func IDct1D(data []float64) []float64 {
	result := make([]float64, len(data))
	c := math.Pi / (2.0 * float64(len(data)))
	scale := math.Sqrt(2.0 / float64(len(data)))

	for k := 0; k < len(data); k++ {
		sum := data[0] / math.Sqrt2
		for n := 1; n < len(data); n++ {
			sum += data[n] * math.Cos((2*float64(k)+1)*float64(n)*c)
		}
		result[k] = scale * sum
	}

	return result
}

func Dct2D(data *Matrix) *Matrix {
	rows := data.Rows()
	cols := data.Cols()

	m := data.Copy()

	for i := 0; i < rows; i++ {
		rowData := m.GetRow(i)
		m.SetRow(i, Dct1D(rowData))
	}

	for i := 0; i < cols; i++ {
		colData := m.GetCol(i)
		m.SetCol(i, Dct1D(colData))
	}

	return m
}

func IDct2D(data *Matrix) *Matrix {
	m := data.Copy()

	for i := 0; i < data.Cols(); i++ {
		colData := m.GetCol(i)
		m.SetCol(i, IDct1D(colData))
	}
	for i := 0; i < data.Rows(); i++ {
		rowData := m.GetRow(i)
		m.SetRow(i, IDct1D(rowData))
	}

	return m
}
