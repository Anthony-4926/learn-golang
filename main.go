package main

import (
	"log"
)

func main() {

	log.Println(`BindWith(\"interface{}, binding.Binding\") error is going to
	be deprecated, please check issue #662 and either use MustBindWith() if you
	want HTTP 400 to be automatically returned if any error occur, or use
	ShouldBindWith() if you need to manage the error.`)
}
