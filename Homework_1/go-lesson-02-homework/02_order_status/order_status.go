package order_status

// CountStatuses возвращает количество new, paid, shipped, cancelled,
// unknown и active (new + paid + shipped) в указанном порядке.
// Любой статус, не равный одному из четырёх известных, считается unknown.
func CountStatuses(statuses []string) (newCount, paidCount, shippedCount, cancelledCount, unknownCount, activeCount int) {
	for _, status := range statuses {
		switch status {
		case "new":
			newCount++
		case "paid":
			paidCount++
		case "shipped":
			shippedCount++
		case "cancelled":
			cancelledCount++
		default:
			unknownCount++
		}
	}
	activeCount = newCount + paidCount + shippedCount
	return
}
