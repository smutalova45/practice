package main

import (
	"fmt"
	"math")






func msiol2(){
    array:=[...]int{2,1,3,0,-5}
	countmusbat:=0
	countmanfiy:=0
	countnol:=0
	counttoq:=0
	countjuft:=0
	for i:=0; i<len(array);i++{
		if array[i]>0{
		countmusbat++ 
		}
		if array[i]<0{
			countmanfiy++
		}
		if array[i]==0{
			countnol++
		}
		if array[i]%2==0{
			countjuft++
		}
		if array[i]%2!=0{
			counttoq++
		}
	}
	fmt.Print(countmusbat)
	fmt.Println(countmanfiy)
	fmt.Println(countjuft)
	fmt.Println(countnol)
	fmt.Println(counttoq)

}
func misol3(){
  slice:=[]int{1,2,3,4,5,6,7,8,9,10}
  newSlice1:=make([]int,len(slice))
  newSlice2:=make([]int,len(slice))
  
  for i:=0; i<len(slice);i++{
	if slice[i]%2==0{
		newSlice1[i]=slice[i]

		
	}
	if slice[i]%2!=0{
		newSlice2[i]=slice[i]
		
	}
  }
  fmt.Println(newSlice1)
  fmt.Println(newSlice2)
}
func misol4(){
	slice:=[]int{1,2,3,4,5,6,7,8,9,10}
	multiple:=1
	
	
	summ:=0
	for _, num:=range slice{
		if num%2==0{
			multiple=multiple*num
			
		}
		if num%2!=0{
			summ+=num
		}
	}
	fmt.Print("kopaytma ",multiple)
	fmt.Print("summa",summ)
}
func misol5(){
	slice:=[]int{1,2,3,4,5,6,7,8,9,10,11}
    
	 min1 := math.MaxInt64
	 min2 := math.MaxInt64
	 
	 for _, num := range slice {
	  if num < min1 {
	   min2 = min1
	   min1 = num
	  } else if num < min2 && num != min1 {
	   min2 = num
	  }
	 }
	 
	 fmt.Println("Ikkinchi minimum element :", min2)
}
func misol6(){   // 0 1  2  3  4 5   6  7  8  9   10
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
 
	max1 := slice[0] 
	max2 := slice[0]
	
	for _, value:= range slice {
	 if value > max1 { // 1>0 max1=2
	  max2 = max1 // max2=0  max2=2
	  max1 = value // max1=1
	 } else if value > max2 { // 1>0
	  max2 =value // max2=1  
	 }
	}
	
	fmt.Println("Ikkinchi maksimum element: ", max2)
   }

   func misol7(){
	slice:=[]int {1,2,3,4,5,6,7,8}
	fmt.Println(slice)
	

	temp:=slice[0]
	slice[0]=slice[len(slice)-1]
	slice[len(slice)-1]=temp
	fmt.Println(slice)
	
   }
   func misol8(){
	var count int=0
	slice:=[]int {2,9,3,-4,5}
	var max int=slice[0]
	for i:=0;i<len(slice);i++{
       if max<slice[i]{
		max=slice[i]
	   }
	
	for i:=max;i<len(slice);i++{
		count++
	}
}
    fmt.Printf("maximum %d dan keyin %d ta son bor",max ,count)
	fmt.Println()

   }
   func misol9(){
	var count int=0
	slice:=[]int {2,9,3,-4,5}
	var min int=slice[0]
	for i:=0;i<len(slice);i++{
       if min>slice[i]{
		min=slice[i]
	   }
	
	for i:=min;i<len(slice);i++{
		count++
	}
}
    fmt.Printf("minimum %d dan keyin %d ta son bor",min ,count)
	fmt.Println()


   }
   func misol10(){
	
	slice:=[]int {2,-4,3,9,5}
	var min int=slice[0]
	for i:=0;i<len(slice);i++{
       if min>slice[i]{
		min=slice[i]
	   }
     }
      var max int=slice[0]
   for i:=0;i<len(slice);i++{
	  if max<slice[i]{
	   max=slice[i]
	  }
	}
	fmt.Println(max)
	fmt.Println(min)
	
		rangeslice:=slice[min:max]
		fmt.Println(rangeslice)
	}

// func fib(n int)int{
// 	a:=0 
// 	b:=1
// 	for i:=0;i<n;i++{
// 		temp:=a+b 
// 		a=b 
// 		b=temp
// 	}
// 	return a
// }

// func triangle(n int){
// 	r:=" "
// 	for i:=n;i>0;i--{
// 		for k:=0;k<n;k++{
// 			r+=" "
			
// 		}
// 		for j:=1;j<=i;j++{
// 			r+="*"
// 		}
// 		fmt.Println(r)
// 		r=" "
// 	}
// }

func main() {

	//Praktika
    // msiol2()
    //  misol3()
	//  misol4()

	// misol5()
    //  misol6() 
	// misol7()
	// misol8()
	// misol9()
	// misol10()











	
	// var n int 
	// fmt.Scan(&n)
	// triangle(n)
	// fmt.Println(fib(n))


	// var(
	// 	a ,b ,c int
	// 	max , min int 
		
	// )


	// fmt.Scan(&a)
	// fmt.Scan(&b)
	// fmt.Scan(&c)

	// max=a
	// if max<b{
	// 	max=b
		
	// } else if max<c{
	// 	max=c 
		
	// }
    // fmt.Println("max: ",max)

	// min=a
	// if min>b{
	// 	min=b
		
	// }else if min>c{

	// 	min=c
		
	// }

	// fmt.Println("min: ",min)
	// fmt.Println("sum: ",min+max)



	














	// finder := 0
	// a := 0
	// b := 0
	// var sum int = 0

	// var slice1 = []int{1, -2, 3, -4, 5, -6, 7, -8, 9, -10}
	// for i := 0; i < len(slice1); i++ {
	// 	if slice1[i] < 0 {
	// 		finder++
	// 		if finder == 2 {

	// 			a = i

	// 		}
	// 		if finder == 4 {
	// 			b = i
	// 			break
	// 		}
	// 	}

	// }
	// for j := a + 1; j < b; j++ {

	// 	sum += slice1[j]

	// }
	// fmt.Println(sum)

}
