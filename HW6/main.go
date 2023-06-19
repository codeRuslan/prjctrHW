package main

import (
	"fmt"
	"testPlayground/vehicles"
)

type Route struct {
	vehiclesList []*Vehicle
}

func (r *Route) AddVehicle(vehicleToAdd Vehicle) {
	r.vehiclesList = append(r.vehiclesList, &vehicleToAdd)
}

func (r *Route) ShowVehicles() {
	fmt.Println("Vehicles on the road:")
	for _, vehicle := range r.vehiclesList {
		// fmt.Printf("- %T\n", *vehicle)
		fmt.Println(*vehicle)
	}
}

type Vehicle interface {
	Move()
	Stop()
	ChangeSpeed(speed int)
}

func main() {

	mainRoute := Route{}
	car := vehicles.Car{
		60,
		2,
	}
	train := vehicles.Train{
		60,
		2,
	}
	plane := vehicles.Plane{
		60,
		2,
	}

	car.ChangeSpeed(60)
	train.ChangeSpeed(60)
	plane.ChangeSpeed(60)

	car.PutPassenger()
	train.PutPassenger()
	plane.PutPassenger()

	fmt.Println(car)
	fmt.Println(train)
	fmt.Println(plane)

	fmt.Println("Initial vehicles on the Route:")
	mainRoute.ShowVehicles()

	mainRoute.AddVehicle(&car)
	mainRoute.AddVehicle(&train)
	mainRoute.AddVehicle(&plane)

	fmt.Println("Final vehicles that were used by traveller:")
	mainRoute.ShowVehicles()

}
