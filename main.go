package main

import (
	"fmt"

	"github.com/theterminalguy/xo/set"
)

func main() {
	s := new(set.Set)
	s.Add("a")
	s.Add("b")
	s.Add("b")
	s.Add("c")

	s.Remove("c")

	s.Add("Simon")
	s.Add("Peter")

	fmt.Println(s.Members(), s.Size())

	fmt.Println("Set contains Simon?", s.Contains("Simon"))

	//arr := []string{"a", "b", "c"}

	//fmt.Println(arr[2+1:])
}
