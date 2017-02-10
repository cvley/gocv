package transform

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
	"strings"

	"log"
)

type Matrix struct {
	data   []float64
	height int
	width  int
}

func New(height, width int) (*Matrix, error) {
	if height <= 0 || width <= 0 {
		return nil, errors.New("invalid height or width")
	}

	return &Matrix{
		data:   make([]float64, height*width),
		height: height,
		width:  width,
	}, nil
}

func NewFromFile(name string) ([]*Matrix, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	im, format, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	log.Printf("open image %s, format %s", name, format)

	return NewFromImage(im)
}

func NewFromImage(img image.Image) ([]*Matrix, error) {
	if img == nil {
		return nil, errors.New("invalid input image")
	}

	rect := img.Bounds()
	matR, _ := New(rect.Max.Y, rect.Max.X)
	matG, _ := New(rect.Max.Y, rect.Max.X)
	matB, _ := New(rect.Max.Y, rect.Max.X)

	for h := 0; h < rect.Max.Y; h++ {
		for w := 0; w < rect.Max.X; w++ {
			r, g, b, _ := img.At(w, h).RGBA()
			matR.Set(h, w, float64(r>>8))
			matG.Set(h, w, float64(g>>8))
			matB.Set(h, w, float64(b>>8))
		}
	}

	return []*Matrix{matR, matG, matB}, nil
}

func clip(v float64) float64 {
	if v < 0 {
		log.Println(v)
		return 0
	}
	if v > 255.0 {
		log.Println(v)
		return 255.0
	}
	return v
}

func ToImage(mats []*Matrix, quality int, name string) error {
	im := image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: mats[0].Width(),
			Y: mats[0].Height(),
		},
	})

	for h := 0; h < mats[0].Height(); h++ {
		for w := 0; w < mats[0].Width(); w++ {
			im.Set(w, h, color.RGBA{
				R: uint8(clip(mats[0].At(h, w))),
				G: uint8(clip(mats[1].At(h, w))),
				B: uint8(clip(mats[2].At(h, w))),
			})
		}
	}

	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := jpeg.Encode(f, im, &jpeg.Options{Quality: quality}); err != nil {
		return err
	}

	return nil
}

func (m *Matrix) Copy() *Matrix {
	matrix, _ := New(m.height, m.width)
	// TODO check this data copy
	matrix.data = m.data
	return matrix
}

func (m *Matrix) EqualShape(matrix *Matrix) bool {
	return m.width == matrix.width && m.height == matrix.height
}

func (m *Matrix) Equal(matrix *Matrix) bool {
	if !m.EqualShape(matrix) {
		return false
	}

	for h := 0; h < m.height; h++ {
		for w := 0; w < m.width; w++ {
			if math.Abs(m.At(h, w)-matrix.At(h, w)) > 1e-9 {
				return false
			}
		}
	}

	return true
}

func (m *Matrix) Rows() int {
	return m.height
}

func (m *Matrix) Height() int {
	return m.height
}

func (m *Matrix) Cols() int {
	return m.width
}

func (m *Matrix) Width() int {
	return m.width
}

func (m *Matrix) At(h, w int) float64 {
	if h >= m.height || w >= m.width || h < 0 || w < 0 {
		panic("get position out of range")
	}

	return m.data[h*m.width+w]
}

func (m *Matrix) Set(h, w int, v float64) {
	if h >= m.height || w >= m.width || h < 0 || w < 0 {
		panic("set postion out of range")
	}

	m.data[h*m.width+w] = v
}

func (m *Matrix) Sub(startH, endH, startW, endW int) (*Matrix, error) {
	if startH > endH || startW > endW || startH < 0 || startW < 0 ||
		endH > m.height || endW > m.width || startH >= m.height || startW >= m.width {
		return nil, errors.New("sub matrix parameters out of range")
	}

	sub, err := New(endH-startH, endW-startW)
	if err != nil {
		return nil, err
	}
	for h := startH; h < endH; h++ {
		for w := startW; w < endW; w++ {
			sub.Set(h-startH, w-startW, m.At(h, w))
		}
	}

	return sub, nil
}

func (m *Matrix) GetRow(index int) []float64 {
	row := make([]float64, m.width)
	for i := 0; i < m.width; i++ {
		row[i] = m.At(index, i)
	}

	return row
}

func (m *Matrix) SetRow(index int, values []float64) {
	if len(values) != m.width {
		panic("set row mismatch length")
	}
	for i := 0; i < m.width; i++ {
		m.Set(index, i, values[i])
	}
}

func (m *Matrix) GetCol(index int) []float64 {
	col := make([]float64, m.height)
	for i := 0; i < m.height; i++ {
		col[i] = m.At(i, index)
	}

	return col
}

func (m *Matrix) SetCol(index int, values []float64) {
	if len(values) != m.height {
		panic("set col mismatch length")
	}
	for i := 0; i < m.height; i++ {
		m.Set(i, index, values[i])
	}
}

func (m *Matrix) Similarity(mat *Matrix) float64 {
	if !m.EqualShape(mat) {
		panic("Similarity fail: mismatch shape")
	}

	var Var, Var1, Var2 float64
	for h := 0; h < m.height; h++ {
		for w := 0; w < m.width; w++ {
			var v float64
			if mat.At(h, w) > 127 {
				v = 255.0
			} else {
				v = 0.0
			}
			Var += m.At(h, w) * v
			Var1 += math.Pow(m.At(h, w), 2)
			Var2 += math.Pow(v, 2)
		}
	}

	sim := Var / (math.Sqrt(Var1) * math.Sqrt(Var2))
	return sim
}

func (m *Matrix) String() string {
	cols := make([]string, m.height)
	for h := 0; h < m.height; h++ {
		row := m.GetRow(h)
		strRow := toString(row)
		cols[h] = "[" + strings.Join(strRow, ", ") + "]"
	}
	return "[" + strings.Join(cols, "\n") + "]"
}

func toString(values []float64) []string {
	str := make([]string, len(values))
	for i, v := range values {
		str[i] = fmt.Sprintf("%.4f", v)
	}
	return str
}
