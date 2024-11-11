package property

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
