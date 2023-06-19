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
	fmt.Println("Moving")
}

func (t *Train) Stop() {
	fmt.Println("Stop")
	t.Speed = 0
}

func (t *Train) PutPassenger() {
	t.Passengers++
}

func (t *Train) OutputPassenger() {
	t.Passengers++
}

func (t Train) String() string {
	return fmt.Sprintf("Train: Speed=%d km/h, Passengers=%d", t.Speed, t.Passengers)
}
