package vehicles

import "fmt"

type Plane struct {
	Speed      int
	Passengers int
}

func (p *Plane) ChangeSpeed(speed int) {
	p.Speed = speed
}

func (p *Plane) Move() {
	fmt.Println("Moving")
}

func (p *Plane) Stop() {
	fmt.Println("Stop")
	p.Speed = 0
}

func (p *Plane) PutPassenger() {
	p.Passengers++
}

func (p *Plane) OutputPassenger() {
	p.Passengers--
}

func (p Plane) String() string {
	return fmt.Sprintf("Plane: Speed=%d km/h, Passengers=%d", p.Speed, p.Passengers)
}
