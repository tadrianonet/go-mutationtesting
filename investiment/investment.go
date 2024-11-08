package investment

func CalculateProfit(initial, final float64) float64 {
	if final > initial {
		return final - initial
	}
	return 0
}
