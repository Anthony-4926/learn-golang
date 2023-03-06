package main

import "fmt"

type trieNode struct {
	Pass  int
	End   int
	Nexts [26]*trieNode
}

func main() {
	a := &trieNode{
		Pass:  0,
		End:   0,
		Nexts: [26]*trieNode{},
	}
	//a.Nexts[0] = &trieNode{
	//	Pass:  0,
	//	End:   0,
	//	Nexts: [26]*trieNode{},
	//}
	fmt.Println(a.Nexts)
}
