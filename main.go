package main

import "fmt"

type Sale struct {
	ID       int
	Product  string
	Category string
	Price    float64
	Quantity int
}

func CalculateRevenue(sales []Sale) float64 {
	var revenue float64
	for _, sale := range sales {
		revenue += sale.Price * float64(sale.Quantity)
	}
	return revenue
}

func GetTopProduct(sales []Sale) string {
	quantities := make(map[string]int)
	for _, sale := range sales {
		quantities[sale.Product] += sale.Quantity
	}
	var topProduct string
	var maxQuantity int
	for product, quantity := range quantities {
		if quantity > maxQuantity {
			topProduct = product
			maxQuantity = quantity
		}
	}
	return topProduct
}

func RevenueByCategory(sales []Sale) map[string]float64 {
	revenue := make(map[string]float64)
	for _, sale := range sales {
		revenue[sale.Category] += sale.Price * float64(sale.Quantity)
	}
	return revenue
}

func ApplyDiscount(sale *Sale, discount float64) {
	sale.Price *= 1 - discount/100
}

func formatMoney(amount float64) string {
	digits := fmt.Sprintf("%.0f", amount)
	var result string
	for i, digit := range digits {
		if i > 0 && (len(digits)-i)%3 == 0 {
			result += " "
		}
		result += string(digit)
	}
	return result
}

func main() {
	sales := []Sale{
		{ID: 1, Product: "Ноутбук", Category: "Электроника", Price: 350000, Quantity: 2},
		{ID: 2, Product: "Мышь", Category: "Аксессуары", Price: 5000, Quantity: 10},
		{ID: 3, Product: "Монитор", Category: "Электроника", Price: 85000, Quantity: 3},
		{ID: 4, Product: "Клавиатура", Category: "Аксессуары", Price: 12000, Quantity: 5},
		{ID: 5, Product: "Наушники", Category: "Аксессуары", Price: 25000, Quantity: 4},
		{ID: 6, Product: "Ноутбук", Category: "Электроника", Price: 350000, Quantity: 1},
	}

	fmt.Println("============ SALES ANALYZER ============")
	fmt.Println("Продажи:")
	var totalQuantity int
	for _, sale := range sales {
		fmt.Printf("%d. %s | Категория: %s | Цена за единицу: %s тг | Количество: %d | Сумма: %s тг\n",
			sale.ID, sale.Product, sale.Category, formatMoney(sale.Price), sale.Quantity,
			formatMoney(sale.Price*float64(sale.Quantity)))
		totalQuantity += sale.Quantity
	}

	fmt.Printf("Количество записей: %d\n", len(sales))
	fmt.Printf("Продано единиц товара: %d\n", totalQuantity)
	revenue := CalculateRevenue(sales)
	fmt.Printf("Общая выручка: %s тг\n", formatMoney(revenue))
	fmt.Printf("Средняя сумма продажи: %s тг\n", formatMoney(revenue/float64(len(sales))))

	topProduct := GetTopProduct(sales)
	var topQuantity int
	for _, sale := range sales {
		if sale.Product == topProduct {
			topQuantity += sale.Quantity
		}
	}
	fmt.Println("Самый продаваемый товар:")
	fmt.Printf("%s — %d шт.\n", topProduct, topQuantity)

	categoryRevenue := RevenueByCategory(sales)
	fmt.Println("Выручка по категориям:")
	for _, category := range []string{"Электроника", "Аксессуары"} {
		fmt.Printf("%s: %s тг\n", category, formatMoney(categoryRevenue[category]))
	}

	discounts := [3]float64{5, 10, 15}
	oldPrice := sales[0].Price
	ApplyDiscount(&sales[0], discounts[1])
	revenueAfterDiscount := CalculateRevenue(sales)
	fmt.Printf("Применение скидки %.0f%%:\n", discounts[1])
	fmt.Printf("Товар: %s\n", sales[0].Product)
	fmt.Printf("Старая цена: %s тг\n", formatMoney(oldPrice))
	fmt.Printf("Новая цена: %s тг\n", formatMoney(sales[0].Price))
	fmt.Printf("Выручка после скидки: %s тг\n", formatMoney(revenueAfterDiscount))
}
