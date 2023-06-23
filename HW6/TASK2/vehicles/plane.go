package vehicles

import "fmt"

type Plane struct {
	Speed      int
	Passengers int
}

func (p *Plane) ChangeSpeed(speed int) {
	if speed < 400 {
		fmt.Println("Plane cannot sustain speed lower than 400 km/h in order to continue a flight")
		return
	}
	p.Speed = speed
}

func (p *Plane) Move() {
	fmt.Println("Started Flying")
}

func (p *Plane) Stop() {
	fmt.Println("Landing in a Airport")
	p.Speed = 0
}

func (p *Plane) PutPassenger() {
	fmt.Println("Airplane cannot pick up anymore passengers while in flight")
	return
}

func (p *Plane) OutputPassenger() {
	fmt.Println("Airplane cannot pick up anymore passengers while in flight")
	return
}

func (p Plane) String() string {
	return fmt.Sprintf("Plane: Speed=%d km/h, Passengers=%d", p.Speed, p.Passengers)
}
