package main

import (
	"fmt"
	"time"
)

// vehicleTypes
// bodyworkIds
// nearCities
// destination
// source
// order.Freight

func main() {
	vehicleTypesChannel := make(chan string)
	bodyworkIdsChannel := make(chan string)
	nearCitiesChannel := make(chan string)
	freight := 10

	go vehicleTypes(freight, vehicleTypesChannel)
	go bodyworkIds(freight, bodyworkIdsChannel)
	go nearCities(freight, nearCitiesChannel)

	fmt.Println(<-vehicleTypesChannel)
	fmt.Println(<-bodyworkIdsChannel)
	fmt.Println(<-nearCitiesChannel)
}

func vehicleTypes(freight int, c chan string) {
	// faz um monte de coisa
	time.Sleep(time.Second * 1)
	c <- "Veículo"
}

func bodyworkIds(freight int, c chan string) {
	// faz um monte de coisa
	time.Sleep(time.Second * 3)
	c <- "Carroceria"
}

func nearCities(freight int, c chan string) {
	// faz um monte de coisa
	time.Sleep(time.Second * 1)
	c <- "10"
}
