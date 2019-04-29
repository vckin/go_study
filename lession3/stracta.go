package main

import "fmt"

type persion struct {
	name string
	age int
	sex string
	hopyy []string
}

func main() {
	var doudou = persion{
		"doudou",
		12,
		"女",
		[]string{"抽烟","喝酒","烫头",},
	}
	fmt.Println(doudou)

	var bb = persion{
		name:"baobao",
		sex:"男",
		age:56,
		hopyy:[]string{"dd","sasda"},

	}
	fmt.Println(bb)
}