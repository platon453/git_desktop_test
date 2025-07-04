interface Product {
  id: number;
  name: string;
  price: number;
}

class ShoppingCart {
  private items: Product[] = [];

  addProduct(product: Product): void {
    this.items.push(product);
  }

  removeProduct(productId: number): void {
    const index = this.items.findIndex(item => item.id === productId);
    if (index !== -1) {
      this.items.splice(index, 1);
    }
  }

  getTotalPrice(): number {
    return this.items.reduce((total, product) => total + product.price, 0);
  }

  listProducts(): void {
    if (this.items.length === 0) {
      console.log("Корзина пуста.");
    } else {
      this.items.forEach(product => {
        console.log(`${product.name} - ${product.price}₽`);
      });
    }
  }
}

const product1: Product = { id: 1, name: 'Ноутбук', price: 50000 };
const product2: Product = { id: 2, name: 'Смартфон', price: 30000 };

const cart = new ShoppingCart();
cart.addProduct(product1);
cart.addProduct(product2);
cart.listProducts();

cart.removeProduct(1);
cart.listProducts();

console.log(`Общая стоимость: ${cart.getTotalPrice()}₽`);
