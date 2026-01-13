//бля чет исправил щас тип лучше


package main

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price float64
}

type ShoppingCart struct {
	Items []Product
}

func (cart *ShoppingCart) AddProduct(product Product) {
	cart.Items = append(cart.Items, product)
}

func (cart *ShoppingCart) RemoveProduct(productID int) {
	for i, item := range cart.Items {
		if item.ID == productID {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			return
		}
	}
	fmt.Println("Product not found")
}

func (cart *ShoppingCart) GetTotalPrice() float64 {
	total := 0.0
	for _, product := range cart.Items {
		total += product.Price
	}
	return total
}

func (cart *ShoppingCart) ListProducts() {
	if len(cart.Items) == 0 {
		fmt.Println("Корзина пуста.")
		return
	}
	for _, product := range cart.Items {
		fmt.Printf("%s - %.2f₽\n", product.Name, product.Price)
	}
}

func main() {
	product1 := Product{ID: 1, Name: "Ноутбук", Price: 50000}
	product2 := Product{ID: 2, Name: "Смартфон", Price: 30000}

	cart := ShoppingCart{}

	cart.AddProduct(product1)
	cart.AddProduct(product2)

	fmt.Println("Товары в корзине:")
	cart.ListProducts()

	cart.RemoveProduct(1)

	fmt.Println("\nПосле удаления товара с ID 1:")
	cart.ListProducts()

	fmt.Printf("\nОбщая стоимость: %.2f₽\n", cart.GetTotalPrice())
}
