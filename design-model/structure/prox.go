package main

import "fmt"

type Patent interface {
	Innovate()
}

type UtilityModelPatent struct{}

func (p *UtilityModelPatent) Innovate() {
	fmt.Println("实用新型专利的创新")
}

type InventionPatent struct{}

func (p *InventionPatent) Innovate() {
	fmt.Println("发明专利的创新")
}

type AppearancePatent struct{}

func (p *AppearancePatent) Innovate() {
	fmt.Println("外观专利的创新")
}

type patentAgent struct{}

func (p *patentAgent) Innovate() {

}
