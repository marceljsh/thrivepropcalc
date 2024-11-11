package property

import "math"

type Property struct {
	Timestamp   string
	Type        string
	Area        float64
	BuildYear   int
	Location    string
	Corner      string
	Parking     int
	Facilities  []string
	Value       float64
	Maintenance float64
}

func (p *Property) CalculateValue(currentYear int) {
	/* base property value */
	baseValue := 0.0
	if p.Location == "PREMIUM" {
		baseValue = premiumValuePerM2 * p.Area
	} else {
		baseValue = standardValuePerM2 * p.Area
	}

	/* annual value adjustments */
	age := currentYear - p.BuildYear

	landApprFactor := math.Pow(1+appreciationRate, float64(age))

	var depreciationRate float64
	if p.Type == "RESIDENTIAL" {
		depreciationRate = residentialDepreciation
	} else {
		depreciationRate = commercialDepreciation
	}
	buildingDeprFactor := math.Pow(1-depreciationRate, float64(age))

	propertyValue := baseValue * landApprFactor * buildingDeprFactor

	/* location bonuses */
	if p.Location == "PREMIUM" {
		propertyValue *= 1.2
	}
	if p.Corner == "CORNER" || p.Corner == "YES" {
		propertyValue *= 1.15
	}

	p.Value = propertyValue

	/* monthly maintenance */
	var baseMaintenance float64
	if p.Type == "RESIDENTIAL" {
		baseMaintenance = residentialBaseFee * p.Area
	} else {
		baseMaintenance = commercialBaseFee * p.Area
	}

	p.Maintenance = baseMaintenance + (securityFee * p.Area) + (cleaningFee * p.Area)
}
