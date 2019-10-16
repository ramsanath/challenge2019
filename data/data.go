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

	partners = make([]model.Partner, 0)
	partnerIndexes := make(map[string]int)

	for i := 1; i < len(contents); i++ {
		content := contents[i]

		partnerId := strings.TrimSpace(content[4])
		theatreId := strings.TrimSpace(content[0])
		theatre := model.Theatre{ID: theatreId}

		capacity := strings.Split(strings.TrimSpace(content[1]), "-")
		minCap, _ := strconv.Atoi(strings.TrimSpace(capacity[0]))
		maxCap, _ := strconv.Atoi(strings.TrimSpace(capacity[1]))
		minCost, _ := strconv.Atoi(strings.TrimSpace(content[2]))
		costPerUnit, _ := strconv.Atoi(strings.TrimSpace(content[3]))

		tariff := model.Tariff{
			MinCost:     minCost,
			CostPerUnit: costPerUnit,
			From:        minCap,
			To:          maxCap,
		}

		partnerIndex, present := partnerIndexes[partnerId]
		if present {
			p := &partners[partnerIndex]
			p.AddTheatreTariff(theatre, tariff)
		} else {
			p := model.Partner{
				ID: partnerId,
			}
			p.AddTheatreTariff(theatre, tariff)
			partners = append(partners, p)
			partnerIndexes[p.ID] = len(partners) - 1
		}
	}

	return
}
