package main

import "fmt"

// 1.
//
// Створити інтерфейс «Посилка» й реалізувати його для двох класів — «Коробка» і «Конверт».
// Для кожної поштової посилки необхідно зберігати адресу отримувача й відправника.
// Додати сортувальний відділ, який залежно від типу посилки відправляє її тим або іншим шляхом.

type Sender interface {
	SendToReceiver() bool
	GetShipmentType() string
}

type Package struct {
	packaging Sender
	address   string
	sender    string
}

func (p Package) ChooseRouting() {
	shipmentType := p.packaging.GetShipmentType()
	fmt.Printf("You should use shipping by %s", shipmentType)
}

type BoxPackage struct{}

func (bp BoxPackage) SendToReceiver() bool {
	fmt.Println("Sending box package to receiver...")
	return true
}

func (bp BoxPackage) GetShipmentType() string {
	return "Truck"
}

type LetterPackage struct{}

func (lp LetterPackage) SendToReceiver() bool {
	fmt.Println("Sending letter package to receiver...")
	return true
}

func (lp LetterPackage) GetShipmentType() string {
	return "Airplane"
}

func main() {
	boxPackage := BoxPackage{}
	letterPackage := LetterPackage{}

	packageBox := Package{
		packaging: boxPackage,
		address:   "Street1 1",
		sender:    "Andre",
	}

	packageLetter := Package{
		packaging: letterPackage,
		address:   "Street 2",
		sender:    "Luizable",
	}

	packageBox.packaging.SendToReceiver()
	packageLetter.packaging.SendToReceiver()

	packageBox.ChooseRouting()
	packageLetter.ChooseRouting()
}
