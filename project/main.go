package main

import (
	"encoding/json"
	"fmt"
	"hm/project/basket"
	"hm/project/client"
	"hm/project/product"
	"hm/project/repository"
	"hm/project/store"
	"os"
)

/* HOMEWORK
Select User type:
1 - Boss
2 - Customer
3 - quit
if 1 {
	Enter password:
	if password {
		Boss login:
		1 - report
		2 - product list
		3 - back
		4 - quit
	} else {
		password is wrong
	}
} else if 2 {
	1 - start shopping
	2 - my basket
	3 - finish
	4 - back
	5 - quit
} else if 3 {
	program should quit
} else {
	not found
}
*/

const ( //cosutumer cmd
	Start = iota + 1
	Mybasket
	Finish
)

type PList struct {
	Name          string `json:"Name"`
	Price         int    `json:"Price"`
	Quantity      int    `json:"Quantity"`
	OriginalPrice int    `json:"PriginalPrice"`
}

const ( //main page
	BossCmd = iota + 1
	CustomerCmd
	UserQuitCmd
)
const ( //boss cmd
	ReportCmd = iota + 1
	ListCmd
	QuitCmd
)
const Password = "1234"

func main() {
	file, err := os.Open("file.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	pls := PList{}

	if err := json.NewDecoder(file).Decode(&pls); err != nil {

		fmt.Println("error ocured")
	}
	fmt.Println(pls)

	repos := repository.NewRepository(product.List)
	s := store.NewStore(repos)

	for true {
		var userCmd int
		fmt.Print("1.boos 2.custsomer 3.quit")
		fmt.Scan(&userCmd)
		//MAIN
		switch userCmd {
		case BossCmd:
			fmt.Println("1,report 2.productlist, 3 quite")
			var password string
			fmt.Println("enter password")
			fmt.Scan(&password)
			if password != Password {
				fmt.Println("password is wrong")
				return
			}
			fmt.Println("BOOS LOGIN")
			var bossCmd int
			fmt.Print("Enter command : ")
			fmt.Scan(&bossCmd)

			//BOSS
			switch bossCmd {
			case ReportCmd:
				fmt.Println("report selected")

			case ListCmd:
				fmt.Println("list")
			case QuitCmd:
				return
			default:
				fmt.Println("error")
			}

			//client
		case CustomerCmd:

			b := basket.NewBasket()              //basket create
			name, cash := client.GetClientInfo() //name bn cash olib client info ilindi

			user := client.NewClient(name, cash, b) //userga owa name bilan cash berildi create qb
			var userCmd int
			//////////********

			fmt.Println("1.startshopping,2,my basket 3.finish 4 back 5 quit")
			fmt.Scan(&userCmd)

			switch userCmd {
			case Start:
				s.Sell(user)
			case Mybasket:
				user.Mybasket()
			case Finish:
				return
			default:
				fmt.Println("error comand")
			}

		}

	}

}
