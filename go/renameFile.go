package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func RenameFile() {
	folder := "F:\\kafka"
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		oldPath := path.Join(folder, file.Name())
		newName := strings.Split(file.Name(), "Java必备 - ")[1]
		newPath := path.Join(folder, newName)
		os.Rename(oldPath, newPath)
		//fmt.Println(newPath)
	}
}
