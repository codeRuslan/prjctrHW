package main

import "fmt"

type Animal struct {
	name         string
	animalOwner  *Zookeper
	cageAssigned *Cage
}

type Zookeper struct {
	firstName  string
	secondName string
	cagesOwned []*Cage // I had some previous experience with Golang, therefore slice were used inside Zookeper structure
}

type Cage struct {
	cageSpeciality string
	cageOwner      *Zookeper
}

func (a *Animal) GetName() string {
	return a.name
}

func (a *Animal) GetAnimalOwner() *Zookeper {
	return a.animalOwner
}

func (a *Animal) GetAnimalCage() *Cage {
	return a.cageAssigned
}

func (a *Animal) GetDetailedSummaryOfAnimal() {
	fmt.Printf("Detailed information about %s ---> ", a.GetName())
	fmt.Println("Owner: "+a.GetAnimalOwner().firstName, ", Cage: "+a.GetAnimalCage().cageSpeciality)
}

func main() {

	MainZookeper := Zookeper{
		firstName:  "Alex",
		secondName: "Yurieev",
	}

	PredatorsCage := Cage{
		cageSpeciality: "Predators",
		cageOwner:      &MainZookeper,
	}

	NonPredatorsCage := Cage{
		cageSpeciality: "Non Predators",
		cageOwner:      &MainZookeper,
	}

	Lion := Animal{
		name:         "Lion",
		animalOwner:  &MainZookeper,
		cageAssigned: &PredatorsCage,
	}

	Zebra := Animal{
		name:         "Zebra",
		animalOwner:  &MainZookeper,
		cageAssigned: &NonPredatorsCage,
	}

	Snake := Animal{
		name:         "Snake",
		animalOwner:  &MainZookeper,
		cageAssigned: &PredatorsCage,
	}

	Tiger := Animal{
		name:         "Tiger",
		animalOwner:  &MainZookeper,
		cageAssigned: &PredatorsCage,
	}

	Elephant := Animal{
		name:         "Elephant",
		animalOwner:  &MainZookeper,
		cageAssigned: &NonPredatorsCage,
	}

	MainZookeper.cagesOwned = append(MainZookeper.cagesOwned, &PredatorsCage, &NonPredatorsCage)

	Tiger.GetDetailedSummaryOfAnimal()
	Lion.GetDetailedSummaryOfAnimal()
	Zebra.GetDetailedSummaryOfAnimal()
	Snake.GetDetailedSummaryOfAnimal()
	Elephant.GetDetailedSummaryOfAnimal()

	for i, cage := range MainZookeper.cagesOwned {
		fmt.Println(i+1, " Speciality of Cage: "+cage.cageSpeciality, ", Cage Owner's Name: "+cage.cageOwner.firstName)
	}

}
