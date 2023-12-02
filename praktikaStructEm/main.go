package main

import "fmt"

type Client struct {
	ID     int
	Name   string
	Basket []CardProducts
}
type CardProducts struct {
	ProductID int
	SizeID    int
	Quantity  int
}
type Product struct {
	IDproduct int
	Name      string
	size      []Size
}
type Size struct {
	IDsize int
	Name   string
	Price  int
}

var clients = []Client{
	{ID: 1, Name: "Akhmedov", Basket: []CardProducts{
		{ProductID: 1, SizeID: 2, Quantity: 4},
		{ProductID: 4, SizeID: 1, Quantity: 3},
		{ProductID: 3, SizeID: 3, Quantity: 1},
	},
	},
	{
		ID: 2, Name: "Komilov", Basket: []CardProducts{
			{ProductID: 2, SizeID: 2, Quantity: 1},
			{ProductID: 2, SizeID: 2, Quantity: 6},
			{ProductID: 5, SizeID: 1, Quantity: 5},
		},
	},
}
var product = []Product{
	{IDproduct: 1, Name: "Pizza", size: []Size{
		{IDsize: 1, Name: "25sm", Price: 30000},
		{IDsize: 2, Name: "30sm", Price: 45000},
		{IDsize: 3, Name: "35sm", Price: 70000},
	},
	},
	{
		IDproduct: 2, Name: "Burger", size: []Size{
			{IDsize: 1, Name: "Small", Price: 15000},
			{IDsize: 2, Name: "Medium", Price: 25000},
			{IDsize: 3, Name: "Big", Price: 30000},
		},
	},
	{
		IDproduct: 3, Name: "Hot-dog", size: []Size{
			{IDsize: 1, Name: "small", Price: 20_000},
			{IDsize: 2, Name: "medium", Price: 25_000},
			{IDsize: 3, Name: "big", Price: 30_000},
		},
	},
	{
		IDproduct: 4, Name: "Cola", size: []Size{
			{IDsize: 1, Name: "0.5", Price: 5000},
			{IDsize: 2, Name: "1.O", Price: 8000},
			{IDsize: 3, Name: "1.5", Price: 10_000},
		},
	},
	{
		IDproduct: 3, Name: "Fanta", size: []Size{
			{IDsize: 1, Name: "0.5", Price: 5000},
			{IDsize: 2, Name: "1.O", Price: 8000},
			{IDsize: 3, Name: "1.5", Price: 11_000},
		},
	},
}

func main() {

	fmt.Println()
	for _, value := range clients {
		fmt.Print("name:= ", value.Name, " ")
		var sum int = 0
		for _, valueBasket := range value.Basket {
			for _, valueProduct := range product {
				for _, valueSize := range valueProduct.size {

					if valueBasket.ProductID == valueProduct.IDproduct {
						sum = valueSize.IDsize * valueSize.Price * valueBasket.Quantity

					}

				}
			}
		}
		fmt.Print("sum:=", sum, " ")

	}
	
		for _,value:=range clients{

		max := 0
		productName:=""
		for _, valueBasket:= range value.Basket{
			for _, valueProduct := range product {

				if valueProduct.IDproduct==valueBasket.ProductID{
				if valueBasket.Quantity> max {
					max = valueBasket.Quantity
					productName = valueProduct.Name
				}
			}}
			
		
			
		}
		fmt.Print("product name:= ",productName," max:",max)
		
	}
		
		
	
}

// p:=[]Product{
// 	{
// 		ID: 1,
// 		Name: "a",
// 		Price: 2000,
// 	},
// 	{
// 		ID: 2,
// 		Name: "b",
// 		Price: 3000,

// 	},
// 	{
// 		ID: 3,
// 		Name: "c",
// 		Price: 45000,
// 	},
// }

// clients := []Client{
// 	{
// 		ID:   1,
// 		Name: "A",
// 		Basket: []BasketProducts{
// 			{
// 				ProductId: 1,
// 				Quantity:  2,
// 			},
// 			{

// 				ProductId: 2,
// 				Quantity:  1,
// 			},
// 			{
// 				ProductId: 3,
// 				Quantity:  4,
// 			},
// 		},
// 	},
// 	{
// 		ID:   2,
// 		Name: "B",
// 		Basket: []BasketProducts{

// 			{
// 				ProductId: 1,
// 				Quantity:  2,
// 			},
// 			{

// 				ProductId: 2,
// 				Quantity:  1,
// 			},
// 			{
// 				ProductId: 3,
// 				Quantity:  8,
// 			},
// 		},
// 	},
// }

// for _,v:=range clients{

// 	sum:=0
// 	fmt.Print("CLIENT name: ",v.Name)
// 	for _,vBasket:=range v.Basket{
// 		for _,vp:=range p{
// 			if vBasket.ProductId==vp.ID{
// 				sum+=vp.Price*vBasket.Quantity
// 			}
// 		}
// 	}
// 	fmt.Println("Total sum: ",sum)
// }

// type Client struct {
// 	ID     int
// 	Name   string
// 	Basket []BasketProducts
// }
// type BasketProducts struct {
// 	ProductId int
// 	Quantity  int
// }
// type Product struct {
// 	ID    int
// 	Name  string
// 	Price int
// }
