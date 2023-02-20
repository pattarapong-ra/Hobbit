package main

import (
	"fmt"
	"time"
)

//base int64
type mathOperator interface {
	Plus() int64
	Minus() int64
	Multi() int64
	Div() int64
}

//base int64
type number struct {
	num1 int64
	num2 int64
}

func (n number) init() number {
	fmt.Printf("input number1 :")
	fmt.Scanf("%d", &n.num1)
	fmt.Printf("input number2 :")
	fmt.Scanf("%d", &n.num2)

	return n
}
func (n number) show() {
	fmt.Println("Both numbers are : ", n.num1, " and ", n.num2)
}
func (n number) Plus() int64 {
	fmt.Println("Answer is ",n.num1 + n.num2)
	return n.num1+n.num2
}
func (n number) Minus() int64 {
	fmt.Println("Answer is ",n.num1 - n.num2)
	return n.num1-n.num2
}
func (n number) Multi() int64 {
	fmt.Println("Answer is ",n.num1 * n.num2)
	return n.num1*n.num2
}
func (n number) Div() int64 {
	fmt.Println("Answer is ",n.num1 / n.num2)
	return n.num1/n.num2
}

func calculator(n mathOperator, c int64) {
	//fmt.Println("++++++++++++++", n)
	switch c {
	case 1:
		n.Plus()
	case 2:
		n.Minus()
	case 3:
		n.Multi()
	case 4:
		n.Div()
/*	case 5:
		n.ChangeNumber()
	case 6:
		n.SwapNumber()
	case 7:
		n.Exit()*/
	default:
		fmt.Println("Input 1-7!!!")
	}
}

func main() {
	n := number{}
	n = n.init()
	for {
		//show input
		n.show()
		//command
		var command int64
		fmt.Println("Enter operation")
		fmt.Println("1. Plus")
		fmt.Println("2. Minus")
		fmt.Println("3. Multi")
		fmt.Println("4. Div")
		fmt.Println("5. Change Number")
		fmt.Println("6. Swap Number") // num1 num2
		fmt.Println("7. Exit")
		fmt.Scanln(&command)
		if command == 5 {
			n = n.init()
		}
		if command == 6 {
			n.num1 += n.num2
			n.num2 = n.num1 - n.num2
			n.num1 = n.num1 - n.num2	// เพิ่มมมมมมม
		}
		if command == 7 {
			break
		}
		calculator(n, command)
		time.Sleep(1 * time.Second)
	}
}
