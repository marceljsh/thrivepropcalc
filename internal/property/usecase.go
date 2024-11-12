package property

import (
	"errors"
	"sort"
	"time"
)

func ProcessRecords(records []string) error {
	if len(records) < 2 {
		return errors.New("input is less than 2 lines of data")
	}

	var properties []Property
	var lastTime time.Time

	for _, record := range records {
		if err := validateInput(record, &lastTime); err != nil {
			return err
		}

		property, err := parseProperty(record)
		if err != nil {
			return err
		}
		properties = append(properties, property)
	}

	totalValue, maintenance := calculateProperties(properties)

	displayResults(totalValue, maintenance, properties)

	return nil
}

func calculateProperties(properties []Property) (float64, float64) {
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
