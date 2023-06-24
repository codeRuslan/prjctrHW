package main

import "fmt"

type Sender interface {
	SendToReceiver(string) bool
	GetShipmentType() string
	GetUniqueField() string
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

type BoxPackage struct {
	boxSize string // Unique field for BoxPackage
}

func (bp BoxPackage) SendToReceiver(address string) bool {
	fmt.Printf("Sending box package to receiver located at %s\n", address)
	return true
}

func (bp BoxPackage) GetShipmentType() string {
	return "Truck"
}

func (bp BoxPackage) GetUniqueField() string {
	return bp.boxSize
}

type LetterPackage struct {
	paperType string // Unique field for LetterPackage
}

func (lp LetterPackage) SendToReceiver(address string) bool {
	fmt.Printf("Sending letter package to receiver located at %s\n", address)
	return true
}

func (lp LetterPackage) GetShipmentType() string {
	return "Airplane"
}

func (lp LetterPackage) GetUniqueField() string {
	return lp.paperType
}

func (packageInfo Package) GetReceiver() string {
	return packageInfo.address
}

func main() {
	boxPackage := BoxPackage{boxSize: "Large"}
	letterPackage := LetterPackage{paperType: "A4"}

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

	uniqueFieldBox := packageBox.packaging.GetUniqueField()
	uniqueFieldLetter := packageLetter.packaging.GetUniqueField()

	fmt.Printf("Unique field for BoxPackage: %s\n", uniqueFieldBox)
	fmt.Printf("Unique field for LetterPackage: %s\n", uniqueFieldLetter)
}
