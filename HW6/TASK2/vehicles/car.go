package vehicles

import "fmt"

type Car struct {
	Speed      int
	Passengers int
}

func (c *Car) ChangeSpeed(speed int) {
	if speed > 90 {
		fmt.Println("In order to follow road rules, speed can not be topped up above 90km/h")
		return
	}
	c.Speed = speed
}

func (c *Car) Move() {
	fmt.Println("Started riding")
}

func (c *Car) Stop() {
	fmt.Println("Stopping the vehicle")
	c.Speed = 0
}

func (c *Car) PutPassenger() {
	if c.Passengers > 4 {
		fmt.Println("There is no space in vehicle anymore")
		return
	}
	c.Passengers++
}

func (c *Car) OutputPassenger() {
	if c.Passengers == 1 {
		fmt.Println("Car can not go without drive!")
		return
	}
	c.Passengers--
}

func (c *Car) String() string {
	return fmt.Sprintf("Car: Speed=%d km/h, Passengers=%d", c.Speed, c.Passengers)
}
