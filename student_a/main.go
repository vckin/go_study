package main

import (
	"fmt"
	"github.com/vckin/go_study/student"
	_ "github.com/vckin/go_study/initc"
)

func main()  {
	a := student.Add(3,6)
	fmt.Println(a)
}
