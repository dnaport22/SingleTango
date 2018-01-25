package core

import (
	"bufio"
	"os"
	"strings"
	"regexp"
)

const triggerTerm = "yo tango"

type Listener struct {
	welcomeMessage string
	typeOf string
	userInput string
}

func (ls *Listener) Listen() {
	if ls.GetType() == typeOne {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if Trigger(input) {
			ls.SetUserInput(input)
		}
	}
}

func Trigger(s string) bool {
	re := regexp.MustCompile(triggerTerm)
	if re.FindString(s) != "" {
		StdPrint("How can I help?")
		return true
	}

	return false
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