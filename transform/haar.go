package transform

import (
	"log"
)

const (
	w0 = 0.5
	w1 = -0.5
	s0 = 0.5
	s1 = 0.5
)

func HaarForward1D(data []float64) []float64 {
	if len(data)%2 == 1 {
		panic("haar transform only support even length data")
	}

	result := make([]float64, len(data))

	h := len(data) >> 1
	for i := 0; i < h; i++ {
		k := i << 1
		result[i] = data[k]*s0 + data[k+1]*s1
		result[i+h] = data[k]*w0 + data[k+1]*w1
	}

	return result
}

func HaarBackward1D(data []float64) []float64 {
	if len(data)%2 == 1 {
		panic("haar transform only support even length data")
	}

	result := make([]float64, len(data))
	h := len(data) >> 1
	for i := 0; i < h; i++ {
		k := i << 1
		result[k] = (data[i]*s0 + data[i+h]*w0) / w0
		result[k+1] = (data[i]*s1 + data[i+h]*w1) / s0
	}

	return result
}

func HaarForward2D(data *Matrix, level int) *Matrix {
	rows := data.Rows()
	cols := data.Cols()

	m := data.Copy()

	for k := 0; k < level; k++ {
		l := 1 << uint(k)
		lCols := cols / l
		lRows := rows / l

		// TODO parallel
		for i := 0; i < lRows; i++ {
			rowData := m.GetRow(i)
			m.SetRow(i, HaarForward1D(rowData))
		}

		for i := 0; i < lCols; i++ {
			colData := m.GetCol(i)
			m.SetCol(i, HaarForward1D(colData))
		}
	}

	return m
}

func HaarBackward2D(data *Matrix, level int) *Matrix {
	rows := data.Rows()
	cols := data.Cols()

	m := data.Copy()

	for k := level - 1; k >= 0; k-- {
		l := 1 << uint(k)
		log.Println(l)
		lCols := cols / l
		lRows := rows / l

		// TODO parallel
		for i := 0; i < lCols; i++ {
			colData := m.GetCol(i)
			m.SetCol(i, HaarBackward1D(colData))
		}

		for i := 0; i < lRows; i++ {
			rowData := m.GetRow(i)
			m.SetRow(i, HaarBackward1D(rowData))
		}
	}

	return m
}
