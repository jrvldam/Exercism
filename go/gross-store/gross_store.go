package gross

var baseUnits = map[string]int{
	"quarter_of_a_dozen": 3,
	"half_of_a_dozen":    6,
	"dozen":              12,
	"small_gross":        120,
	"gross":              144,
	"great_gross":        1728,
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return baseUnits
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// getUnitValue returns the value of given unit in units if exists
func getUnitValue(units map[string]int, unit string) (int, bool) {
	value, ok := units[unit]
	if !ok {
		return 0, false
	}

	return value, true
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	qty, ok := getUnitValue(units, unit)
	if !ok {
		return false
	}

	_, ok = GetItem(bill, item)
	if ok {
		bill[item] += qty
	} else {
		bill[item] = qty
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	oldQty, ok := GetItem(bill, item)
	if !ok {
		return false
	}

	qty, ok := getUnitValue(units, unit)
	if !ok {
		return false
	}

	newQty := oldQty - qty
	if newQty < 0 {
		return false
	} else if newQty > 0 {
		bill[item] = newQty
		return true
	} else {
		delete(bill, item)
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
	if !ok {
		return 0, false
	}

	return qty, true
}
