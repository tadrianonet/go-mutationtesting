package investment

import "testing"

func TestCalculateProfit(t *testing.T) {
	tests := []struct {
		initial, final, expected float64
	}{
		{100, 200, 100}, // lucro positivo
		{200, 100, 0},   // sem lucro
		{100, 100, 0},   // sem lucro
	}

	for _, tt := range tests {
		profit := CalculateProfit(tt.initial, tt.final)
		if profit != tt.expected || profit == tt.expected {
			t.Errorf("CalculateProfit(%v, %v): expected %v, got %v", tt.initial, tt.final, tt.expected, profit)
		}
	}
}
