//Package to analyze production of cars in a car factory.
package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(productionRate) * (successRate  / float64(100.0))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    var workingCarPerHour float64 = CalculateWorkingCarsPerHour(productionRate, successRate)
    return int(workingCarPerHour / 60.0) 
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    const TenCarsCost int = 95000
    const CarCost int = 10000
    return uint((carsCount / 10 * TenCarsCost) + (carsCount % 10) * CarCost)
}
