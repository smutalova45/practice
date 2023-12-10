package main

import (
	"encoding/json"
	"fmt"
	"os"
	_"sort"

)

type Branch struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
type Branchproduct struct {
	BranchId  int `json:"branch_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
type Categoriess struct {
	Id   int `json:"id"`
	Name int `json:"name"`
}
type Products struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryId int    `json:"category_id"`
}
type Transaction struct {
	Id        int    `json:"id"`
	BranchId  int    `json:"branch_id"`
	ProductId int    `json:"product_id"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
	CreatedAT string `json:"created_at"`
}

type Count struct {
	Key   int
	Value int
}
type Sort struct {
	Key   int
	Value int
}

func main() {

	// branchproduct:=[]Branchproduct{}
	// file3,err:=os.Open("branchproduct.json")
	// if err!=nil{
	// 	fmt.Println("error : ",err)
	// }
	// defer file3.Close()

	// countermap := make(map[int]int)
	// productmap := make(map[int]Products)
	// branchmap := make(map[int]Branch)

	// categories:=[]Categoriess{}
	// file5,err:=os.Open("categoriess.json")
	// if err!=nil{
	// 	fmt.Println("error",err)
	// }
	// defer file5.Close()

	// if err:=json.NewDecoder(file5).Decode(&categories);err!=nil{
	// 	fmt.Println("error",err)
	// }
	products := []Products{}
	file2, err := os.Open("products.json")
	if err != nil {
		fmt.Println("error", err)
	}
	defer file2.Close()

	if err := json.NewDecoder(file2).Decode(&products); err != nil {
		fmt.Println("error", err)
		return
	}

	branches := []Branch{}
	file1, err := os.Open("branch.json")
	if err != nil {
		fmt.Println("error : ", err)
	}
	defer file1.Close()

	if err := json.NewDecoder(file1).Decode(&branches); err != nil {
		fmt.Println("error", err)
		return
	}

	transaction := []Transaction{}
	file4, err := os.Open("transaction.json")
	if err != nil {
		fmt.Println("error ", err)
	}
	defer file4.Close()

	if err := json.NewDecoder(file4).Decode(&transaction); err != nil {
		fmt.Println("error", err)
		return
	}
	// branchmap := make(map[int]string)
	// tmap := make(map[int]int)
	// for _, v := range products {
	// 	productmap[v.Id] = v //hamma malumot
	// }
	// for _, v := range transaction {
	// 	countermap[v.BranchId] += productmap[v.ProductId].Price * v.Quantity
	// }

	// for _, v := range branches {
	// 	branchmap[v.Id] = v
	// }
	// for i, v := range countermap {
	// 	fmt.Printf("%s->%d\n", branchmap[i].Name, v)
	// }
	// sort1:=[]Sort{}
	// for i,v:=range countermap{
	// 	sort1=append(sort1, Sort{i,v})
	// }

	// sort.Slice(sort1,func(i,j int )bool{
	// 	return sort1[i].Value>sort1[j].Value
	// })


	// for _, v := range branches {
	// 	branchmap[v.Id] = v.Name
	// }

	
	// for _, v := range transaction {
	// 	tmap[v.BranchId] = tmap[v.BranchId] + 1
	// }

	

	// for key, v := range tmap {
	// 	fmt.Println(branchmap[key], v,"marta takrorlangan")
	// }

	// counts:=[]Count{}
	// for i,v:=range tmap{
	// 	counts=append(counts,Count{i,v})
	// }

	// sort.Slice(counts,func(i,j int )bool{
	// 	return counts[i].Value>counts[j].Value
	// })

	// for _,v:=range counts{
	// 	fmt.Println(v.Key)
	// 	fmt.Println("Branch name:",branchmap[v.Key],"->",v.Value)
	// }

	// productmap := make(map[int]string)
	// tmap := make(map[int]int)

	// for _, v := range products {
	// 	productmap[v.Id] = v.Name
	// }
	// for _, v := range transaction {
	// 	tmap[v.ProductId] = tmap[v.ProductId] + 1
	// }
	
	// counts := []Count{}
	// for i, v := range tmap {
	// 	counts = append(counts, Count{i, v})
	// }
	
	// sort.Slice(counts,func(i,j int)bool{
	// 	return counts[i].Value>counts[j].Value
	// })
     
	// for _,v:=range counts{
	// 	fmt.Println("productname",productmap[v.Key],"value",v.Value)
	// }
}
