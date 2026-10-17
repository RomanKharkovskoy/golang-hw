package order_statistics

import "fmt"

func OrderStatistics(orders []int) {
	countOrders := len(orders)
	if countOrders == 0 {
		fmt.Println("Нет заказов для анализа")
		return
	}
	var allOrders int
	minOrder := orders[0]
	maxOrder := orders[0]
	for _, v := range orders {
		allOrders += v
		if v < minOrder {
			minOrder = v
		} 
		if v > maxOrder {
			maxOrder = v
		}
	}
	meanOrder := float64(allOrders) / float64(countOrders)
	fmt.Printf("Количество заказов: %d\n", countOrders)
	fmt.Printf("Минимальный заказ: %d коп.\n", minOrder)
	fmt.Printf("Максимальный заказ: %d коп.\n", maxOrder)
	fmt.Printf("Общая сумма: %d коп.\n", allOrders)
	fmt.Printf("Средний заказ: %.0f коп.\n", meanOrder)
}