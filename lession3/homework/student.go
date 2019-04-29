package main

import (
	"fmt"
	"os"
)

type StudentList map[int]struct {
	id int
	name string
	classname string
	age int
}
type Student struct {
	id int
	name string
	classname string
	age int
}
var(
	StudentArray = make(StudentList,200)
	Studentone Student
)
//生成一个学生信息
func (stu Student)new(){
	fmt.Println("id")
	fmt.Scan(&Studentone.id)
	fmt.Println("姓名")
	fmt.Scan(&Studentone.name)
	fmt.Println("班级")
	fmt.Scan(&Studentone.classname)
	fmt.Println("年龄")
	fmt.Scan(&Studentone.age)
}

//添加信息到用户列表
func (stuall StudentList) add(){
	Studentone.new()
	_,ok := StudentArray[Studentone.id]
	if ok {
		fmt.Println("用户id已经存在！")
		return
	}
	StudentArray[Studentone.id] = Studentone
	fmt.Println()
	fmt.Println("用户添加成功！")
	fmt.Println("信息如下:")
	fmt.Printf("id:%d 姓名:%s 班级:%s 年龄:%d \n",Studentone.id,Studentone.name,Studentone.classname,Studentone.age)
}

//查看所有用户信息
func (stuall StudentList)seeall(){
	fmt.Println()
	for _,v := range StudentArray{
		fmt.Printf("id:%d 姓名:%s 班级:%s 年龄:%d \n",v.id,v.name,v.classname,v.age)
	}
}

//修改信息
func (stuall StudentList) update(){
	Studentone.new()
	_,ok := StudentArray[Studentone.id]
	if !ok {
		fmt.Println("用户id不存在！")
		return
	}
	StudentArray[Studentone.id] = Studentone
	fmt.Println()
	fmt.Println("用户修改成功！")
	fmt.Println("信息如下:")
	fmt.Printf("id:%d 姓名:%s 班级:%s 年龄:%d \n",Studentone.id,Studentone.name,Studentone.classname,Studentone.age)
}
//删除学生信息
func (stull StudentList)delete(){
	var id int
	fmt.Println("请输入要删除用户的id！")
	fmt.Scan(&id)
	_,ok := StudentArray[id]
	if !ok {
		fmt.Println("用户id不存在！")
		return
	}
	delete(StudentArray,id)
	fmt.Println("用户删除成功！")

}

//选择操作
func chose() int{
	fmt.Println("1. 添加学生")
	fmt.Println("2. 删除学生")
	fmt.Println("3. 修改学生信息")
	fmt.Println("4. 展示所有学生")
	fmt.Println("5. 退出")
	fmt.Println()
	var id int
	fmt.Scan(&id)
	return id
}
func main() {
	fmt.Println("欢迎来到学生管理系统")
	fmt.Println()
	for{
		switch chose() {
			case 1:
				StudentArray.add()
				fmt.Println()
			case 2:
				StudentArray.delete()
				fmt.Println()
			case 3:
				StudentArray.update()
				fmt.Println()
			case 4:
				StudentArray.seeall()
				fmt.Println()
			case 5:
				os.Exit(0)
			default:
				fmt.Println("选择参数错误")
		}
	}

}
