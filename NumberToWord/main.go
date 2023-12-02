package main

import "fmt"

func main() {
	var number int
	fmt.Println("enter number: ")
	fmt.Scan(&number)
	if number < 0 {
		number = -1 * number
		fmt.Print("-")
	}

}
func NumberToWord(number int) {
	var (
		dig        int = 0
		tenPower   int = 1
		workNumber int = number
	)
	for workNumber != 0 {
		workNumber = number / 10
		dig++
	}
	for i := 1; i < dig; i++ {
		tenPower += workNumber
	}
	for dig >= 1 {
		var s string
		workNumber = number / tenPower
		if dig%3 == 2 {
			switch workNumber {
			case 1:
				s = "o'n"
				fmt.Print(s)
			case 2:
				s = "yigirma"
				fmt.Print(s)
			case 3:
				s = "o'ttiz"
				fmt.Print(s)
			case 4:
				s = "qirq"
				fmt.Print(s)
			case 5:
				s = "ellik"
				fmt.Print(s)
			case 6:
				s = "oltmish"
				fmt.Print(s)
			case 7:
				s = "yetmish"
				fmt.Print(s)
			case 8:
				s = "sakson"
				fmt.Print(s)
			case 9:
				s = "to'qson"
				fmt.Print(s)

			}
		} else {
			switch workNumber {
			case 1:
				s = "bir"
				fmt.Print(s)
			case 2:
				s = "ikki"
				fmt.Print(s)
			case 3:
				s = "uch"
				fmt.Print(s)
			case 4:
				s = "qirq"
				fmt.Print(s)
			case 5:
				s = "ellik"
				fmt.Print(s)
			case 6:
				s = "oltmish"
				fmt.Print(s)
			case 7:
				s = "yetmish"
				fmt.Print(s)
			case 8:
				s = "sakson"
				fmt.Print(s)
			case 9:
				s = "to'qson"
				fmt.Print(s)

			}

		}
		if dig%3 == 0 && workNumber != 0 {
			fmt.Print("yuz")
		}
		if workNumber != 0 {
			switch dig {
			case 4:
				s = "ming"
				fmt.Print(s)
			case 7:
				s = "million"
				fmt.Print(s)
			case 10:
				s = "trillion"
				fmt.Print(s)
			}
		}

		number %= tenPower
		tenPower /= 10
		dig--

	}

}
