package main

import (
	_ "embed"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"math/rand"
	"time"
)

const SAVE_FILE = "save.json"

//go:embed game.json
var game []byte

var CurrentGame = &Game{}

func NewChoice() *Choice {
	return &Choice{}
}

func main() {
	rand.Seed(time.Now().Unix())
	current := loadGame()
	previous := current
	color.Cyan("Bienvenue %v, vous avez %v PV, %v ATK et %v DEF", CurrentGame.Player.Name, CurrentGame.Player.HP, CurrentGame.Player.ATK, CurrentGame.Player.DEF)
	for {
		current = handleChoice(previous)
		if current == nil {
			break
		}
		previous = current

	}
	finishGame(previous)
}

func handleChoice(choice *Choice) *Choice {

	GamePrint(MESSAGE, color.Cyan, choice.Message)

	if choice.Fight != nil {
		return choice.Choice[FightOpponents(choice)]
	}
	e := []string{}

	for _, item := range choice.Choice {
		e = append(e, item.Name)
	}
	switch len(choice.Choice) {
	case 0:
		return nil

	case 1:
		return choice.Choice[0]

	default:
		for {
			prompt := promptui.Select{
				Label: choice.Prompt,
				Items: e,
			}
			i, _, err := prompt.Run()

			if handleEnd(err, choice) {
				continue
			}
			return choice.Choice[i]
		}

	}

}
