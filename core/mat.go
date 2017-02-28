package core

import (
	"errors"
	"image"
	"os"
	"strings"
)

type Mat struct {
	data []Value
	rows int
	cols int
}

func NewFromImage(img image.Image) (*Mat, error) {
	if img == nil {
		return nil, errors.New("NewFromImage fail: invalid input image")
	}

	cols := img.Bounds().Dx()
	rows := img.Bounds().Dy()
	mat := NewMat(rows, cols)

	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			r, g, b, alpha := img.At(c, r).RGBA()
			if alpha == 65536 {
				// only RGB
				value := NewValue(3, []int{
					int(r >> 8),
					int(g >> 8),
					int(b >> 8),
				})
				mat.Set(r, c, value)
			} else {
				value := NewValue(4, []int{
					int(r >> 8),
					int(g >> 8),
					int(b >> 8),
					int(alpha >> 8),
				})
				mat.Set(r, c, value)
			}
		}
	}

	return mat, nil
}

func NewFromFile(file string) (*Mat, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	im, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return NewFromImage(im)
}

func NewMat(rows, cols int) *Mat {
	if rows <= 0 || cols <= 0 {
		panic("New Mat fail: invalid rows or columns")
	}

	return &Mat{
		data: make([]Value, rows*cols),
		rows: rows,
		cols: cols,
	}
}

func (m *Mat) Rows() int {
	return m.rows
}

func (m *Mat) Cols() int {
	return m.cols
}

func (m *Mat) At(row, col int) Value {
	if row >= m.rows || row < 0 || col >= m.cols || col < 0 {
		panic("Mat: at position out of range")
	}

	return m.data[m.cols*row+col]
}

func (m *Mat) Set(row, col int, v Value) {
	if row >= m.rows || row < 0 || col >= m.cols || col < 0 {
		panic("Mat: set position out of range")
	}

	m.data[m.cols*row+col] = v
}

func (m *Mat) Copy() *Mat {
	matrix := NewMat(m.rows, m.cols)
	matrix.data = m.data
	return matrix
}

func (m *Mat) EqualShape(matrix *Mat) bool {
	return m.rows == matrix.rows && m.cols == matrix.cols
}

func (m *Mat) Shape() (int, int) {
	return m.rows, m.cols
}

func (m *Mat) Row(i int) []Value {
	if i < 0 {
		i += m.rows
	}

	row := make([]Value, m.cols)
	for col := 0; col < m.cols; col++ {
		row[col] = m.At(i, col)
	}

	return row
}

func (m *Mat) SetRow(i int, v []Value) {
	if i < 0 {
		i += m.rows
	}

	if len(v) < m.cols {
		panic("Mat SetRow fail: mismatch length")
	}

	for col := 0; col < m.cols; col++ {
		m.Set(i, col, v[col])
	}
}

func (m *Mat) Col(i int) []Value {
	if i < 0 {
		i += m.cols
	}

	col := make([]Value, m.rows)
	for row := 0; row < m.rows; row++ {
		col[row] = m.At(row, i)
	}

	return col
}

func (m *Mat) SetCol(i int, v []Value) {
	if i < 0 {
		i += m.cols
	}

	if len(v) < m.rows {
		panic("Mat SetCol fail: mismatch length")
	}

	for row := 0; row < m.rows; row++ {
		m.Set(row, i, v[row])
	}
}

func (m *Mat) String() string {
	result := make([]string, m.rows)
	for row := 0; row < m.rows; row++ {
		rowResult := make([]string, m.cols)
		for col := 0; col < m.cols; col++ {
			rowResult[col] = m.At(row, col).String()
		}
		result[row] = "[" + strings.Join(rowResult, ",") + "]"
	}

	return "[" + strings.Join(result, "\n") + "]\n"
}
