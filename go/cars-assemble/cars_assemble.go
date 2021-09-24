package cars

const carPerHour float64 = 221

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	productionWithoutError := carPerHour * float64(speed)

	return (productionWithoutError * successRate(speed)) / 100
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	productionPerHour := int(CalculateProductionRatePerHour(speed))

	return productionPerHour / 60
}

// successRate is used to calculate the ratio of an item being created without
// error for a given speed
func successRate(speed int) float64 {
	if speed >= 1 && speed <= 4 {
		return 100
	}

	if speed >= 5 && speed <= 8 {
		return 90
	}

	if speed >= 9 && speed <= 10 {
		return 77
	}

	return 0
}
