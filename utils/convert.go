package utils

func ConvertSpeed(speedMs float64) float64 {
	const conversionFactor = 3.6
	speedKmh := speedMs * conversionFactor
	return speedKmh
}
