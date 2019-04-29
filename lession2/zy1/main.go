package main

import "fmt"

func main() {
	a := [5]int{1,3,5,7,8}
	var a2 int
	for _,a1 := range a{
		a2 += a1
	}
	fmt.Println(a2)

	for i := 0;i < len(a);i++{
		for j := i+1; j < len(a); j++ {
			if(a[i] + a[j] == 8){
				fmt.Printf("{%d,%d} \n",i,j)
			}
		}
	}

}
