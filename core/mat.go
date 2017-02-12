package core

import (
	"log"
)

// Value interface is used to convert value to desired type
type Value interface {
	Int() int
	Int32() int32
	Int64() int64
	Uint() uint
	Uint32() uint32
	Uint64() uint64
	Float32() float32
	Float64() float64
}

type Int int

func (i Int) Int() int {
	return int(i)
}

func (i Int) Int32() int32 {
	return int32(i)
}

func (i Int) Int64() int64 {
	return int64(i)
}

func (i Int) Uint() uint {
	return uint(i)
}

func (i Int) Uint32() uint32 {
	return uint32(i)
}

func (i Int) Uint64() uint64 {
	return uint64(i)
}

func (i Int) Float32() float32 {
	return float32(i)
}

func (i Int) Float64() float64 {
	return float64(i)
}

type Int32 int32

func (i Int32) Int() int {
	return int(i)
}

func (i Int32) Int32() int32 {
	return int32(i)
}

func (i Int32) Int64() int64 {
	return int64(i)
}

func (i Int32) Uint() uint {
	return uint(i)
}

func (i Int32) Uint32() uint32 {
	return uint32(i)
}

func (i Int32) Uint64() uint64 {
	return uint64(i)
}

func (i Int32) Float32() float32 {
	return float32(i)
}

func (i Int32) Float64() float64 {
	return float64(i)
}

type Int64 int64

func (i Int64) Int() int {
	return int64(i)
}

func (i Int64) Int32() int32 {
	return int32(i)
}

func (i Int64) Int64() int64 {
	return int64(i)
}

func (i Int64) Uint() uint {
	return uint(i)
}

func (i Int64) Uint32() uint32 {
	return uint32(i)
}

func (i Int64) Uint64() uint64 {
	return uint64(i)
}

func (i Int64) Float32() float32 {
	return float32(i)
}

func (i Int64) Float64() float64 {
	return float64(i)
}

type Uint uint

func (i Uint) Int() int {
	return uint(i)
}

func (i Uint) Int32() int32 {
	return int32(i)
}

func (i Uint) Int64() int64 {
	return int64(i)
}

func (i Uint) Uint() uint {
	return uint(i)
}

func (i Uint) Uint32() uint32 {
	return uint32(i)
}

func (i Uint) Uint64() uint64 {
	return uint64(i)
}

func (i Uint) Float32() float32 {
	return float32(i)
}

func (i Uint) Float64() float64 {
	return float64(i)
}

type Uint32 uint32

func (i Uint32) Int() int {
	return int(i)
}

func (i Uint32) Int32() int32 {
	return int32(i)
}

func (i Uint32) Int64() int64 {
	return int64(i)
}

func (i Uint32) Uint() uint {
	return uint(i)
}

func (i Uint32) Uint32() uint32 {
	return uint32(i)
}

func (i Uint32) Uint64() uint64 {
	return uint64(i)
}

func (i Uint32) Float32() float32 {
	return float32(i)
}

func (i Uint32) Float64() float64 {
	return float64(i)
}

type Uint64 uint64

func (i Uint64) Int() int {
	return int(i)
}

func (i Uint64) Int32() int32 {
	return int32(i)
}

func (i Uint64) Int64() int64 {
	return int64(i)
}

func (i Uint64) Uint() uint {
	return uint(i)
}

func (i Uint64) Uint32() uint32 {
	return uint32(i)
}

func (i Uint64) Uint64() uint64 {
	return uint64(i)
}

func (i Uint64) Float32() float32 {
	return float32(i)
}

func (i Uint64) Float64() float64 {
	return float64(i)
}

type Float32 float32

func (i Float32) Int() int {
	return int(i)
}

func (i Float32) Int32() int32 {
	return int32(i)
}

func (i Float32) Int64() int64 {
	return int64(i)
}

func (i Float32) Uint() uint {
	return uint(i)
}

func (i Float32) Uint32() uint32 {
	return uint32(i)
}

func (i Float32) Uint64() uint64 {
	return uint64(i)
}

func (i Float32) Float32() float32 {
	return float32(i)
}

func (i Float32) Float64() float64 {
	return float64(i)
}

type Float64 float64

func (i Float64) Int() int {
	return int(i)
}

func (i Float64) Int32() int32 {
	return int32(i)
}

func (i Float64) Int64() int64 {
	return int64(i)
}

func (i Float64) Uint() uint {
	return uint(i)
}

func (i Float64) Uint32() uint32 {
	return uint32(i)
}

func (i Float64) Uint64() uint64 {
	return uint64(i)
}

func (i Float64) Float32() float32 {
	return float32(i)
}

func (i Float64) Float64() float64 {
	return float64(i)
}

type Point struct {
	X int
	Y int
}

type Mat struct {
	data     []Value
	rows     int
	cols     int
	channels int
}
