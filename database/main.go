package main

import (
	"database/sql"
	_ "errors"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Country struct {
	ID       uuid.UUID
	Name     string
	Currency string
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=person password=12345 database=country3 sslmode=disable")
	if err != nil {
		fmt.Println("error in line 9", err.Error())
		return
	}
	defer db.Close()

	//inserting
	// c := Country{
	// 	ID:       uuid.New(),
	// 	Name:     "Russia",
	// 	Currency: "rub",
	// }
	// if _, err = db.Exec(`insert into countries values($1,$2,$3)`, c.ID, c.Name, c.Currency); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	//selecting rows
	// countries := []Country{}

	// rows,err:= db.Query(`select * from countries`)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// for rows.Next(){
	// 	country:=Country{}
	// 	if err=rows.Scan(&country.ID,&country.Name,&country.Currency);err!=nil{

	// 		fmt.Println(err.Error())
	// 	}
	// 	countries=append(countries, country)
	// }
	// fmt.Println(countries)

	//Updating data

	// id:="823ae27b-8fae-4dc8-bb0f-be11e2e967ef"
	// if _,err=db.Exec(`delete from countried where id =$1`,id);err!=nil{
	// 	fmt.Println(err.Error())
	// }


	//query deleting
	id:="823ae27b-8fae-4dc8-bb0f-be11e2e967ef"
	if _,err=db.Exec(`update countries set name=$2 where id=$1`,id,"Sydeny");err!=nil{
		fmt.Println(err.Error())
	}

	
	//selecting query by id
	// country:=Country{}
	// id:=uuid.New()
	// row:=db.QueryRow(`select id,currency,name from countries where id=$1`,id)
	// if err=row.Scan(&country.ID,&country.Currency,&country.Name);err!=nil{
	// 	if !errors.Is(err,sql.ErrNoRows){
	// 		fmt.Println(err.Error())
	// 	}
	// }
	// fmt.Println(country)

	//

}
