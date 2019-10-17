package service

import (
	"qube/model"
)

func FindBestPartner(request model.DeliveryRequest, partners []model.Partner) (result model.BestDelivery) {
	result = model.BestDelivery{
		Delivery: request.ID,
		Possible: false,
		Partner:  "",
		Cost:     -1,
	}

	for pI := range partners {
		p := partners[pI]
		cost := GetQuote(request, p)
		if cost == nil {
			continue
		}
		if !result.Possible || *cost < result.Cost {
			result = model.BestDelivery{
				Delivery: request.ID,
				Possible: true,
				Partner:  p.ID,
				Cost:     *cost,
			}
		}
	}

	return
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
