package main

import (
	"errors"
	"fmt"
)

type errorString string

func (e errorString) Error() string {
	return string(e)
}

func New(text string) error {
	b := errorString(text)
	return &b
}

var ErrNameType = New("EOF")
var ErrStructType = errors.New("EOF")

func TestGoError() {
	if ErrNameType == New("EOF") {
		fmt.Println("Name Type Error")
	}
	if ErrStructType == errors.New("EOF") {
		fmt.Println("struct Eype Error")
	}

}
