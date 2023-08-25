package main

import "testing"

func TestFacade(t *testing.T) {
	face := NewFacade()
	face.open()
}
