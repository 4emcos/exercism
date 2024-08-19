package purchase

import "fmt"

var (
	vehicles = map[string]bool{
		"car":   true,
		"bike":  false,
		"truck": true,
	}
	estimationPercent = map[string]float64{
		"3y":   0.8,
		"10y":  0.5,
		"midY": 0.7,
	}
)

const chooseVehicleMessage = "%s is clearly the better choice."

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return vehicles[kind]
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return fmt.Sprintf(chooseVehicleMessage, option1)
	}
	return fmt.Sprintf(chooseVehicleMessage, option2)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * estimationPercent["3y"]
	}

	if age >= 10 {
		return originalPrice * estimationPercent["10y"]
	}

	return originalPrice * estimationPercent["midY"]
}
