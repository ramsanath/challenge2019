package model

import "fmt"

type BestDelivery struct {
	Delivery string
	Possible bool
	Partner  string
	Cost     int
}

func (bd BestDelivery) ToCSV() string {
	return fmt.Sprintf("%s, %t, %s, %d", bd.Delivery, bd.Possible, bd.Partner, bd.Cost)
}
