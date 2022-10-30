package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func RenameFile() {
	folder := "F:\\MySQL\\视频"
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		//oldPath := path.Join(folder, file.Name())
		newName := strings.Split(file.Name(), "-")[1]
		newPath := path.Join(folder, newName)
		//os.Rename(oldPath, newPath)
		fmt.Println(newPath)
	}
}
