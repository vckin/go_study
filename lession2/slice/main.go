package main

import "fmt"

func main() {
	a := [...]int64{1,2000,3,4,5,6,7}
	b := a[:]
	fmt.Printf("b:%v %d %d %p \n",b,len(b),cap(b),b)

	c := a[3:5]
	fmt.Printf("c:%v %d %d %p \n",c,len(c),cap(c),c)

	c = append(c,100,200,300)
	fmt.Printf("c:%v %d %d %p \n",c,len(c),cap(c),c)

	fmt.Printf("a:%v %d %d %p %p %p %p %p %p \n",a,len(a),cap(a),&a[0],&a[1],&a[2],&a[3],&a[4],&a[5])

	d := [...]string{"a","fssdafsfdsfdsfssfdsb","c","d","e","f"}
	fmt.Printf("d:%v %d %d %p %p %p %p %p %p \n",d,len(d),cap(d),&d[0],&d[1],&d[2],&d[3],&d[4],&d[5])


}
