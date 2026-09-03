package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	value := float64(productionRate) * successRate/100
	return value
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	value := CalculateWorkingCarsPerHour(productionRate, successRate)
	som := int(value/60)
    return som

}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groups:= (carsCount/10)*95000
    unitCars:= (carsCount%10)*10000

    value:= uint(groups+unitCars)

    return value
}
