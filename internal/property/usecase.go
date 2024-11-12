package property

import (
	"errors"
	"fmt"
	"sort"
	"strings"
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

func validateInput(line string, lastTime *time.Time) error {
	fields := strings.Fields(line)
	if len(fields) != 8 {
		return errors.New("input does not follow the correct format")
	}

	currentTime, err := time.Parse(timestampLayout, fields[0])
	if err != nil {
		return errors.New("invalid timestamp format")
	}

	if !lastTime.IsZero() && currentTime.Sub(*lastTime) > 5*time.Minute {
		return errors.New("gap between records exceeds 5 minutes")
	}
	*lastTime = currentTime

	if fields[1] != "RESIDENTIAL" && fields[1] != "COMMERCIAL" {
		return errors.New("invalid property type")
	}

	var area float64
	if _, err := fmt.Sscan(fields[2], &area); area <= 0 || err != nil {
		return errors.New("area less than or equal to zero")
	}

	var year int
	if _, err := fmt.Sscan(fields[3], &year); year < 1900 || year > time.Now().Year() || err != nil {
		return errors.New("invalid build year")
	}

	if _, ok := baseValues[fields[4]]; !ok {
		return errors.New("invalid location type")
	}

	if _, ok := cornerBonuses[fields[5]]; !ok {
		return errors.New("invalid corner specification")
	}

	var parking int
	if _, err := fmt.Sscan(fields[6], &parking); parking < 0 || parking > 99 || err != nil {
		return errors.New("parking units must be non-negative")
	}

	return nil
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
