package core

import (
	_ "log"
)

type Mat struct {
	data     []Value
	rows     int
	cols     int
	channels int
}

func New(rows, cols, channels int) *Mat {
	if row <= 0 || width <= 0 || channels <= 0 {
		panic("New Mat fail: invalid rows, columns or channels")
	}

	return &Mat{
		data:     make([]Value, rows*cols*channels),
		rows:     rows,
		cols:     cols,
		channels: channels,
	}
}

func (m *Mat) Rows() int {
	return m.rows
}

func (m *Mat) Cols() int {
	return m.cols
}

func (m *Mat) Channels() int {
	return m.channels
}

func (m *Mat) At(row, col, channel int) Value {
	if row >= m.rows || row < 0 || col >= m.cols ||
		col < 0 || channel >= m.channels || channel < 0 {
		panic("Mat: at position out of range")
	}

	index := m.rows*m.cols*channel + m.cols*row + col
	return m.data[index]
}

func (m *Mat) Set(row, col, channel int, v Value) {
	if row >= m.rows || row < 0 || col >= m.cols ||
		col < 0 || channel >= m.channels || channel < 0 {
		panic("Mat: set position out of range")
	}

	index := m.rows*m.cols*channel + m.cols*row + col
	m.data[index] = v
}
