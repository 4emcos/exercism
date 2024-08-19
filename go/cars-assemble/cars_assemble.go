package cars

// determine percentage
func determinePercentage(percent float64) float64 {
	return percent / 100
}

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * determinePercentage(successRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var productionPerMinute = float64(productionRate / 60)
	return int(productionPerMinute * determinePercentage(successRate))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var firstTerm = carsCount / 10
	var secondTerm = carsCount % 10

	return uint((firstTerm * 95000) + (secondTerm * 10000))

}
