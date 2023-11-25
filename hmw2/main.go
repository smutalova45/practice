package main

import (
	"fmt"
)

// func kordinata(x1,x2,y1,y2 float64, c float64){

// 	fmt.Scan(&x1)
// 	fmt.Scan(&x2)
// 	fmt.Scan(&y1)
// 	fmt.Scan(&y2)

// 	a:=x1-x2
// 	b:=y1-y2

// 	k:=float64(math.Pow(2,a)) + float64((math.Pow(2,b)))
// 	fmt.Println(math.Sqrt(k))

// }
// func fibonacci(n  int) int{
// 	f0:=0
// 	f1:=1
// 	f2:=1

// 	fmt.Scan(&n)
// 	f2=f1+f0
//fib int
// 	fib(n):=f(n-1)+f(n-2)
// 	return fib
// }

// func main()  {

// for i:=1;i<=5;i++{

// 	for j:=1;j<=5;j++{
// 		if j>=i{
// 		fmt.Print("*")
// 		}else {
// 			fmt.Print(" ")
// 		}

//    }
//    fmt.Println(" ")
// }

// for i:=1; i<=5; i++{
// 	for j:=1; j<=5; j++{
// 		if j<=i{
// 			fmt.Print(" *")
// 		}else{
// 			fmt.Print(" ")
// 		}
// 	}
// 	fmt.Println(" ")
// }

// var(
// 	a float64
// 	b float64
// 	sum float64
// )

// fmt.Scan(&a)
// fmt.Scan(&b)

// max := math.Max(a,b)

// min:= math.Min(a,b)
// sum = max+min
// fmt.Printf("max %f , min %f, sum %f ", max, min, sum)
//}

// homework
func xonalison(n int) string {
	var s string
	if n < 10 && n > 0 {
		s = "1 xonali"
		fmt.Println(s)
	} else if n >= 10 && n < 100 {
		s = "2 xonali"
		fmt.Println(s)
	} else if n >= 100 && n < 1000 {
		s = "3 xonali"
		fmt.Println(s)
	} else if n >= 1000 && n < 10000 {
		s = "4 xonali"
		fmt.Println(s)
	}
	return s
}
func add(a, b int) int {
	digit := 1
	for i := b; i > 0; i /= 10 {
		digit *= 10
	}
	result := a*digit + b
	return result
}
// func factorial(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	return n * factorial(n-1)
// }

func reverse(n int) int {
	result := 0
	var k int
	for n > 0 {
		k = n % 10
		result *= 10
		result += k

		n /= 10
	}
	return result
}

func son(n int) string {
 var s string
	if n > 0 && n < 10 {
		var k int
		k = (n / 1) % 10
		switch k {

		case 1:
			s = "bir "
			fmt.Print(s)
		case 2:
			s = "ikki "
			fmt.Print(s)
		case 3:
			s = "uch "
			fmt.Print(s)
			break
		case 4:
			s = "to'rt  "
			fmt.Print(s)
		case 5:
			s = "besh "
			fmt.Print(s)
		case 6:
			s = "olti "
			fmt.Print(s)
		case 7:
			s = "yetti "
			fmt.Print(s)
		case 8:
			s = "sakkiz "
			fmt.Print(s)
		case 9:
			s = "to'qqiz "
			fmt.Print(s)
		}
	}else if n>9 && n<100{
		var k int
	
		k = (n / 100) % 10
		switch k {
		case 1:
			s = "on "
			fmt.Print(s)
		case 2:
			s = "yigirma"
			fmt.Print(s)
			
		case 3:
			s = "ottiz "
			fmt.Print(s)
		case 4:
			s = "qirq "
			fmt.Print(s)
		case 5:
			s = "ellik"
			fmt.Print(s)
			
		case 6:
			s = "oltmush "
			fmt.Print(s)
		case 7:
			s = "yettimush"
			fmt.Print(s)
		case 8:
			s = "sakson "
			fmt.Print(s)
		case 9:
			s = "to'qson"
			fmt.Print(s)

		k = (n / 1) % 10
		switch k {

		case 1:
			s = "bir "
			fmt.Print(s)
		case 2:
			s = "ikki "
			fmt.Print(s)
		case 3:
			s = "uch "
			fmt.Print(s)
			break
		case 4:
			s = "to'rt  "
			fmt.Print(s)
		case 5:
			s = "besh "
			fmt.Print(s)
		case 6:
			s = "olti "
			fmt.Print(s)
		case 7:
			s = "yetti "
			fmt.Print(s)
		case 8:
			s = "sakkiz "
			fmt.Print(s)
		case 9:
			s = "to'qqiz "
			fmt.Print(s)
		}
		}
    } else if n>99 && n < 1000 {
		var k int
		k = (n / 100) % 10
		switch k {
		case 1:
			s = "bir yuz "
			fmt.Print(s)
		case 2:
			s = "ikki yuz "
			fmt.Print(s)
			
		case 3:
			s = "uch yuz "
			fmt.Print(s)
		case 4:
			s = "to'rt yuz"
			fmt.Print(s)
		case 5:
			s = "besh yuz "
			fmt.Print(s)
			
		case 6:
			s = "olti yuz "
			fmt.Print(s)
		case 7:
			s = "yetti yuz "
			fmt.Print(s)
		case 8:
			s = "sakkiz yuz "
			fmt.Print(s)
		case 9:
			s = "to'qqiz yuz"
			fmt.Print(s)
		}

		k = (n % 100) / 10 //490
		switch k {
		case 1:
			s = "on "
			fmt.Print(s)
		case 2:
			s = "yigirma"
			fmt.Print(s)
			
		case 3:
			s = "ottiz "
			fmt.Print(s)
		case 4:
			s = "qirq "
			fmt.Print(s)
		case 5:
			s = "ellik"
			fmt.Print(s)
			
		case 6:
			s = "oltmush "
			fmt.Print(s)
		case 7:
			s = "yettimush"
			fmt.Print(s)
		case 8:
			s = "sakson "
			fmt.Print(s)
		case 9:
			s = "to'qson"
			fmt.Print(s)


		} 
		k = (n / 1) % 10
		switch k {

		case 1:
			s = "bir "
			fmt.Print(s)
		case 2:
			s = "ikki "
			fmt.Print(s)
		case 3:
			s = "uch "
			fmt.Print(s)
			break
		case 4:
			s = "to'rt  "
			fmt.Print(s)
		case 5:
			s = "besh "
			fmt.Print(s)
		case 6:
			s = "olti "
			fmt.Print(s)
		case 7:
			s = "yetti "
			fmt.Print(s)
		case 8:
			s = "sakkiz "
			fmt.Print(s)
		case 9:
			s = "to'qqiz "
			fmt.Print(s)
		}



		

	}

	


return s
}

func main() {

	//1
	// var n int
	// fmt.Scan(&n)
	// xonalison(n)

	//2
	// var (
	// 	a ,b int
	// )
	//   fmt.Scan(&a)
	//   fmt.Scan(&b)
	//   fmt.Print(add(a,b))

	//3
	// var n int
	// fmt.Scan(&n)
	// factorial(n)

	// 4 reverse

	var n int
	fmt.Scan(&n)
	fmt.Print(reverse(n))


	//4 
	// var  n int 
	// fmt.Scan(&n)
	// son(n)

}
