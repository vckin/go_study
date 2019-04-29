package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "how do you do"
	s1 := strings.Split(s," ");
	s2 := make(map[string]int,len(s1))
	for _,v := range  s1{
		s2[v] += 1
	}
	for k,v := range  s2{
		fmt.Printf("%s = %d \n",k,v)
	}
}
