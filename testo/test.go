package main

import (
	"fmt"
	"math"
)

func main() {
	message := greetMe("world")
	fmt.Println(message)
	s := 35000/((1-(1/ (math.Pow(1+0.09120/12,4)) ))/(0.09120/12))
	fmt.Println(s)
}

func greetMe(name string) string {
	return "Hello, " + name + "!"
}
