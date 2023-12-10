package store

import (
	"fmt"

	"hm/project/client"
	"hm/project/dealer"
	"hm/project/product"
	"hm/project/repository"
	"os"
	"text/tabwriter"
)

type Store struct {
	Repository repository.Repository
	Dealer     dealer.Dealer
	Budget     int
	Profit     int
}

// store ni metodi orqali new  store ochiladi u oziga reponi oladi
func NewStore(repository repository.Repository) Store {
	return Store{
		Repository: repository,
		Budget:     10_000,
		Profit:     0,
	}

}

func (s *Store) printStatus() {
	w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, '\t', 0)
	fmt.Fprintln(w, "NAME :\tQUANTITY :\tPRICE :\tORIGINAL PRICE :\t")

	for _, product := range s.Repository.Product {
		fmt.Fprintf(w, "%s\t%d\t%d\t%d\n", product.Name, product.Quantity, product.Price, product.OriginalPrice)
	}
	fmt.Fprintf(w, "\t\t\t\tBudget: %d\n", s.Budget)
	fmt.Fprintf(w, "\t\t\t\tProfit: %d\n", s.Profit)

	w.Flush()

}

func (s *Store) Sell(user client.Client) {
	// productSellRequest := product.GetProductInfo()
	// 		//sell requestni ornig aclientINfo func

	productSellRequest := product.GetProductInfo()
	p, exist := s.Repository.Search(productSellRequest.Name)
	if !exist {
		if user.Cash < p.Price*productSellRequest.Quantity {
			fmt.Println("you dont have enogh cash")
			return
		}
		fmt.Printf("we dont have %s this product\n WE will BRING ' %s ' in the NEXT time) ", productSellRequest.Name, productSellRequest.Name)

		product, success := s.Dealer.ProvideProduct(productSellRequest.Name, s.Budget, productSellRequest.Quantity)

		if !success {
			fmt.Println("WE WILL BUY!!!")
			return
		}
		s.Repository.Product.AddProduct(product)
		s.Budget -= product.OriginalPrice * product.Quantity
		return
	}

	if p.Quantity < productSellRequest.Quantity {

		fmt.Printf("we have %s enought, left %d ", productSellRequest.Name, p.Quantity)
		return
	}
	s.Repository.TakeProduct(productSellRequest.Name, productSellRequest.Quantity)
	user.AddProduct(product.Product{
		Name: p.Name,
		Quantity: productSellRequest.Quantity,
		Price: p.Price,
	})
	// fmt.Println("user basket", user.Basket)
	
	s.Profit += productSellRequest.Quantity * (p.Price - p.OriginalPrice)
	s.printStatus()
    

}
