package analytics

type Sale struct {
	Product          string
	Category         string
	Quantity         int
	UnitPriceKopecks int64
	Cancelled        bool
}

type Report struct {
	RevenueKopecks    int64
	UnitsSold         int
	RevenueByCategory map[string]int64
	TopProduct        string
	InvalidSales      int
}

func AnalyzeSales(sales []Sale) Report {
	report := Report{RevenueByCategory: make(map[string]int64)}
	productUnits := make(map[string]int)
	maxUnits := 0
	for _, sale := range sales {
		if sale.Cancelled {
			continue
		}
		if sale.Product == "" || sale.Category == "" || sale.Quantity <= 0 || sale.UnitPriceKopecks <= 0 {
			report.InvalidSales++
			continue
		}
		revenue := int64(sale.Quantity) * sale.UnitPriceKopecks
		report.RevenueKopecks += revenue
		report.UnitsSold += sale.Quantity
		report.RevenueByCategory[sale.Category] += revenue
		productUnits[sale.Product] += sale.Quantity
		units := productUnits[sale.Product]
		if units > maxUnits || (units == maxUnits && sale.Product < report.TopProduct) {
			maxUnits = units
			report.TopProduct = sale.Product
		}
	}
	return report
}
