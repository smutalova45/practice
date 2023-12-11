package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	_ "sort"
	"time"
)

type Transaction struct {
	Id        int    `json:"id"`
	BranchId  int    `json:"branch_id"`
	ProductId int    `json:"product_id"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
}
type Branch struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type Products struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryId int    `json:"category_id"`
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
	transactin := []Transaction{}
	branch := []Branch{}
	category := []Category{}
	products := []Products{}

	file1, err := os.Open("transaction.json")
	if err != nil {
		fmt.Println("error in transaction", err)
	}
	defer file1.Close()

	if err := json.NewDecoder(file1).Decode(&transactin); err != nil {
		fmt.Println(err)
	}

	file2, err := os.Open("branches.json")

	if err != nil {
		fmt.Println(err)
	}
	defer file2.Close()

	if err := json.NewDecoder(file2).Decode(&branch); err != nil {
		fmt.Println(err)
	}

	file3, err := os.Open("category.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file3.Close()

	if err := json.NewDecoder(file3).Decode(&category); err != nil {
		fmt.Println(err)
	}

	file4, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file4.Close()

	if err := json.NewDecoder(file4).Decode(&products); err != nil {
		fmt.Println(err)
	}

	//question 1
	fmt.Println("--------------1-----------------")

	categoryMap := make(map[int]string)

	productMap := make(map[int]int)
	for _, v := range category {
		categoryMap[v.Id] = v.Name
	}

	for _, tv := range transactin {

		for _, pv := range products {

			if pv.Id == tv.ProductId {
				productMap[pv.CategoryId]++
			}
		}
	}
	fmt.Println(productMap)
	counts := []Count{}
	for i, v := range productMap {
		counts = append(counts, Count{i, v})
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Value > counts[j].Value
	})
	for _, v := range counts {
		fmt.Println("name:", categoryMap[v.Key], v.Value)
	}
	fmt.Println("----------------2--------------------")

	//question2

	//.har bir branchda har bir categorydan qancha transaction bolgani

	bMap := make(map[int]string)
	catMap := make(map[int]string)
	Cmap := make(map[int]map[int]int)
	for _, b := range branch {
		bMap[b.Id] = b.Name
		Cmap[b.Id] = make(map[int]int) //idlani har bitasiga map ni tenglashtirib chiqdi
	}
	fmt.Println(Cmap)
	for _, c := range category {
		catMap[c.Id] = c.Name
	}
	for _, p := range products {
		for _, t := range transactin {
			for _, b := range branch {
				if p.Id == t.ProductId && b.Id == t.BranchId {
					Cmap[b.Id][p.CategoryId]++
				}

			}
		}
	}
	// fmt.Println(Cmap)
	// fmt.Println("**********")

	for k, v := range Cmap {
		fmt.Println(bMap[k], v) //k maksim gorkiy
		count5 := []Count{}
		for k, c := range v {
			count5 = append(count5, Count{k, c})
		}
		for _, c := range count5 {
			fmt.Printf("%s -> %d\n", catMap[c.Key], c.Value)
		}
	}

	// 3 har bir branch har bir categoryda qancha transaction bolgani

	fmt.Println("-------------------3---------------")

	namemap := map[int]string{}
	productmap1 := make(map[int]int)
	totalsumplus := map[int]int{}
	totalsumminus := map[int]int{}
	minuscount := map[int]int{}
	pluscount := map[int]int{}
	for _, v := range branch {
		namemap[v.Id] = v.Name
	}
	for _, v := range products {
		productmap1[v.Id] = v.Price

	}
	for _, t := range transactin {
		if t.Type == "plus" {
			pluscount[t.BranchId]++
			totalsumplus[t.BranchId] += productmap1[t.ProductId] * t.Quantity
		} else {
			minuscount[t.BranchId]++
			totalsumminus[t.BranchId] += productmap1[t.ProductId] * t.Quantity

		}
	}
	for key := range namemap {
		fmt.Printf("%s\n Transaction: -> plus %d,minus->%d\n Totalsum Plus->%d minus->%d\n\n", namemap[key], pluscount[key], minuscount[key], totalsumplus[key], totalsumminus[key])
	}

	fmt.Println("-------------4----------------------")

	// har bir kunda kirgan productlar sonini kameyish tartibda chiqarish

	// productCountByDay := make(map[string]int)
	// for _, t := range transactin {
	// 	createdTime, err := time.Parse("2006-01-02 15:04:05",t.CreatedAt)
	// 	if err != nil {
	// 		fmt.Println("error in parsing:::", err)
	// 	}

	// 	//  fmt.Println(createdTime)

	// 	for _, v := range transactin {
	// 		if v.Type=="plus"{
	// 		day := createdTime.Format(createdTime.Format(v.CreatedAt))
	// 		productCountByDay[day] += v.Quantity
	// 	}
	// }
	// }

	// fmt.Println("Product count by day:")
	// for day, count := range productCountByDay {
	// 	fmt.Printf("%s: %d\n", day, count)
	// }
	counterMap := make(map[string]int)

	for _, t := range transactin {

		if t.Type == "plus" {

			times, err := time.Parse("2006-01-02 15:04:05", t.CreatedAt)
			if err != nil {
				fmt.Println("Error in Parsing Time")
				return
			}

			str := times.Format("2006-01-02")

			counterMap[str] += t.Quantity

		}
	}
	fmt.Println("Product count by day:")
	for day, count := range counterMap {
		fmt.Printf("%s: %d\n", day, count)
	}

	fmt.Println("------------------------5----------------------------")

	name2 := map[int]string{}
	minuscount2 := map[int]int{}
	pluscount2 := map[int]int{}

	for _, v := range products {
		name2[v.Id] = v.Name

	}
	for _, t := range transactin {
		if t.Type == "plus" {
			pluscount2[t.ProductId] += t.Quantity
		} else {
			minuscount2[t.ProductId] += t.Quantity
		}
	}

	for key := range name2 {
		if pluscount2[key] != 0 || minuscount2[key] != 0 {
			fmt.Println(name2[key], "-> ", "plus: ", pluscount2[key], "minus:", minuscount2[key])
		}
	}

	fmt.Println("-------------6--------------")
	/*
		    Uacademy ->18000
			C1->58000
			maksim goktiy->20000
			ucademy ->40000
	*/

	nameb := map[int]string{}
	result := map[int]int{}
	pluscount3 := map[int]int{}
	minuscount3 := map[int]int{}

	totalsum := map[int]int{}

	for _, v := range branch {
		nameb[v.Id] = v.Name

	}
	for _, v := range products {
		result[v.Id] = v.Price
	}

	for _, v := range transactin {
		if v.Type == "plus" {
			pluscount3[v.BranchId] += result[v.ProductId] * v.Quantity
		} else {
			minuscount3[v.BranchId] += result[v.ProductId] * v.Quantity
		}

		totalsum[v.BranchId] = pluscount3[v.BranchId] - minuscount3[v.BranchId]

	}
	fmt.Println(pluscount3)
	fmt.Println("---------")
	fmt.Println(minuscount3)
	fmt.Println("********")
	fmt.Println(totalsum)

	for key:=range nameb{
		fmt.Println(nameb[key],totalsum[key])
	}
}
