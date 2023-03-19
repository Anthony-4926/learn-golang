package data_structure

import (
	"testing"
)

func TestHeapGreater_Add(t *testing.T) {
	heap := NewHeapGreater(func(a, b interface{}) bool {
		aa := a.(int)
		bb := b.(int)
		if aa < bb {
			return true
		}
		return false
	})

	heap.Add(1)
}

func TestHeapGreater_Contains(t *testing.T) {
	heap := NewHeapGreater(func(a, b interface{}) bool {
		aa := a.(int)
		bb := b.(int)
		if aa < bb {
			return true
		}
		return false
	})

	heap.Add(1)
	heap.Add(2)
	heap.Add(3)
	if heap.Contains(3) != true {
		t.Errorf("heap中本应该存在3")
	}
	if heap.Contains(4) == true {
		t.Errorf("heap中本不应该存在4")
	}
}

func TestHeapGreater_Remove(t *testing.T) {
	heap := NewHeapGreater(func(a, b interface{}) bool {
		aa := a.(int)
		bb := b.(int)
		if aa < bb {
			return true
		}
		return false
	})

	heap.Add(1)
	heap.Add(2)
	heap.Add(3)
	if heap.Contains(3) != true {
		t.Errorf("heap中本应该存在3")
	}
	heap.Remove(3)
	if heap.Contains(3) == true {
		t.Errorf("heap中本不应该存在3")
	}
}

func TestHeapGreater_Remove2(t *testing.T) {
	heap := NewHeapGreater(func(a, b interface{}) bool {
		aa := a.(int)
		bb := b.(int)
		if aa < bb {
			return true
		}
		return false
	})

	heap.Add(1)
	heap.Add(2)
	heap.Add(3)
	if heap.Contains(2) != true {
		t.Errorf("heap中本应该存在2")
	}
	heap.Remove(2)
	if heap.Contains(2) == true {
		t.Errorf("heap中本不应该存在2")
	}
}
