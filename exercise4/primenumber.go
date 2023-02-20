package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func primeNumber(num int) {
	tempNum := abs(num)
	if tempNum == 1 || tempNum == 0 {
		fmt.Println(num, " is not a prime number.")
	} else if tempNum > 1 {
		for i := 2; i < int(tempNum); i++ {
			if (num % i) == 0 {
				fmt.Println(num, " is not a prime number.")
				return
			}
		}
		fmt.Println(num, " is a prime number.")
	}
}

func main() {
	var number int
	fmt.Printf("Input a number to check :")
	fmt.Scanf("%d", &number)
	primeNumber(number)
}
