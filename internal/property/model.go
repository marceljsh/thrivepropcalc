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
	age := currentYear - p.BuildYear

	baseValue := p.Area * float64(baseValues[p.Location])

	landApprFactor := math.Pow(1+appreciationRate, float64(age))

	depreciationRate := deppreciationRates[p.Type]
	buildingDeprFactor := math.Pow(1-depreciationRate, float64(age))

	propertyValue := baseValue * landApprFactor * buildingDeprFactor

	propertyValue = propertyValue * (1 + locationBonuses[p.Location]) * (1 + cornerBonuses[p.Corner])

	p.Value = propertyValue

	baseMaintenance := baseFees[p.Type]

	p.Maintenance = float64(baseMaintenance) + (securityFee * p.Area) + (cleaningFee * p.Area)
}
