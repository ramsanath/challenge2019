package service

import (
	"fmt"
	"qube/model"
)

func FindBestPartner(request model.DeliveryRequest, partners []model.Partner) {
	var bestCost *int = nil
	var bestPartner *model.Partner = nil

	for pI := range partners {
		p := partners[pI]
		cost := GetQuote(request, p)
		if cost == nil {
			continue
		}
		if bestPartner == nil || *cost < *bestCost {
			bestCost = cost
			bestPartner = &p
		}
	}

	if bestCost != nil {
		fmt.Printf("Best partner is %s\n", bestPartner.ID)
		fmt.Printf("Best cost is %d\n", *bestCost)
	} else {
		fmt.Println("Delivery not possible")
	}
}

func GetQuote(request model.DeliveryRequest, p model.Partner) (cost *int) {
	for ttI := range p.TheatreTariffs {
		tt := p.TheatreTariffs[ttI]
		if tt.Theatre != request.Theatre {
			continue
		}
		for tI := range tt.Tariffs {
			tariff := tt.Tariffs[tI]
			cost := tariff.CostFor(request.Quantity)
			if cost > 0 {
				return &cost
			}
		}
	}
	return nil
}
