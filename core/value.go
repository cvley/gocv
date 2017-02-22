package core

import (
	"fmt"
	"strings"
)

// Value interface is used to convert value to desired type
type Value interface {
	Char() Char
	Int() Int
	Float() Float
	String() string
	Channels() int
}

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

func (i Char) Char() Char {
	return i
}

func (i Char) Int() Int {
	intv := Char{
		data:     make([]int, len(i.data)),
		channels: i.channels,
	}
	for i, v := range i.data {
		intv.data[i] = int(v)
	}
	return intv

}
func (i Char) Float() Float {
	float := Float{
		data:     make([]float64, len(i.data)),
		channels: i.channels,
	}

	for i, v := range i.data {
		float.data[i] = float64(v)
	}
	return float
}

func (i Char) String() string {
	result := make([]string, i.channels)
	for i, v := range i.data {
		result[i] = fmt.Sprintf("%d", v)
	}

	return "[" + strings.Join(result, ",") + "]"
}

func (i Char) Channels() int {
	return i.channels
}

func (i Int) Char() Char {
	char := Char{
		data:     make([]int, len(i.data)),
		channels: i.channels,
	}
	for i, v := range i.data {
		char.data[i] = char(v)
	}
	return char
}

func (i Int) Int() Int {
	return i
}

func (i Int) Float() Float {
	float := Float{
		data:     make([]float64, len(i.data)),
		channels: i.channels,
	}

	for i, v := range i.data {
		float.data[i] = float64(v)
	}
	return float
}

func (i Int) String() string {
	result := make([]string, i.channels)
	for i, v := range i.data {
		result[i] = fmt.Sprintf("%d", v)
	}

	return "[" + strings.Join(result, ",") + "]"
}

func (i Int) Channels() int {
	return i.channels
}

func (i Float) Char() Char {
	char := Char{
		data:     make([]int, len(i.data)),
		channels: i.channels,
	}
	for i, v := range i.data {
		char.data[i] = char(v)
	}
	return char
}

func (i Float) Int() Int {
	return int(i)
}

func (i Float) Float() Float {
	float := Char{
		data:     make([]float64, len(i.data)),
		channels: i.channels,
	}
	for i, v := range i.data {
		float.data[i] = float64(v)
	}
	return float
}

func (i Float) String() string {
	result := make([]string, i.channels)
	for i, v := range i.data {
		result[i] = fmt.Sprintf("%.4f", v)
	}

	return "[" + strings.Join(result, ",") + "]"
}

func (i Float) Channels() int {
	return i.channels
}

func char(v interface{}) int {
	switch v.(type) {
	case int:
		if v < 0 {
			return 0
		}
		if v > 255 {
			return 255
		}
		return v

	case float64:
		if v < 0 {
			return 0
		}
		if v > 255 {
			return 255
		}
		return int(v)

	default:
		panic("char func not support value type")
	}
}
