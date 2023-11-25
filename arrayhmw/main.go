package main

import (
	"fmt"

	"math/rand"
	"time"
)

// array := []int{2, 3, 4, 5, 6, 7, 8, 9, 10


func misol1(){

	slice:=[]int {1,26,7,8,123,0,-4}
	max :=slice[0]
	maxIndex := 0
    minIndex := 0
	
	for i,value:=range slice{
		if max<value{
			max=value
			maxIndex=i
		}
	
	}

	min:=slice[0]
	for i, value:=range slice{
		if min>value{
			min=value
			maxIndex=i
		}
	}


	fmt.Println(slice)
	fmt.Println("max",max)
	fmt.Println("min",min)
	fmt.Println("after swaping")
	slice[maxIndex],slice[minIndex]=slice[minIndex],slice[maxIndex]
	fmt.Println(slice)



}
//2 misol
func isPrime(num int) bool {
 if num <= 1 { 
  return false
 }

 for i := 2; i*i <= num; i++ {
  if num%i == 0 {
   return false
  }
 }

 return true
}

func findPrimes(arr []int) []int {
 var primes []int

 for _, num := range arr {
  if isPrime(num) {
   primes = append(primes, num)
  }
 }

 return primes
}
 func misol3(){
	slice:=[]int {1,2,3,4,5,6}
	index:=0
	sumindex:=0
	sum:=0
	for i,value:=range slice{
		sumindex=index+i
		sum+=value

		
	}
	if sumindex>sum{
		fmt.Println("INDEX")
	}else if sumindex<sum{
		fmt.Println("Values")
	}else {
		fmt.Println("wow")
	}
	
 }
 func misol4(n int){
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, n)
    for i := 0; i < n; i++ {
     slice[i] = rand.Intn(26) - 12
    }
	fmt.Println(slice)

	for i:=0; i<len(slice);i++{
      if slice[i]<0{
		slice[i]=0
	  }
	  if slice[i]>0{
		slice[i]=1
	  }
	}
	fmt.Println(slice)


}

func misol5(n int){
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, n)
    for i := 0; i < n; i++ {
     slice[i] = rand.Intn(35) +14
    }
	fmt.Println(slice)
	for i:=0; i<len(slice);i++{
		if slice[i]>0{
			fmt.Print(slice[i],"+")
		}
	}


}
func misol6(n int){

rand.Seed(time.Now().UnixNano())
	slice := make([]int, n)
    for i := 0; i < n; i++ {
     slice[i] = rand.Intn(5) -46
    }
	fmt.Println(slice)
	for i:=0; i<len(slice);i++{
		if slice[i]%2!=0{
			fmt.Print("_",slice[i])
		}
}
}
func misol7(n int){
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, n)
    for i := 0; i < n; i++ {
     slice[i] = rand.Intn(25) -5
    }
	fmt.Println(slice)
	temp:=slice[0]
	for i:=0; i<len(slice)-1;i++{

		slice[i]=slice[i+1]
		slice[i-1]=temp



	}
	fmt.Println(slice)

}

func misol9(n int){
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, n)
    for i := 0; i < n; i++ {
     slice[i] = rand.Intn(8) -15
    }
	fmt.Println(slice)
	temp:=slice[0]
	for i:=0; i<len(slice)-1;i++{
		slice[i]=slice[i+1]
		slice[len(slice)-1]=temp

	}
	fmt.Println(slice)

}
func misol10(n , k int){
 
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, n)
    for i := 0; i < n; i++ {
     slice[i] = rand.Intn(8) +21
    }
	fmt.Println(slice)
	temp:=slice[0]
	for i:=0; i<len(slice)-1;i++{
		slice[i]=slice[i+k]
		slice[len(slice)-1]=temp

	}
	fmt.Println(slice)

}
func misol11(n,k int){
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, n)
    for i := 0; i < n; i++ {
     slice[i] = rand.Intn(70) -11
    }
	fmt.Println(slice)
	temp:=slice[0]
	for i:=0; i<len(slice)-1;i++{
		slice[i]=slice[i+k]
		slice[k-1]=temp

	}
	fmt.Println(slice)

}
func main(){
	// rand.Seed(time.Now().UnixNano())
	//misol4 
	// var n int 
	// fmt.Scan(&n)
	// misol4(n)
	// var n int 
	// fmt.Scan(&n)
	// misol5(n)
	// var n int 
	// fmt.Scan(&n)
	// misol6(n)

// misol1()
//misol2
// numbers := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}

// primes := findPrimes(numbers)

// fmt.Println("Prime numbers in the array:", primes)
//   misol3()


 





}