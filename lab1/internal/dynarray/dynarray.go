package dynarray

import (
	"errors"

	. "github.com/stewkk/iu9-networks/lab1/internal/vertex"
)

type Array []Vertex

type Iterator struct {
	idx int
	arr Array
}

func (it *Iterator) Next() {
	it.idx++
}

func (it Iterator) HasNext() bool {
	return it.idx+1 < len(it.arr)
}

func (it Iterator) Vertex() *Vertex {
	return &it.arr[it.idx]
}

func (arr *Array) Insert(idx int, v Vertex) error {
	if idx < 0 || idx >= len(*arr) {
		return ErrOutOfBounds
	}
	*arr = append(*arr, Vertex{})
	copy((*arr)[idx+1:], (*arr)[idx:len(*arr)-1])
	(*arr)[idx] = v
	return nil
}

func (arr Array) Iterator(idx int) (Iterator, error) {
	if idx < 0 || idx >= len(arr) {
		return Iterator{}, ErrOutOfBounds
	}
	return Iterator{idx, arr}, nil
}

func (arr *Array) Remove(idx int) error {
	if idx < 0 || idx >= len(*arr) {
		return ErrOutOfBounds
	}
	*arr = append((*arr)[:idx], (*arr)[idx+1:]...)
	return nil
}

func (arr Array) First() Iterator {
	return Iterator{
		idx: 0,
		arr: arr,
	}
}

func (arr Array) Get(idx int) Iterator {
	return Iterator{
		idx: idx,
		arr: arr,
	}
}

var ErrOutOfBounds = errors.New("index is out of bounds")
