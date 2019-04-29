package main

import (
	"github.com/vckin/go_study/lession4/mylog/log_time/log"
)




// 写了一个项目想要在代码中记录日志
// 要使用mylog这个包
func main() {
	//fmt.Println(mylog.DEBUG)
	fl := mylog.NewFileLogger(mylog.WARN, "./runtime", "test.log")
	for i:=0;i<50;i++{

		userID := 10
		fl.Warn("id是%d的用户一直在尝试登陆", userID)
	}
}
