package main

import (
	"fmt"
)

type Numbers struct {
	x int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (num Numbers) check() string {
	tempNum := abs(num.x)
	if tempNum == 1 || tempNum == 0 {
		fmt.Println(num, " is not a prime number.")
	} else if tempNum > 1 {
		for i := 2; i < int(tempNum); i++ {
			if (tempNum % i) == 0 {
				return "This is not a prime number."
			}
		}
	}
	return "This is a prime number."
}

func main() {
	var inputNumber int
	fmt.Printf("Input a number to check :")
	fmt.Scanf("%d", &inputNumber)
	numberStruct := Numbers{inputNumber}
	fmt.Println(numberStruct.check())
}

/*func(reciver_name Type) method_name(parameter_list)(return_type){
// Code
}*/
