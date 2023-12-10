package client

import (
	"fmt"
	"hm/project/basket"
	"hm/project/product"
	"os"
	"text/tabwriter"
)

type Client struct {
	Name   string
	Basket basket.Basket
	Cash   int
}

func NewClient(name string, cash int, basket basket.Basket) Client {
	return Client{
		Name:   name,
		Cash:   0,
		Basket: basket,
	}
}
func GetClientInfo() (name string, cash int) {
	fmt.Print("your name ")
	fmt.Scan(&name)
	fmt.Print("your cash")
	fmt.Scan(&cash)
	return
}
func (c *Client) AddProduct(p product.Product) {
	c.Basket.ProductList.AddProduct(p) //productni add producti orqali qowib qoydik basketga
	for _, p := range c.Basket.ProductList {
		c.Basket.Total += p.Price
	}
	
	
}

func (c *Client) Mybasket() {
	w:=tabwriter.NewWriter(os.Stdout,1,8,1,'\t',0)
	fmt.Fprintln(w,"PRODUCT_NAME: \tQUANTITY : \tTOTAL_SUM")
	for _,v:=range c.Basket.ProductList{
		fmt.Fprintf(w,"%s\t%d\t%d\n",v.Name,v.Quantity,v.Price)
	}
	w.Flush()
	
}
