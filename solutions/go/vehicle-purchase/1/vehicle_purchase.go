package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {	
    return kind == "car" || kind == "truck" 
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	ChoosedVehicle := ""
    if option1 < option2 {
        ChoosedVehicle = option1
    }else{
        ChoosedVehicle = option2
        }
    return ChoosedVehicle + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	price := originalPrice
    if age < 3 {
        price = originalPrice*0.8
    }else if age >= 10{
        price = originalPrice*0.5
    }else{
        price = originalPrice*0.7
    }
    return price
}
