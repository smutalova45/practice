package product

import (
	"fmt"
	"math/rand"
)

// mahsulot struktturasi

var (
	List = ProductList{
		
	}
)


type Product struct {
	Name          string
	Price         int
	Quantity      int
	OriginalPrice int
}

// skladga productni qabul qilish uchun product create qilish kere va bu uchun bizaga Productni
// metodi kere boladi
func NewProduct(name string, price, quantity int) Product { //Product uchun ishlidi faqat
	return Product{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}
func GetProductInfo() ProductSellRequest {
	var productName string
	var productQuantity = 0
	fmt.Print("enter product name")
	fmt.Scan(&productName)
	fmt.Print("enter quantity ")
	fmt.Scan(&productQuantity)

	return ProductSellRequest{
		Name:     productName,
		Quantity: productQuantity,
	}

}

type ProductSellRequest struct {
	//odam keb request soridi requestda
	Name     string //boladi
	Quantity int
}
type ProductList []Product //alies

func (p *ProductList) AddProduct(product Product) { ///type ProductList []Product
	*p = append(*p, product)
}
func (p *ProductList) RemoveProduct(index int) {
	*p = append((*p)[:index], (*p)[index+1:]...) //type ProductList []Product
}
func GenerateProductPrice(min, max int) int {
	return rand.Intn(max-min) + min /// to generate original price of product
}
