package core

import (
	"fmt"
	"strings"
	"errors"
)

var (
	ErrChannels = errors.New("mismatch channels")
)

// Value interface is used to convert value to desired type
type Value interface {
	Char() *Char
	Int() *Int
	Float() *Float
	Scale(float64)
	Mul(Value) error
	Add(Value) error
	String() string
	Channels() int
}

// Char represents values in range [0, 255]
type Char struct {
	data     []int
	channels int
}

type Int struct {
	data     []int
	channels int
}
type Float struct {
	data     []float64
	channels int
}

func NewValue(c int, v interface{}) Value {
	switch v.(type) {
	case int:
		char := &Char{
			channels: c,
			data:     make([]int, c),
		}
		char.data[0] = v.(int)
		return char

	case float64:
		float := &Float{
			channels: c,
			data:     make([]float64, c),
		}
		float.data[0] = v.(float64)
		return float

	case []int:
		char := &Char{
			channels: c,
			data:     make([]int, c),
		}
		for i := 0; i < c; i++ {
			char.data[i] = v.([]int)[i]
		}
		return char

	case []float64:
		float := &Float{
			channels: c,
			data:     make([]float64, c),
		}
		for i := 0; i < c; i++ {
			float.data[i] = v.([]float64)[i]
		}
		return float

	default:
		panic("Value type not supported")
	}
}

func (i *Char) Char() *Char {
	return i
}

func (i *Char) Int() *Int {
	intv := &Int{
		data:     make([]int, i.channels),
		channels: i.channels,
	}
	for i, v := range i.data {
		intv.data[i] = v
	}
	return intv
}

func (i *Char) Float() *Float {
	float := &Float{
		data:     make([]float64, i.channels),
		channels: i.channels,
	}

	for i, v := range i.data {
		float.data[i] = float64(v)
	}
	return float
}

func (i *Char) Scale(scale float64) {
}

func (i *Char) Mul(v Value) error {
	if i.channels != v.Channels() {
		return ErrChannels
	}
}

func (i *Char) At(idx int) int {
	if idx < 0 {
		idx += i.channels
	}

	return i.data[idx]
}

func (i *Char) String() string {
	result := make([]string, i.channels)
	for i, v := range i.data {
		result[i] = fmt.Sprintf("%d", v)
	}

	return "[" + strings.Join(result, ",") + "]"
}

func (i *Char) Channels() int {
	return i.channels
}

func (i *Int) Char() *Char {
	char := &Char{
		data:     make([]int, i.channels),
		channels: i.channels,
	}
	for i, v := range i.data {
		char.data[i] = toChar(v)
	}
	return char
}

func (i *Int) Int() *Int {
	return i
}

func (i *Int) Float() *Float {
	float := &Float{
		data:     make([]float64, i.channels),
		channels: i.channels,
	}

	for i, v := range i.data {
		float.data[i] = float64(v)
	}
	return float
}

func (i *Int) String() string {
	result := make([]string, i.channels)
	for i, v := range i.data {
		result[i] = fmt.Sprintf("%d", v)
	}

	return "[" + strings.Join(result, ",") + "]"
}

func (i *Int) Channels() int {
	return i.channels
}

func (i *Int) At(idx int) int {
	if idx < 0 {
		idx += i.channels
	}

	return i.data[idx]
}

func (i *Float) Char() *Char {
	char := &Char{
		data:     make([]int, i.channels),
		channels: i.channels,
	}
	for i, v := range i.data {
		char.data[i] = toChar(v)
	}
	return char
}

func (i *Float) Int() *Int {
	out := &Int{
		data:     make([]int, i.channels),
		channels: i.channels,
	}

	for i, v := range i.data {
		out.data[i] = int(v)
	}

	return out
}

func (i *Float) Float() *Float {
	return i
}

func (i *Float) String() string {
	result := make([]string, i.channels)
	for i, v := range i.data {
		result[i] = fmt.Sprintf("%.4f", v)
	}

	return "[" + strings.Join(result, ",") + "]"
}

func (i *Float) Channels() int {
	return i.channels
}

func (i *Float) At(idx int) int {
	if idx < 0 {
		idx += i.channels
	}

	return i.data[idx]
}

func toChar(v interface{}) int {
	switch v.(type) {
	case int:
		if v.(int) < 0 {
			return 0
		}
		if v.(int) > 255 {
			return 255
		}
		return v.(int)

	case float64:
		if v.(float64) < 0 {
			return 0
		}
		if v.(float64) > 255 {
			return 255
		}
		return int(v.(float64))

	default:
		panic("char func not support value type")
	}
}
