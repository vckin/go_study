package main

import (
	"fmt"
)
//执行判断范围
func rangPrime(a int,b int){
	if b % 2 == 0{
		b--
	}
	for ;b > a; b = b-2 {
		if isPrime(b){
			fmt.Printf("%d ",b)
		}
	}
}
//判断是否为素数
func isPrime(a int)(bool){
	if(a == 2){
		return true
	}
	//再次对2取余是为了函数的准确性
	if a < 2 || a % 2 == 0{
		return false
	}
	//跳度为2 是为了性能
	b := (a+1)/2;
	for i := 3;i < b;i = i + 2{
		if a % i == 0 {
			//fmt.Println(i,a)
			return false
		}
	}
	return true
}

func main() {
	rangPrime(200,1000)
}
