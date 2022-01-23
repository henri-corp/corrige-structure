package main

import (
	"github.com/enescakir/emoji"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	rand "math/rand"
)

func FightOpponents(choice *Choice) int {
	var err error

	fighting := choice.Fight
	GamePrint(MESSAGE, color.Red, fighting.Message)

	for {

		opponentNames := []string{}
		for _, opponent := range fighting.Opponents {
			opponentNames = append(opponentNames, opponent.Name)
		}
		selected := 0
		if len(opponentNames) > 1 {
			prompt := promptui.Select{
				Label: "choisissez",
				Items: opponentNames,
			}

			selected, _, err = prompt.Run()
			if handleEnd(err, choice) {
				continue
			}
		}
		if oppose(CurrentGame.Player, fighting.Opponents[selected], color.Green) {
			fighting.Opponents = removeOpponent(fighting.Opponents, selected)
			if len(fighting.Opponents) == 0 {
				return fighting.SuccessfulChoice
			}
		}

		for _, opponent := range fighting.Opponents {
			if oppose(opponent, CurrentGame.Player, color.Red) {
				return fighting.FailureChoice
			}
		}

	}
}

func removeOpponent(s []*Character, index int) []*Character {
	return append(s[:index], s[index+1:]...)
}
func oppose(attacker *Character, defender *Character, output func(format string, a ...interface{})) bool {
	roll := diceRoll()

	if len(defender.Photo) > 0 {
		StoreInHistory(IMAGE, defender.Photo)
	}
	sentence := "%v(%v) attaque %v(%v) (%v%v)"
	arguments := []interface{}{attacker.Name, attacker.HP, defender.Name, defender.HP, emoji.GameDie, roll}
	if roll >= defender.DEF {
		defender.HP -= attacker.ATK
		sentence = sentence + " et lui inflige %v points de dégât"
		arguments = append(arguments, attacker.ATK)
	} else {
		sentence = sentence + " mais ne lui inflige aucun dégât"
	}
	if roll >= CRITICAL {
		defender.HP -= attacker.ATK
		sentence = sentence + " doublé car critique"
	}
	dead := false
	if defender.HP <= 0 {
		sentence = sentence + " et l'achève."
		dead = true
	}
	GamePrint(QUOTE, output, sentence, arguments...)
	return dead
}

const DICE = 20
const CRITICAL = 20

func diceRoll() int {
	return rand.Intn(DICE) + 1
}
