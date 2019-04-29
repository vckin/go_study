package repeat

const repeatCount = 5

func Repeat(s string)string  {
	s1 := ""
	for i := 0;i< repeatCount ;i++{
		s1 +=  s
	}
	return s1
}
