package vehicles

import "fmt"

type Car struct {
	Speed      int
	Passengers int
}

func (c *Car) ChangeSpeed(speed int) {
	c.Speed = speed
}

func (c *Car) Move() {
	fmt.Println("Moving")
}

func (c *Car) Stop() {
	fmt.Println("Stop")
	c.Speed = 0
}

func (c *Car) PutPassenger() {
	c.Passengers++
}

func (c *Car) OutputPassenger() {
	c.Passengers--
}

func (c *Car) String() string {
	return fmt.Sprintf("Car: Speed=%d km/h, Passengers=%d", c.Speed, c.Passengers)
}
