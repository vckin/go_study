package helloword

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(1,2)
	if sum != 3{
		t.Error("add 1 and 2 resoult not 3")
	}
}
