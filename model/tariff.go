package model

// Tariff - Defines the cost for a delivery
type Tariff struct {
	From        int
	To          int
	CostPerUnit int
	MinCost     int
}

func (t Tariff) CostFor(quantity int) (cost int) {
	cost = -1
	if quantity >= t.From && quantity <= t.To {
		cost = t.CostPerUnit * quantity
		if cost < t.MinCost {
			cost = t.MinCost
		}
	}
	return
}