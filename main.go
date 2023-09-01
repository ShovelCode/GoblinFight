package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Entity struct {
	HP      int
	Attacks map[string]int // Attack name and its damage
}

type Player struct {
	Entity
}

type Enemy struct {
	Entity
	Name string
}

func main() {
	player := Player{
		Entity: Entity{
			HP: 100,
			Attacks: map[string]int{
				"slash": 10,
				"stab":  20,
			},
		},
	}

	goblin := Enemy{
		Entity: Entity{
			HP: 80,
			Attacks: map[string]int{
				"punch":   5,
				"headbut": 10,
			},
		},
		Name: "Goblin",
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		if goblin.HP <= 0 {
			fmt.Println("You defeated the goblin!")
			break
		}
		if player.HP <= 0 {
			fmt.Println("You were defeated by the goblin!")
			break
		}

		fmt.Println("Player HP:", player.HP)
		fmt.Println(goblin.Name, "HP:", goblin.HP)

		// Player's turn
		fmt.Println("Choose your attack:")
		for attack := range player.Attacks {
			fmt.Println("-", attack)
		}

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		damage, ok := player.Attacks[input]
		if ok {
			fmt.Println("You used", input, "and dealt", damage, "damage!")
			goblin.HP -= damage
		} else {
			fmt.Println("Invalid attack!")
			continue
		}

		// Goblin's turn (could add randomness, but for now just use punch)
		player.HP -= goblin.Attacks["punch"]
		fmt.Println(goblin.Name, "used punch and dealt", goblin.Attacks["punch"], "damage!")
	}
}
