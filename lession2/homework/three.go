package main

import (
	"fmt"
	"strings"
)
var(
	coins = 50
	users = []string{
		"Mattew","Sarah","Augustus","Heidi","Emilie","Peter","Giana","Adriano","Aaron","Elizabeth",
	}
	distribution = make(map[string]int,len(users))
)

func dispatchCoin()int {
	cosNumber := map[string]int{
		"e":1,
		"i":2,
		"o":3,
		"u":4,
	}
	for _,v := range users{
		s := strings.ToLower(v)
		for _,i := range s{
			if(coins > 0){
				distribution[v] += cosNumber[string(i)]
				coins -= cosNumber[string(i)]
			}
		}
	}
	fmt.Println(distribution)
	return coins
}
func main() {
	left := dispatchCoin()
	fmt.Println("剩下",left)
}