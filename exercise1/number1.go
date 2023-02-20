package main

import (
	"fmt"
	"strconv"
)

func number(S string) string {

	return S
}

func main() {
	tr := "h99 ส!"
	ansslice := []string{}
	numberslice := []string{}
	specialslice := []string{}
	englishslice := []string{}
	thaislice := []string{}
	numbersslice := []string{}hgf
	for i := 0; i < len(tr); i++ {
		temp := rune(tr[i])
		fmt.Println(i)
		switch {
		case (temp >= 65 && temp <= 90) || (temp >= 97 && temp <= 122): //english character
			englishslice = append(ansslice, "A,B,C,D,E,F,G,H,I,J")
		case (temp >= 161 && temp <= 251): //thai character
			thaislice = append(ansslice, "ก,ข,ค,ง,จ")
		case (temp >= 48 && temp <= 57): //number
			if rune(tr[i+1]) >= 48 && rune(tr[i+1]) <= 57 {
				numberslice = append(numberslice, string(tr[i]))
				for j := 1; (i + j) < len(tr); j++ {
					if rune(tr[i+j]) >= 48 && rune(tr[i+j]) <= 57 {
						numberslice = append(numberslice, string(tr[i]))
					} else {
						sum := 0
						for k := 0; k < len(numberslice); k++ {
							l, _ := strconv.Atoi(numberslice[k])
							sum += l
						}
						fmt.Println(sum)
						sumString := strconv.Itoa(sum)
						tempchars := []rune(sumString)
						for i := len(tempchars) - 1; i >= 0; i-- {
							char := string(tempchars[i])
							numbersslice = append(numbersslice, char)
						}
						break
					}
				}
			}
		default:
			specialslice = append(specialslice, string(tr[i]))
		}

		//fmt.Println(ansslice)

	}
	ansslice = append(ansslice, specialslice...)
	ansslice = append(ansslice, englishslice...)
	ansslice = append(ansslice, thaislice...)
	ansslice = append(ansslice, numbersslice...)
	fmt.Printf("%+v", ansslice)
}
