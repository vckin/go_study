package main

import (
	"fmt"
	"os"
)

type Book struct {
	title   string
	author  string
	price   float32
	publish bool
}
//打印菜单
func pritcd(){
	fmt.Println("1. 添加书籍")
	fmt.Println("2. 修改书籍信息")
	fmt.Println("3. 展示所有书籍")
	fmt.Println("4. 退出")
	fmt.Println()
}
func addbook(a *Book){
	fmt.Println("标题")
	fmt.Scan(&a.title)
	fmt.Println("作者")
	fmt.Scan(&a.author)
	fmt.Println("价格")
	fmt.Scan(&a.price)
	fmt.Println("是否公开")
	fmt.Scan(&a.publish)

}

func updatebook(a *Book){
	fmt.Println("标题")
	fmt.Scan(&a.title)
	fmt.Println("作者")
	fmt.Scan(&a.author)
	fmt.Println("价格")
	fmt.Scan(&a.price)
	fmt.Println("是否公开")
	fmt.Scan(&a.publish)

}
func seeBook(a *Book){
	fmt.Printf("标题:%v 作者:%v 价格%v 是否公开:%v\n",a.title,a.author,a.price, a.publish)

}
func main() {
	var bookArray = Book{}

	var chose int
	//fmt.Println(bookArray)
	for {
		if(chose == 0){
			pritcd()
		}
		fmt.Scan(&chose)
		if(chose == 1){
			addbook(&bookArray)
			chose = 0
		}
		if(chose == 2){
			updatebook(&bookArray)
			chose = 0
		}
		if(chose == 3){
			seeBook(&bookArray)
			chose = 0
		}
		if(chose == 4){
			fmt.Println("退出")
			os.Exit(0)
		}
	}

	//fmt.Println(bookArray)


}
