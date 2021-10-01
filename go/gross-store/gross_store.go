package gross

// Units store the Gross Store unit measurement
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill create a new bill
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem add item to customer bill
func AddItem(bill, units map[string]int, item, unit string) bool {
	itemUnit, exists := itemMapExists(units, unit)

	if !exists {
		return false
	}

	bill[item] = itemUnit

	return true
}

// RemoveItem remove item from customer bill
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	quantityItem, itemToRemoveExists := itemMapExists(bill, item)
	itemUnit, unitExists := itemMapExists(units, unit)

	if !itemToRemoveExists || !unitExists {
		return false
	}

	newQuantity := quantityItem - itemUnit

	if newQuantity < 0 {
		return false
	}

	bill[item] = newQuantity

	if newQuantity == 0 {
		delete(bill, item)
	}

	return true
}

// GetItem return the quantity of item that the customer has in his/her bill
func GetItem(bill map[string]int, item string) (int, bool) {
	return itemMapExists(bill, item)
}

func itemMapExists(mapItems map[string]int, item string) (int, bool) {
	itemInMap := mapItems[item]
	exists := true

	if itemInMap == 0 {
		exists = false
	}

	return itemInMap, exists
}
