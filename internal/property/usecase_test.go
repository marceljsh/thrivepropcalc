package property

import "testing"

func TestCalculateValue(t *testing.T) {
	tests := []struct {
		name                string
		property            Property
		expectedValue       float64
		expectedMaintenance float64
	}{
		{
			name: "Residential Premium Corner",
			property: Property{
				Type:      "RESIDENTIAL",
				Area:      150.0,
				BuildYear: 2020,
				Location:  "PREMIUM",
				Corner:    "YES",
			},
			expectedValue:       3410650856.62,
			expectedMaintenance: 272500.00,
		},
		{
			name: "Commercial Standard",
			property: Property{
				Type:      "COMMERCIAL",
				Area:      200.0,
				BuildYear: 2018,
				Location:  "STANDARD",
				Corner:    "NO",
			},
			expectedValue:       2.1643608527022924e+09,
			expectedMaintenance: 363500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.property.CalculateValue(2024)
			if tt.property.Value != tt.expectedValue {
				t.Errorf("expected value: %.2f, got: %.2f", tt.expectedValue, tt.property.Value)
			}
			if tt.property.Maintenance != tt.expectedMaintenance {
				t.Errorf("expected maintenance: %.2f, got: %.2f", tt.expectedMaintenance, tt.property.Maintenance)
			}
		})
	}
}
