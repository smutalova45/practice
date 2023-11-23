package main

import "fmt"

func tubSon(son int) {
	counter := 0
	for j := 2; j <= son/2; j++ {
		if son%j == 0 {
			counter++
		}
	}
	if counter == 0 {
		fmt.Println(son, " tub son.")
	} else {
		fmt.Println(son, " tub son emas.")
	}
}

func swaping(num int) {

	var temp int

	for num != 0 {
		temp = num % 10
		fmt.Printf("%d ", temp)
		num /= 10
	}
}

func sum(s int) {
	var count = 0
	temps := 0
	temps = s
	for temps != 0 {
		temps /= 10
		count++
	}
	temp := 1
	sum := 0
	for i := 0; i < count; i++ {
		if i == 0 {
			for j := 1; j < count; j++ {
				temp *= 10
			}
		}
		sum += s / temp
		s %= temp
		temp /= 10
	}
	fmt.Println(sum)

}

// 5
func printout(a, b int) {
	if a < b {
		for i := a; i <= b; i++ {
			fmt.Print(i, " ")
		}
	} else if a > b {
		for j := a; j >= b; j-- {
			fmt.Print(j, " \n")
		}
	}
	fmt.Println()
}

func juftson(n int) {
	for i := n; i > 1; i-- {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
func middle(a, b, c int) {
	if a > c && b < c {
		fmt.Println(c)
	} else if a > b && c < b {
		fmt.Println(b)
	} else if b > a && c < a {
		fmt.Println(a)
	} else if b > c && a < c {
		fmt.Println(c)
	} else if c > b && a < b {
		fmt.Println(b)
	} else {
		fmt.Println(a)
	}

}

func evenSum(n int) bool {
	tempN := 0
	tempN = n
	count := 0
	for tempN != 0 {
		tempN /= 10
		count++
	}
	var temp = 1
	var sum = 0
	for i := 0; i < count; i++ {
		if i == 0 {
			for j := 1; j < count; j++ {
				temp *= 10
			}
		}
		sum += n / temp

		n %= temp
		temp /= 10
	}
	if sum%2 == 0 {
		return true
	} else {
		return false
	}

}

func dividenum(n, n1 int) int {
	return n / n1
}

func sumnumbers(n int) int {
	sum := 0
	for i := 1; i < n+1; i++ {
		sum += i
	}

	return sum
}
func algorithm(a int) {
	temp := 0
	temp = a
	count := 0
	for temp != 0 {
		temp = temp / 10
		count++
	}

	tempN := 1
	for i := 0; i < count; i++ {
		if i == 0 {
			for j := 1; j < count; j++ {
				tempN *= 10
			}
		}
		fmt.Print(a/tempN, " ")

		a %= tempN
		tempN /= 10
	}
	fmt.Println()
}
func primeDivider(num int) {
	for i := 2; i <= num; i++ {
		var sum int = 0
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				sum++
			}
		}
		if sum == 0 {
			if num%i == 0 {
				fmt.Print(i, " ")
			}
		}
	}
	fmt.Println()
}
func multiplication(m int) int {
	m2 := 1
	for i := 1; i <= m; i++ {
		m2 *= i
	}
	return m2
}

func adding(a, b, c int) int {
	if a+b == c {
		return 1
	} else {
		return 0
	}
}

// task-15
func pwN(n float32, m int) float32 {
	for i := 1; i < m; i++ {
		n *= n
	}
	return n
}

func toqSon(n int) {
	for i := 1; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func jadval(jadval int) {
	for i := 1; i <= jadval; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println()
}
func main() {

	// var son int
	// fmt.Scan(&son)
	// tubSon(son)

	// var num int
	// fmt.Scan(&num)
	// swaping(num)

	// var s int
	// fmt.Scan(&s)
	// sum(s)

	// //5
	// var a, b int
	// fmt.Scan(&a,&b)
	// printout(a,b)

	// var juftsonga int
	// fmt.Scan(&juftsonga)
	// juftson(juftsonga)

	//   var (
	// 	a ,b , c int
	//   )
	//   fmt.Scan(&a,&b,&c)
	//   middle(a,b,c)

	//8

	// var n int
	// fmt.Scan(&n)
	// evenSum(n)

	// 9
	// var (
	// 	a,b int
	// )
	// fmt.Scan(&a,&b)
	// dividenum(a,b)

	//10
	// var n int
	// fmt.Scan(&n)
	// sumnumbers(n)

	//11
	// var a int
	// fmt.Scan(&a)
	// algorithm(a)

	//12
	// 	var num int
	// 	fmt.Scan(&num)
	// primeDivider(num)

	// var m2 int
	// fmt.Scan(&m2)
	// multiplication(m2)

	// var (
	// 	a ,b,c int
	// )
	// fmt.Scan(&a,&b,&c)
	// adding(a,b,c)

	// var(
	// 	m float32
	// 	n int
	// )
	// fmt.Scan(&m,&n)
	// pwN(m,n)

	// var n int
	// fmt.Scan(&n)
	// toqSon(n)

	var n int
	fmt.Scan(&n)
	jadval(n)

}
