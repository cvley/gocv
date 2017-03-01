package core

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// Format represents image type
type Format int

var (
	// FormatJPEG represents JPEG format
	FormatJPEG Format = iota
	// FormatPNG represents JPEG format
	FormatPNG Format
)

var (
	// ErrShape represents mismatch shape error
	ErrShape = errors.New("mismatch shape")
)

var (
	defaultJPEGQuality = 95
)

// Mat represents a matrix of Values
type Mat struct {
	data []Value
	rows int
	cols int
}

// NewFromImage returns a Mat from image
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

// NewFromFile returns a Mat from a image file
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

// NewMat returns an all zeros Mat of input rows and columns
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

// Init returns a Mat with input values
func Init(rows, cols int, value Value) *Mat {
	mat := NewMat(rows, cols)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			mat.Set(r, c, value)
		}
	}

	return mat
}

// InitEye returns a Mat with input values in the eye position
func InitEye(rows, cols int, value Value) *Mat {
	mat := NewMat(rows, cols)
	min := rows
	if min > cols {
		min = cols
	}

	for i := 0; i < min; i++ {
		mat.Set(i, i, value)
	}

	return mat
}

// Rows returns the rows number of Mat
func (m *Mat) Rows() int {
	return m.rows
}

// Cols returns the columns number of Mat
func (m *Mat) Cols() int {
	return m.cols
}

// At returns Values in the position of input row and col
func (m *Mat) At(row, col int) Value {
	if row >= m.rows || row < 0 || col >= m.cols || col < 0 {
		panic("Mat: at position out of range")
	}

	return m.data[m.cols*row+col]
}

// Set will reset the value in the position of Mat
func (m *Mat) Set(row, col int, v Value) {
	if row >= m.rows || row < 0 || col >= m.cols || col < 0 {
		panic("Mat: set position out of range")
	}

	m.data[m.cols*row+col] = v
}

// Copy returns a Mat with the same shape and data
func (m *Mat) Copy() *Mat {
	matrix := NewMat(m.rows, m.cols)
	copy(matrix.data, m.data)
	return matrix
}

// Sub returns a sub mat with the input shape
func (m *Mat) Sub(rect Rect) (*Mat, error) {
	rows := rect.Dx()
	cols := rect.Dy()
	
	mat := NewMat(rows, cols)
	// TODO Rect
}

// Reshape returns a new shape Mat
func (m *Mat) Reshape(rows, cols int) (*Mat, error) {
	if m.rows == rows && m.cols == cols {
		return m, nil
	}

	if m.rows * m.cols != rows * cols {
		return nil, ErrShape
	}

	mat := NewMat(rows, cols)
	copy(mat.data, m.data)
	return mat
}

// Scale performs bitwise scale of the Mat
func (m *Mat) Scale(scale float64) {
	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			v := m.At(r, c)
			m.Set(r, c, v.Scale(scale))
		}
	}
}

// Mul performs two matrix multiplication with the same shape
func (m *Mat) Mul(matrix *Mat) error {
	if !m.EqualShape(matrix) {
		return ErrShape
	}

	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			v := m.At(r, c)
			if err := m.At(r, c).Mul(matrix.At(r, c)); err != nil {
				return err
			}
			m.Set(r, c, v)
		}
	}

	return nil
}

// Add performs two matrix addition with the same shape
func (m *Mat) Add(matrix *Mat) {
	if !m.EqualShape(matrix) {
		return ErrShape
	}

	for r := 0; r < m.rows; r++ {
		for c := 0; c < m.cols; c++ {
			v := m.At(r, c)
			if err := m.At(r, c).Add(matrix.At(r, c)); err != nil {
				return err
			}
			m.Set(r, c, v)
		}
	}

	return nil
}

// EqualShape returns whether the input Mat has the same shape
func (m *Mat) EqualShape(matrix *Mat) bool {
	return m.rows == matrix.rows && m.cols == matrix.cols
}

// Equal returns whether the input Mat is the same
func (m *Mat) Equal(matrix *Mat) bool {
	if !m.EqualShape(matrix) {
		return false
	}

	for c := 0; c < m.cols; c++ {
		for r := 0; r < m.rows; r++ {
			if !m.At(r, c).Equal(matrix.At(r, c)) {
				return false
			}
		}
	}

	return true
}

// Shape returns the shape of Mat, i.e. rows and columns
func (m *Mat) Shape() (int, int) {
	return m.rows, m.cols
}

// Row returns values of the input index row of Mat
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

// SetRow will reset the input index row of Mat
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

// Col returns values of the input index column of Mat
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

// SetCol will reset the input index column of Mat
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

// ToImage returns a image container of Mat
func (m *Mat) ToImage() image.Image {
	im := image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: m.cols,
			Y: m.rows,
		},
	})

	switch m.data[0].Channels() {
	case 1:
		for c := 0; c < m.cols; c++ {
			for r := 0; r < m.rows; r++ {
				value := m.At(r, c).Char()
				im.Set(c, r, color.Gray{
					Y: uint8(value.At(0)),
				})
			}
		}

	case 3:
		for c := 0; c < m.cols; c++ {
			for r := 0; r < m.rows; r++ {
				value := m.At(r, c).Char()
				im.Set(c, r, color.RGBA{
					R: uint8(value.At(0)),
					G: uint8(value.At(1)),
					B: uint8(value.At(2)),
				})
			}
		}

	case 4:
		for c := 0; c < m.cols; c++ {
			for r := 0; r < m.rows; r++ {
				value := m.At(r, c).Char()
				im.Set(c, r, color.RGBA{
					R: uint8(value.At(0)),
					G: uint8(value.At(1)),
					B: uint8(value.At(2)),
					A: uint8(value.At(3)),
				})
			}
		}

	default:
		panic("Mat to Image fail: only 1, 3 or 4 channels are supported.")
	}
}

// Imwrite will save the Mat to the disk with specified format and name
func (m *Mat) Imwrite(name string, format Format) error {
	im := m.ToImage()

	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	switch format {
	case FormatJPEG:
		if err := jpeg.Encode(f, im, &jpeg.Options{Quality: defaultJPEGQuality}); err != nil {
			return err
		}

	case FormatPNG:
		if err := png.Encode(f, im); err != nil {
			return err
		}

	default:
		panic("Mat Imwrite only support jpeg and png")
	}
}

// String returns the string format of Mat data
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
