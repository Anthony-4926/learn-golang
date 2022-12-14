package data_structure

import "unsafe"

type Bit struct {
	bitArr  []uint
	intSize int
	arrSize int
}

func New(bitNum int) *Bit {
	b := &Bit{}
	b.intSize = int(unsafe.Sizeof(uint(0)))
	b.arrSize = bitNum/b.intSize + 1
	b.bitArr = make([]uint, b.arrSize, b.arrSize)
	return b
}

func (b Bit) locate(index int) (arrIndex, intIndex int) {
	arrIndex = index / b.intSize
	intIndex = index % b.intSize
	return
}

func (b Bit) Set(index int) {
	arrIndex, intIndex := b.locate(index)
	b.bitArr[arrIndex] |= (1 << intIndex)
}

func (b Bit) Get(index int) int {
	arrIndex, intIndex := b.locate(index)
	return int((b.bitArr[arrIndex] & (1 << intIndex)) >> intIndex)
}
