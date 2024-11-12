package property

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/marceljsh/thrivepropcalc/pkg/formatter"
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

func displayResults(totalValue float64, maintenance float64, properties []Property) {
	fmt.Printf("Property Value: Rp %s\n", formatter.FormatInteger(int(totalValue)))
	fmt.Printf("Monthly Maintenance: Rp %s\n", formatter.FormatInteger(int(maintenance)))

	for _, prop := range properties {
		valPerM2 := int(prop.Value / prop.Area)
		fmt.Printf("%s %s %.1f %d %s %s %s/m²\n",
			prop.Timestamp, prop.Type, prop.Area, prop.BuildYear, prop.Location,
			prop.Corner, formatter.FormatInteger(valPerM2))
	}
}
