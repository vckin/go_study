package main

import "fmt"

func main() {
	a := [3][2]int{
		{2,5},
		{6,8},
	}
	b := a
	b[2][1] = 100
	fmt.Println(a)
	fmt.Println(b)

}
