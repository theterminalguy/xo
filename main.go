package main

import (
	"fmt"
	"log"

	"github.com/theterminalguy/xo/set"
)

func main() {
	s := new(set.Set)
	s.Add("a")
	_, err := s.Add(1)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(s.Members())
}
