package util

import "math"

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func IsRoughly(value, isRoughly float64, epsilon float64) bool {
	return isRoughly >= (value-epsilon) && isRoughly <= (value+epsilon)
}
