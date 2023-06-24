package vehicles

import "fmt"

type Train struct {
	Speed      int
	Passengers int
}

func (t *Train) ChangeSpeed(speed int) {
	t.Speed = speed
}

func (t *Train) Move() {
	fmt.Println("Starting movement from Train station")
}

func (t *Train) Stop() {
	fmt.Println("Stopping at the train station")
	t.Speed = 0
}

func (t *Train) PutPassenger() {
	fmt.Println("Stopping at the train station to pick up Passenger")
	t.Passengers++
}

func (t *Train) OutputPassenger() {
	fmt.Println("Stopping at the train station to leave a Passenger")
	t.Passengers++
}

func (t Train) String() string {
	return fmt.Sprintf("Train: Speed=%d km/h, Passengers=%d", t.Speed, t.Passengers)
}
