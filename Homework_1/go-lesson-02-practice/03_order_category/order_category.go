package order_category

import "fmt"

func OrderCategory (amount int) {
	var category string
	switch {
	case amount < 0:
		fmt.Println("Ошибка: сумма заказа не может быть отрицательной")
		return
	case amount < 10000:
		category = "small"
	case amount < 50000:
		category = "regular"
	default:
		category = "large"
	}
	fmt.Printf("Сумма заказа: %d коп.\n", amount)
	fmt.Printf("Категория: %s", category)
}
