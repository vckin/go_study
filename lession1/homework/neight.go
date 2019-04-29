package main

import "fmt"

func neightNum(a int,b int){
	if a == 1 && b==1{
		fmt.Printf("%d*%d=%d ",a,b,a*b)
		return
	}
	if a > 0 {
		neightNum(a-1,b)
		fmt.Printf("%d*%d=%d ",a,b,a*b)
	}else{
		b--
		a = b
		neightNum(a,b)
		fmt.Printf("\n")
	}
}

func main() {
	neightNum(9,9)
}
