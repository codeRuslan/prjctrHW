package main

import "fmt"

// 1.
//
// Створити інтерфейс «Посилка» й реалізувати його для двох класів — «Коробка» і «Конверт».
// Для кожної поштової посилки необхідно зберігати адресу отримувача й відправника.
// Додати сортувальний відділ, який залежно від типу посилки відправляє її тим або іншим шляхом.

func (bp BoxPackage) SendToReceiver() bool {
	fmt.Println("Sending to receiver...")
	return true
}

func (bp BoxPackage) GetShipmentType() string {
	return "Truck"
}

func (lp LetterPackage) SendToReceiver() bool {
	fmt.Println("Sending to receiver...")
	return true
}

func (lp LetterPackage) GetShipmentType() string {
	return "Airplane"
}

func (p Package) ChooseRouting() {
	shipmentType := p.packaging.GetShipmentType()
	fmt.Printf("You should use shipping by %s", shipmentType)
}

type PackageType interface {
	SendToReceiver() bool
	GetShipmentType() string
}

type BoxPackage struct{}

type LetterPackage struct{}

type Package struct {
	packaging PackageType
	address   string
	sender    string
}

func main() {
	boxPackage := BoxPackage{}
	letterPackage := LetterPackage{}

	packageBox := Package{
		boxPackage,
		"Street1 1",
		"Andre",
	}

	packageLetter := Package{
		letterPackage,
		"Street 2",
		"Luizable",
	}

	packageBox.packaging.SendToReceiver()
	packageLetter.packaging.SendToReceiver()

	packageBox.ChooseRouting()
	packageLetter.ChooseRouting()
}
