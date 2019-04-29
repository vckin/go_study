package main

import "fmt"

//const (
//	a = 1
//)
//
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
const (
	n1 = iota
	_
	n3
	n4
)

const (
	a, b = iota + 1, iota +2
	c, d
	e, f
)

func main() {
	fmt.Println(KB,MB,GB,TB,PB)
	fmt.Println(a,b,c,d,e,f)
	fmt.Println(n1,n3,n4)
}
