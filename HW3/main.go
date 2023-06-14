package main

import (
	"fmt"
	"os"
)

type Inventory struct {
	items        []string
	limitOfItems int
}

type Character struct {
	characterName string
	inventory     *Inventory
	health        int
	mana          int
}

type Adventure struct {
	mainPart         string
	optionsPart      [2]string
	deltaHealth      int
	deltaMana        int
	end              bool
	nextAdventureOne *Adventure
	nextAdventureTwo *Adventure
}

func (i *Inventory) PutItemToInventory(inputItem string) {
	i.items = append(i.items, inputItem)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!")
	fmt.Printf("You got a new item to inventory named %s", inputItem)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!")
}

func main() {
	var userChoice int

	itemsInventory := make([]string, 0)

	playerInventory := Inventory{
		itemsInventory,
		10,
	}

	playerCharacter := Character{
		"Alexis",
		&playerInventory,
		50,
		50,
	}

	Start := Adventure{
		"You find yourself standing at the entrance of an ancient, mysterious cave. Legends speak of hidden treasures and untold secrets lurking within its depths. As you peer into the darkness, you can sense an eerie energy emanating from the cavern, both enticing and foreboding",
		[2]string{"You take a deep breath, steeling your nerves, and step forward into the yawning abyss. The air grows colder, and the faint sound of dripping water echoes through the cavernous walls. Your heart races with a mix of excitement and trepidation as you venture deeper into the unknown. With each step, the darkness envelops you, leaving only the faint glow of your torch to guide your way. The path ahead splits into two narrow tunnels",
			"You hesitate for a moment, a nagging sense of caution tugging at your instincts. Despite your curiosity, you decide that discretion is the better part of valor. Slowly, you back away from the entrance of the cave, your mind filled with what-ifs and unanswered questions. Perhaps it's best to seek a different adventure or gather more information before delving into the unknown. As you turn your back on the cave, you feel a mix of relief and a tinge of regret for the path unexplored, however sudenly you got hit by a  small rock into your head and see a monster in front of ancient cave"},
		0,
		0,
		false,
		nil,
		nil,
	}

	StartA := Adventure{
		"You take a deep breath, steeling your nerves, and step forward into the yawning abyss. The air grows colder, and the faint sound of dripping water echoes through the cavernous walls. Your heart races with a mix of excitement and trepidation as you venture deeper into the unknown. With each step, the darkness envelops you, leaving only the faint glow of your torch to guide your way. The path ahead splits into two narrow tunnels",
		[2]string{"You decide to turn right, following the tunnel that veers off to the side. The narrow passage becomes even more claustrophobic as you proceed, and the air grows thick with an unsettling stillness. After what feels like an eternity of winding through the labyrinthine tunnels, you come across a faint glimmer of light ahead. Intrigued, you quicken your pace, eager to uncover the source of the illumination.It happens to be a strange monster glowing green.",
			"You opt to turn left, intrigued by the mysteries that lie in that direction. The tunnel gradually widens, and you find yourself in a vast, dimly lit chamber. The walls are adorned with intricate carvings, hinting at a forgotten civilization. A faint sound of rushing water catches your attention, leading your gaze to a distant waterfall cascading from a crevice above. The sight is mesmerizing, beckoning you to explore further."},
		10,
		15,
		false,
		nil,
		nil,
	}
	StartAA := Adventure{
		"You decide to turn right, following the tunnel that veers off to the side. The narrow passage becomes even more claustrophobic as you proceed, and the air grows thick with an unsettling stillness. After what feels like an eternity of winding through the labyrinthine tunnels, you come across a faint glimmer of light ahead. Intrigued, you quicken your pace, eager to uncover the source of the illumination.It happens to be a strange monster glowing green.",
		[2]string{"You decide to fight him and fight for life and death begins now!", "You start to run around and screaming spells from a game!"},
		0,
		0,
		true,
		nil,
		nil,
	}

	StartAB := Adventure{
		"You opt to turn left, intrigued by the mysteries that lie in that direction. The tunnel gradually widens, and you find yourself in a vast, dimly lit chamber. The walls are adorned with intricate carvings, hinting at a forgotten civilization. A faint sound of rushing water catches your attention, leading your gaze to a distant waterfall cascading from a crevice above. The sight is mesmerizing, beckoning you to explore further.",
		[2]string{"You noticed the strange chest near it and decide to check it",
			"Then you suddenly notice women sitting in front off the wall"},
		0,
		0,
		true,
		nil,
		nil,
	}

	StartB := Adventure{
		"You hesitate for a moment, a nagging sense of caution tugging at your instincts. Despite your curiosity, you decide that discretion is the better part of valor. Slowly, you back away from the entrance of the cave, your mind filled with what-ifs and unanswered questions. Perhaps it's best to seek a different adventure or gather more information before delving into the unknown. As you turn your back on the cave, you feel a mix of relief and a tinge of regret for the path unexplored, however sudenly you got hit by a  small rock into your head and see a monster in front of ancient cave",
		[2]string{"You start running away from it and screaming",
			"You take a small pocket knife out of your pocket and prepare to duel with it"},
		-15,
		5,
		false,
		nil,
		nil,
	}

	StartBA := Adventure{
		"You start running away from it and screaming",
		[2]string{"You got hit by a strange creature with a fireball",
			"You got hit by a strange creature with a fireball"},
		-100,
		-100,
		true,
		nil,
		nil,
	}

	StartBB := Adventure{
		"You take a small pocket knife out of your pocket and prepare to duel with it",
		[2]string{"Then you take a empty water bottle near you, as additional item to fight",
			"Then you take a branch near you, as additional item to fight"},
		5,
		2,
		true,
		nil,
		nil,
	}

	Start.nextAdventureOne = &StartA
	Start.nextAdventureTwo = &StartB

	StartA.nextAdventureOne = &StartAA
	StartA.nextAdventureTwo = &StartAB

	StartB.nextAdventureOne = &StartBA
	StartB.nextAdventureTwo = &StartBB

	currentAdventure := &Start

	for !currentAdventure.end {
		currentAdventure.printAdventureText(&playerCharacter)
		fmt.Scan(&userChoice)

		if userChoice == 1 {
			currentAdventure = currentAdventure.nextAdventureOne
		} else if userChoice == 2 {
			currentAdventure = currentAdventure.nextAdventureTwo
		}
	}
	fmt.Println("The End")
	fmt.Println("Good Bye!")
	os.Exit(1)

}

func (adv *Adventure) printAdventureText(c *Character) {

	c.editHealthAndMana(adv.deltaMana, adv.deltaHealth)

	fmt.Println("----------------------------------------")
	fmt.Println(adv.mainPart)
	fmt.Println("***---***\nType 1 or 2 to choose next option")
	fmt.Println("\nOption 1:", adv.optionsPart[0])
	fmt.Println("\nOption 2", adv.optionsPart[1])

	fmt.Println("****************************************")
	fmt.Println("\nChracter health: ", c.health)
	fmt.Println("Charachter mana: ", c.mana)
	fmt.Println("----------------------------------------")
}

func (c *Character) editHealthAndMana(deltaMana int, deltaHealth int) {
	c.mana += deltaMana
	c.health += deltaHealth

	if c.mana >= 100 { // In order to keep the limit of mana at 100
		c.mana = 100
	}

	if c.health >= 100 { // In order to keep the limit of health at 100
		c.health = 100
	}
}
