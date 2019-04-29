package sumall

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(l ...[]int)(sums []int){
	sums = make([]int, len(l))
	for i,num := range l{
		sums[i] = Sum(num)
	}
	return
}
