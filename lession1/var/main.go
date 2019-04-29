package main

import "fmt"

var a string
func main() {
	fmt.Println(a)

	var (
		b string
		c int
		d bool
	)
	fmt.Println(b,c,d)

	e := 100
	fmt.Printf("%d 你好！",e)

	f := "dsssa"
	fmt.Printf("%s 666",f)

}
