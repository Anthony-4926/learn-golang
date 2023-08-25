package main

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	phone := Phone{}
	target := Adapter{}
	phone.charge(target)
}
