package main

import (
	"fmt"
	"qube/data"
	"qube/service"
)

func TestProblem1() {
	requests := data.GetInput()
	partners := data.GetPartners()

	for i := range requests {
		bestDelivery := service.FindBestPartner(requests[i], partners)
		fmt.Println(bestDelivery.ToCSV())
	}
}
