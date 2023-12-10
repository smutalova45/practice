package repository
import (
	"hm/project/product"
)

type Repository struct {
	Product product.ProductList
}

// skladni new metodi orqali create qilamz
func NewRepository(product product.ProductList) Repository {
	return Repository{
		Product: product,
	}
}
func (r *Repository) Search(productName string) (product.Product, bool) {
	for _, product := range r.Product {

		if product.Name == productName {
			return product, true
		}
	}
	return product.Product{}, false
}
func (r *Repository) TakeProduct(productName string, quantity int) {
	for index, product := range r.Product {
		if product.Name == productName {
			r.Product[index].Quantity -= quantity
			if r.Product[index].Quantity == 0 {
				r.Product.RemoveProduct(index)
			}
			return
		}
	}
}
