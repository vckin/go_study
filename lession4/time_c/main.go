package main

import (
	"fmt"
	"net/http"
	"time"
)

func FormatTime()  {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:03:04"))
	fmt.Println(now.Format("2006-01-02 15:03:04.999"))
	fmt.Println(now.Format("15:03:04.999"))
}
func httpTime()  {
	fmt.Println(http.TimeFormat)
}
func main() {
	FormatTime()
	httpTime()
}
