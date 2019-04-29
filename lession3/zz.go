package main

import "fmt"

func main() {
	a := [3]int{
		1,2,3,
	}
	modifyAarray(&a)
	fmt.Println(a)

}
func modifyAarray(a1 *[3]int){
	a1[1] = 100
}
