package property

const (
	appreciationRate = .05

	securityFee = 1_000
	cleaningFee = 800

	timestampLayout = "2006-01-02T15:04:05"
)

var (
	baseValues = map[string]int{
		"STANDARD": 10_000_000,
		"PREMIUM":  15_000_000,
	}

	deppreciationRates = map[string]float64{
		"RESIDENTIAL": .025,
		"COMMERCIAL":  .035,
	}

	baseFees = map[string]int{
		"RESIDENTIAL": 2_500,
		"COMMERCIAL":  3_500,
	}

	locationBonuses = map[string]float64{
		"PREMIUM":  .2,
		"STANDARD": .0,
	}

	cornerBonuses = map[string]float64{
		"CORNER": .15,
		"YES":    .15,
		"NO":     .0,
	}
)
