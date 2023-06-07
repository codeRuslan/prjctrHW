package main

import (
	"fmt"
	"strconv"
)

type Animal struct {
	name           string
	animalOwner    *Zookeper
	caughtIntoCage bool
}

type Zookeper struct {
	firstName  string
	secondName string
	cagesOwned []*Cage // I had some previous experience with Golang, therefore slice were used inside Zookeper structure
}

type Cage struct {
	cageSpeciality string
	cageOwner      *Zookeper
	caughtAnimal   []*Animal
}

func (a *Animal) GetName() string {
	return a.name
}

func (a *Animal) GetAnimalOwner() *Zookeper {
	return a.animalOwner
}

func (a *Animal) GetTrueIfAnimalCage() bool {
	return a.caughtIntoCage
}

func (a *Animal) GetDetailedSummaryOfAnimal() {
	fmt.Printf("Detailed information about %s ---> ", a.GetName())
	fmt.Println("Owner: "+a.GetAnimalOwner().firstName, ", Cage: "+strconv.FormatBool(a.GetTrueIfAnimalCage()))
}

func (c *Cage) GetAnimalIntoCage(a *Animal) {
	a.caughtIntoCage = true
	c.caughtAnimal = append(c.caughtAnimal, a)
}

func main() {

	mainZookeper := Zookeper{
		firstName:  "Alex",
		secondName: "Yurieev",
	}

	predatorsCage := Cage{
		cageSpeciality: "Predators",
		cageOwner:      &mainZookeper,
	}

	nonPredatorsCage := Cage{
		cageSpeciality: "Non Predators",
		cageOwner:      &mainZookeper,
	}

	lion := Animal{
		name:           "Lion",
		animalOwner:    &mainZookeper,
		caughtIntoCage: false,
	}

	zebra := Animal{
		name:           "Zebra",
		animalOwner:    &mainZookeper,
		caughtIntoCage: false,
	}

	snake := Animal{
		name:           "Snake",
		animalOwner:    &mainZookeper,
		caughtIntoCage: false,
	}

	tiger := Animal{
		name:           "Tiger",
		animalOwner:    &mainZookeper,
		caughtIntoCage: false,
	}

	elephant := Animal{
		name:           "Elephant",
		animalOwner:    &mainZookeper,
		caughtIntoCage: false,
	}

	mainZookeper.cagesOwned = append(mainZookeper.cagesOwned, &predatorsCage, &nonPredatorsCage)
	fmt.Println("\nDetailed Summary about animals: ")
	tiger.GetDetailedSummaryOfAnimal()
	lion.GetDetailedSummaryOfAnimal()
	zebra.GetDetailedSummaryOfAnimal()
	snake.GetDetailedSummaryOfAnimal()
	elephant.GetDetailedSummaryOfAnimal()

	fmt.Println("\nCages:")
	for i, cage := range mainZookeper.cagesOwned {
		fmt.Println(i+1, " Speciality of Cage: "+cage.cageSpeciality, ", Cage Owner's Name: "+cage.cageOwner.firstName)
	}

	fmt.Println("\nPrint out if there any Animals inside predators Cage: ")
	fmt.Println(predatorsCage.caughtAnimal)
	fmt.Println("Catching tiger and lion...")
	predatorsCage.GetAnimalIntoCage(&lion)
	predatorsCage.GetAnimalIntoCage(&tiger)
	fmt.Println("\nPrint out if there any Animals inside predators Cage once more: ")
	for i, animal := range predatorsCage.caughtAnimal {
		fmt.Println(i+1, animal.name)
		fmt.Println(i+1, animal.caughtIntoCage)
	}

}
