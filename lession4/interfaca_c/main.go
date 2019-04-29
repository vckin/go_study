package main

import "fmt"

type animol interface {
	move()
	speak()
}

type cat struct {
	name string
}

func (c *cat)move()  {
	fmt.Println(&c)
	fmt.Println("会动")
}

func (c cat)speak()  {
	fmt.Println(&c)
	fmt.Println("会说")
}

func main() {
	var c animol
	c = &cat{"sadas"}
	//fmt.Printf("%v\n",&cc)
	//c = cc
	fmt.Println(&c)
	c.speak()
	c.move()
}
