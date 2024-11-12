package property

import (
	"sort"
	"time"
)

func CalculateProperties(properties []Property) (float64, float64) {
	currentYear := time.Now().Year()

	var totalValue, totalMaintenance float64
	for i := range properties {
		properties[i].CalculateValue(currentYear)
		totalValue += properties[i].Value
		totalMaintenance += properties[i].Maintenance
	}

	sort.Slice(properties, func(i, j int) bool {
		return properties[i].Value/properties[i].Area > properties[j].Value/properties[j].Area
	})

	return totalValue, totalMaintenance
}
