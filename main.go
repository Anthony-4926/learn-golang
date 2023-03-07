package main

//type People interface {
//	speek()
//}
//
//type Student struct {
//	name string
//}
//
//type Worker struct {
//	name string
//}
//
//func (s *Student) speek() {
//	fmt.Printf("我的名字叫%s", s.name)
//}

func main() {
	var ch chan int = make(chan int, 2)
	a := 1
	ch <- a

}
