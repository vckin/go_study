package main

import "fmt"

func main() {
	studentMap := []map[string]string{
		{
			"username": "a",
			"age":"18",
			"grade":"60",
		},
		{
			"username": "b",
			"age":"19",
			"grade":"70",
		},
		{
			"username": "d",
			"age":"20",
			"grade":"80",
		},
		{
			"username": "h",
			"age":"21",
			"grade":"90",
		},
	}

	studentInfo := map[string]string{
			"username": "d",
			"age":"17",
			"grade":"20",
	}
	studentMap = append(studentMap,studentInfo)
	fmt.Println(studentMap)
	fmt.Println(studentMap[1]["age"])
}
