package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Company struct {
	Id     int       `json:"id"`
	Name   string    `json:"name"`
	Owner  Owner     `json:"owner"`
	Workers []Worker `json:"workers"`
}
type Owner struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type Worker struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Salary    int    `json:"salary"`
	Level     string `json:"level"`
}

func main() {

 companies := []Company{}

	file, err := os.Open("file.json")
	if err != nil {
		fmt.Println("error ", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&companies); err != nil {
		fmt.Println("error", err)
		return
	}
	//fmt.Println(companies)


	// sort.Slice(companies, func(i, j int) bool {
	// 	return len(companies[i].Workers)>len(companies[i].Workers)

	// })


	// for _,company := range companies{
	// 	fmt.Println(company.Name)
	// }


	
      slice:=[]Worker{}
		for _,v:=range companies{
		for _,v2:=range v.Workers{
             slice = append(slice, v2)
		}
			
		}
		// fmt.Println(slice)


       sort.Slice(slice,func(i,j int) bool{


	  return slice[i].Salary> slice[j].Salary
 })
     for i:=0;i<3;i++{
		fmt.Println(slice[i].FirstName)
		fmt.Print(slice[i].LastName)
	 }


	

}
