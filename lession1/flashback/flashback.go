package main

import (
	"fmt"
)

func mbBack(s string)(string){
	s2 := ""
	for _,s1 := range s{
		s2 = string(s1) + s2
	}
	return s2
}
func main() {
	fmt.Println(mbBack("hello你好666小姐姐"))
}
