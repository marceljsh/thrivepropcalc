package property

import (
	"errors"
	"fmt"
	"strings"
)

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
