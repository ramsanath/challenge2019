package data

import (
	"fmt"
	"qube/model"
	"strconv"
	"strings"
)

func GetPartners() (partners []model.Partner) {
	contents, err := ReadCSV("/data/partners.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	capacities := GetPartnerCapacity()
	partners = make([]model.Partner, 0)
	partnerIndexes := make(map[string]int)

	for i := 1; i < len(contents); i++ {
		content := contents[i]

		partnerId := strings.TrimSpace(content[4])
		theatreId := strings.TrimSpace(content[0])
		theatre := model.Theatre{ID: theatreId}

		slab := strings.Split(strings.TrimSpace(content[1]), "-")
		fromSlab, _ := strconv.Atoi(strings.TrimSpace(slab[0]))
		toSlab, _ := strconv.Atoi(strings.TrimSpace(slab[1]))
		minCost, _ := strconv.Atoi(strings.TrimSpace(content[2]))
		costPerUnit, _ := strconv.Atoi(strings.TrimSpace(content[3]))

		tariff := model.Tariff{
			MinCost:     minCost,
			CostPerUnit: costPerUnit,
			From:        fromSlab,
			To:          toSlab,
		}

		partnerIndex, present := partnerIndexes[partnerId]
		if present {
			p := &partners[partnerIndex]
			p.AddTheatreTariff(theatre, tariff)
		} else {
			p := model.Partner{
				ID:       partnerId,
				Capacity: capacities[partnerId],
			}
			p.AddTheatreTariff(theatre, tariff)
			partners = append(partners, p)
			partnerIndexes[p.ID] = len(partners) - 1
		}
	}

	return
}

func GetPartnerCapacity() (capacities map[string]int) {
	capacities = make(map[string]int, 0)

	contents, err := ReadCSV("/data/capacities.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i < len(contents); i++ {
		content := contents[i]
		partnerId := strings.TrimSpace(content[0])
		capacity, _ := strconv.Atoi(strings.TrimSpace(content[1]))
		capacities[partnerId] = capacity
	}

	return
}

func GetInput() (requests []model.DeliveryRequest) {
	contents, err := ReadCSV("/data/input.csv")
	requests = make([]model.DeliveryRequest, 0)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(contents); i++ {
		content := contents[i]

		deliveryId := strings.TrimSpace(content[0])
		units, _ := strconv.Atoi(strings.TrimSpace(content[1]))
		theatreId := strings.TrimSpace(content[2])

		dr := model.DeliveryRequest{
			ID:       deliveryId,
			Quantity: units,
			Theatre:  model.Theatre{ID: theatreId},
		}
		requests = append(requests, dr)
	}

	return
}
