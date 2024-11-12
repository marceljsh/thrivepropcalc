package property

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

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

	if fields[4] != "STANDARD" && fields[4] != "PREMIUM" {
		return errors.New("invalid location type")
	}

	if fields[5] != "YES" && fields[5] != "NO" && fields[5] != "CORNER" {
		return errors.New("invalid corner specification")
	}

	var parking int
	if _, err := fmt.Sscan(fields[6], &parking); parking < 0 || err != nil {
		return errors.New("parking units must be non-negative")
	}

	return nil
}

func parseProperty(line string) (Property, error) {
	fields := strings.Fields(line)
	area := 0.0
	year := 0
	parking := 0

	fmt.Sscan(fields[2], &area)
	fmt.Sscan(fields[3], &year)
	fmt.Sscan(fields[6], &parking)

	facilities := strings.Split(fields[7], ",")
	uniqueFacilities := make(map[string]struct{})
	for _, facility := range facilities {
		if _, exists := uniqueFacilities[facility]; exists {
			return Property{}, errors.New("invalid or duplicate facilities")
		}
		uniqueFacilities[facility] = struct{}{}
	}

	return Property{
		Timestamp:  fields[0],
		Type:       fields[1],
		Area:       area,
		BuildYear:  year,
		Location:   fields[4],
		Corner:     fields[5],
		Parking:    parking,
		Facilities: facilities,
	}, nil
}
