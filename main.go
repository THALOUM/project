package main

import (
	"fmt"

	puppy "github.com/THALOUM/Puppy"
)

func main() {
	s1 := puppy.Bark()
	fmt.Println(s1)

	s2 := puppy.BigBarks()
	fmt.Println(s2)
}
