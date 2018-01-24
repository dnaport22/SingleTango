package core

import (
	"bufio"
	"os"
	"strings"
)

type Listener struct {
	welcomeMessage string
	typeOf string
	userInput string
}

func (ls *Listener) Listen() {
	if ls.GetType() == typeOne {
		reader := bufio.NewReader(os.Stdin)
		StdPrint("How can I help?")
		input, _ := reader.ReadString('\n')
		ls.SetUserInput(input)
	}
}

func (ls *Listener) SetType(t string) {
	ls.typeOf = strings.ToLower(t)
}

func (ls *Listener) GetType() string {
	return ls.typeOf
}

func (ls *Listener) SetUserInput(in string) {
	ls.userInput = in
}

func (ls *Listener) GetUserInput() string {
	return ls.userInput
}