package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func loadGame() *Choice {
	choice := NewChoice()
	if _, err := os.Stat(SAVE_FILE); !errors.Is(err, os.ErrNotExist) {
		savedContent, _ := ioutil.ReadFile(SAVE_FILE)
		save := &Save{}
		json.Unmarshal(savedContent, save)
		CurrentGame = save.History
		os.Remove(SAVE_FILE)
		return save.Choice
	} else {
		initGame()
		json.Unmarshal(game, choice)
	}
	return choice
}

func initGame() {
	prompt := promptui.Prompt{
		Label: "Entrez votre nom",
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(0)
	}
	StoreInHistory(TITLE, "L'histoire de %v", result)

	CurrentGame.Player = &Character{
		Name: result,
		HP:   (rand.Intn(20) + 1) * 5,
		ATK:  6 + (rand.Intn(6) + 1) + (rand.Intn(6) + 1),
		DEF:  6 + (rand.Intn(6) + 1) + (rand.Intn(6) + 1),
	}

}

func saveState(choice *Choice) {

	color.New(color.FgWhite, color.BgRed).Println("Votre partie va s'arrêter, si vous étiez en combat, celui-ci sera sauvegardé. Voulez-vous continuer?")

	prompt := promptui.Select{
		Label: "êtes vous sûr ?",
		Items: []string{"Oui", "Non"},
	}
	i, _, err := prompt.Run()
	if err != nil {
		return
	}
	if i == 0 {

		save := Save{
			History: CurrentGame,
			Choice:  choice,
		}
		item, _ := json.Marshal(save)
		ioutil.WriteFile(SAVE_FILE, item, 0777)

		os.Exit(0)
	}

}

func handleEnd(err error, choice *Choice) bool {
	if err != nil {
		if err.Error() == "^C" {
			saveState(choice)
		}
		return true
	}
	return false
}

func finishGame(choice *Choice) {
	file, _ := os.OpenFile("end-"+time.Now().Format("2006-02-01-15:04:05")+".md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	datawriter := bufio.NewWriter(file)

	datawriter.WriteString("| property | value |\n|-----------|--------|\n")
	datawriter.WriteString(fmt.Sprintf("| Name | %v |\n", CurrentGame.Player.Name))
	datawriter.WriteString(fmt.Sprintf("| ATK | %v |\n", CurrentGame.Player.ATK))
	datawriter.WriteString(fmt.Sprintf("| DEF | %v |\n", CurrentGame.Player.DEF))
	datawriter.WriteString(fmt.Sprintf("| HP | %v |\n\n\n", CurrentGame.Player.HP))

	for _, msg := range CurrentGame.Messages {
		datawriter.WriteString(msg + "\n\n")
	}
	datawriter.WriteString(fmt.Sprintf("![](https://loremflickr.com/g/300/400/%v/all)\n\n", choice.Feeling))

	datawriter.Flush()
	file.Close()
}

const (
	MESSAGE int = iota
	QUOTE
	TITLE
	IMAGE
)

func GamePrint(messageType int, printFunc func(format string, a ...interface{}), message string, a ...interface{}) {
	StoreInHistory(messageType, message, a...)
	printFunc(message, a...)
}

func StoreInHistory(messageType int, message string, a ...interface{}) {
	switch messageType {
	case MESSAGE:
		CurrentGame.Messages = append(CurrentGame.Messages, fmt.Sprintf(message, a...))
	case QUOTE:
		CurrentGame.Messages = append(CurrentGame.Messages, fmt.Sprintf("> "+message, a...))
	case TITLE:
		CurrentGame.Messages = append(CurrentGame.Messages, fmt.Sprintf("# "+message, a...))
	case IMAGE:
		CurrentGame.Messages = append(CurrentGame.Messages, fmt.Sprintf("![](%v)", message))
	}
}
