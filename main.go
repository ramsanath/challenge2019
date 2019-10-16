package main

import (
	"qube/data"
	"qube/model"
	"qube/service"
)

func main() {
	partners := data.GetPartners()

	service.FindBestPartner(model.DeliveryRequest{
		ID:       "D1",
		Theatre:  model.Theatre{ID: "T2"},
		Quantity: 400,
	}, partners)
}
