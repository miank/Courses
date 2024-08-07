package main

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	m1 := make(map[string]int)
	m1["quarter_of_a_dozen"] = 3
	m1["half_of_a_dozen"] = 6
	m1["dozen"] = 12
	m1["small_gross"] = 120
	m1["gross"] = 144
	m1["great_gross"] = 1728

	return m1
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	m := make(map[string]int)
	return m
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// Check if the unit is valid
	quantity, validUnit := units[unit]
	if !validUnit {
		return false
	}

	// Add or update the item in the bill
	bill[item] += quantity
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// Check if the item is in the bill
	currentQuantity, itemInBill := bill[item]
	if !itemInBill {
		return false
	}

	// Check if the unit is valid
	quantity, validUnit := units[unit]
	if !validUnit {
		return false
	}

	// Calculate the new quantity
	newQuantity := currentQuantity - quantity
	if newQuantity < 0 {
		return false
	} else if newQuantity == 0 {
		delete(bill, item)
	} else {
		bill[item] = newQuantity
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, found := bill[item]
	if !found {
		return 0, false
	}
	return quantity, true
}

func main() {

}
