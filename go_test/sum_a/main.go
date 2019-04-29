package sum_a

func Sum(numbers [5]int)int{
	num := 0
	for _,i := range numbers{
		num += i
	}
	return num
}
