package main

import (
	"testing"
)

func createContext() component {
	file5 := file{name: "file5"}

	folder3 := folder{
		name:       "folder3",
		components: []component{file5},
	}

	file4 := file{name: "file4"}
	folder2 := folder{
		name:       "folder2",
		components: []component{file4, folder3},
	}

	file1 := file{name: "file1"}
	file2 := file{name: "file2"}
	file3 := file{name: "file3"}

	root := folder{
		name:       "root",
		components: []component{file1, file2, file3, folder2},
	}
	return root
}

func f(c component) {
	c.open()
	if folder, ok := c.(folder); ok {
		for _, next := range folder.components {
			f(next)
		}
	}
	c.close()
}

func TestBuilder(t *testing.T) {
	root := createContext()
	f(root)
}
