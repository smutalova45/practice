package main

import (
	"fmt"
)

// type Car struct {
// 	Name       string
// 	Price      float64
// 	Model      string
// 	HorsePower uint32
// }
//1
type Student struct {
	Name string 
	Age int 
	Scholarship int 
	Course int
	// Score []int
}
 type Aeroport struct{
	PlaneType  string
	FlightDate int 
	FromCity string 
	ToCity string 
	FlightTime int

 }
 type Team struct{
	Name string 
	Coach string  
	PlayersCount int 
	Captain string 



 }





func main() {


// ///Homework of struct//1
// 	students:=[]Student{
// 	{
// 		Name:"BillGates" ,
// 		Age: 20,
// 		Scholarship: 12000,
// 		Course: 3,

// 	},	
// 	{
//         Name:"TomKKK" ,
// 		Age: 20,
// 		Scholarship: 11000,
// 		Course: 2,
// 	},
// 	{
// 		Name:"Sara" ,
// 		Age: 20,
// 		Scholarship: 12000,
// 		Course: 1,
// 	},
// 	{
// 		Name:"BA" ,
// 		Age: 20,
// 		Scholarship: 9000,
// 		Course: 2,
// 	},
// 	{
// 		Name:"Bill" ,
// 		Age: 18,
// 		Scholarship: 12000,
// 		Course: 1,
// 	},
// 	{

// 		Name:"II" ,
// 		Age: 22,
// 		Scholarship: 12000,
// 		Course: 4,
// 	},
// 	{
// 		Name:"Tompson" ,
// 		Age: 20,
// 		Scholarship: 5000,
// 		Course: 2,

// 	},
// 	}
// // 	var sum int=0
// // 	for i:=0;i<len(students);i++{
		
// // 		if students[i].Course==2{
// //             sum+=students[i].Scholarship
// // 		}
// // 	}
// // 	fmt.Println("total scholarship ",sum)
// // 	//2 homework

// // 	for i:=0;i<len(students);i++{
// // 		if len(students[i].Name)<5{
// // 			fmt.Println(students[i].Name)

// // 		}
// // 	}

//3 
//    a:=[]Aeroport{
// 	{
// 	PlaneType:"Something",
// 	FlightDate : 7,
// 	FromCity :"USA",
// 	ToCity :"Uzb",
// 	FlightTime :4,
        
// 	},
// 	{
// 		PlaneType:"JJJJJ",
// 		FlightDate : 8,
// 		FromCity :"USA",
// 		ToCity :"Uzb",
// 		FlightTime : 2,
			
//    },
//    {
// 	PlaneType:"Something",
// 	FlightDate : 9,
// 	FromCity :"USA",
// 	ToCity :"Uzb",
// 	FlightTime : 9,
        
// 	},

// 	{
// 		PlaneType:"XXX",
// 		FlightDate : 8,
// 		FromCity :"Latvia",
// 		ToCity :"Tashkent",
// 		FlightTime : 3,
			
// 	},
// 	{
// 		PlaneType:"BBB",
// 		FlightDate : 6,
// 		FromCity :"USA",
// 		ToCity :"Uzb",
// 		FlightTime : 5,
			
// 	},
// 	{
// 		PlaneType:"Something",
// 		FlightDate : 12,
// 		FromCity :"USA",
// 		ToCity :"Tashkent",
// 		FlightTime : 2,
			
// 	},
// 	{
// 		PlaneType:"AAA",
// 		FlightDate : 4,
// 		FromCity :"UAA",
// 		ToCity :"Tashkent",
// 		FlightTime : 3,
			
// 	},
// 	{
// 		PlaneType:"PPPPP",
// 		FlightDate : 4,
// 		FromCity :"Kazakhistan",
// 		ToCity :"Tashkent",
// 		FlightTime : 3,
			
// 	},
//    }
    // for i:=0;i<len(a);i++{
    //  if a[i].FlightDate>=6 && a[i].FlightDate<=8{
	// 	fmt.Println(a[i])
	//  }
	// }
 
	// for i:=0;i<len(a);i++{
	// 	if a[i].FlightTime==2 || a[i].FlightTime==3{
	// 		fmt.Println(a[i])
	// 	}
	// }
	// for i:=0;i<len(a);i++{
	// 	if a[i].ToCity=="Tashkent"{
	// 		fmt.Println(a[i])
	// 	}
	// }



teams:=[]Team{
	{
       Name:"AAAA",
	   Coach:"aaaaa",
	   PlayersCount: 15,
	   Captain: "qqqq",
	},
	{
		Name:"A",
		Coach:"aaaaa",
		PlayersCount: 18,
		Captain: "qqqq",
	 },
	 {
		Name:"BBBB",
		Coach:"aaaaa",
		PlayersCount: 16,
		Captain: "qqqq",
	 },
	 {
		Name:"CCCC",
		Coach:"aaaaa",
		PlayersCount: 17,
		Captain: "qqqq",
	 },
	 {
		Name:"FFFFF",
		Coach:"aaaaa",
		PlayersCount: 9,
		Captain: "qqqq",
	 },
	 {
		Name:"AAANNN",
		Coach:"aaaaa",
		PlayersCount: 20,
		Captain: "qqqq",
	 },
}

// slice1:=[]string{}
fmt.Print("enter name of team ")
var nameInput string 
fmt.Scan(&nameInput)
fmt.Println()
for i:=0;i<len(teams);i++{
    if teams[i].Name==nameInput{
		fmt.Println(teams[i])
		break
		
	}else{
		fmt.Println("not found ")
		
	}
}























	//practice in lecture
	// students:=[]Student{

	// 	{
	// 		Name: "AA",
	// 		Age: 10,
	// 		Scholarship: 2000,
	// 		Score: []int {1,2,3,4,5},
	// 	},
	// 	{
	// 		Name: "CC",
	// 		Age: 10,
	// 		Scholarship: 2000,
	// 		Score: []int {1,2,3,4,5},

	// 	},
	// 	{
	// 		Name: "BB",
			// Age: 10,
	// 		Scholarship: 2000,
	// 		Score: []int {4,4,4,4,4},
	// 	},

		
	// }
	// var sum int=0
	// for i:=0;i<len(students);i++{

	// 	for j:=0;j<len(students[i].Score);j++{
    //       sum+= students[i].Score[j]
                
	// 	}
	// 	sum/=len(students[i].Score)
	// 	if sum>=4{
	// 		fmt.Println(students[i])
	// 	}
	// }
	
	

	//1 practise
	// cars := []Car{

	// 	{
	// 		Name:       "VUT",
	// 		Price:      12.000,
	// 		Model:      "BBB",
	// 		HorsePower: 400,
	// 	},
	// 	{
	// 		Name:       "ADT",
	// 		Price:      12.000,
	// 		Model:      "BHJ",
	// 		HorsePower: 400,
	// 	},
	// 	{
	// 		Name:       "CWT",
	// 		Price:      12.000,
	// 		Model:      "BHJ",
	// 		HorsePower: 400,
	// 	},
	// }

	// for i := 0; i < len(cars); i++ {
	// 	for j := i + 1; j < len(cars); j++ {
	// 		if cars[i].Name > cars[j].Name {
	// 			cars[j].Name,cars[i].Name=cars[i].Name,cars[j].Name
	// 		}
	// 	}
	// }

	// fmt.Print(cars)

	// words:=[]string{"A","D","B"}
	// for i:=0;i<len(words);i++{
	// 	for j:=i+1;j<len(words)-1;i++{

	//            for j=i+1;j<len(words);j++{
	// 			if words[i]>words[j]{
	// 				words[j],words[i]=words[i],words[j]
	// 			}
	// 		   }

	// 	}
	// }
}
