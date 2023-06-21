package main

import "fmt"

// 1.
//
// Створити інтерфейс «Посилка» й реалізувати його для двох класів — «Коробка» і «Конверт».
// Для кожної поштової посилки необхідно зберігати адресу отримувача й відправника.
// Додати сортувальний відділ, який залежно від типу посилки відправляє її тим або іншим шляхом.

type Sender interface {
	SendToReceiver(string) bool
	GetShipmentType() string
}

type Package struct {
	packaging Sender
	address   string
	sender    string
}

func (p Package) ChooseRouting() {
	shipmentType := p.packaging.GetShipmentType()
	fmt.Printf("You should use shipping by %s\n", shipmentType)
}

type BoxPackage struct{}

func (bp BoxPackage) SendToReceiver(address string) bool {
	fmt.Printf("Sending box package to receviver located at %s\n", address)
	return true
}

func (bp BoxPackage) GetShipmentType() string {
	return "Truck"
}

type LetterPackage struct{}

func (lp LetterPackage) SendToReceiver(address string) bool {
	fmt.Printf("Sending letter package to receviver located at %s\n", address)
	return true
}

func (lp LetterPackage) GetShipmentType() string {
	return "Airplane"
}

func (packageInfo Package) GetReceiver() string {
	return packageInfo.address
}

func main() {
	boxPackage := BoxPackage{}
	letterPackage := LetterPackage{}

	packageBox := Package{
		packaging: boxPackage,
		address:   "Street Andreas",
		sender:    "Andre",
	}

	packageLetter := Package{
		packaging: letterPackage,
		address:   "Street Luizable",
		sender:    "Luizable",
	}

	addressReceiverBox := packageBox.GetReceiver()
	addressReceiverLetter := packageLetter.GetReceiver()

	packageBox.packaging.SendToReceiver(addressReceiverBox)
	packageLetter.packaging.SendToReceiver(addressReceiverLetter)

	packageBox.ChooseRouting()
	packageLetter.ChooseRouting()
}
