package core

import (
	"bufio"
	"os"
	"strings"
	"regexp"
	"time"
)

const triggerTerm = "yo tango"

type Listener struct {
	welcomeMessage string
	typeOf string
	userInput string
	sessionId string
	sessionState bool
}

func (ls *Listener) Listen() {
	if ls.GetType() == typeOne {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if !ls.IsSessionActive() {
			if Trigger(input) {
				ls.StartSession()
				ls.SetUserInput(input)
			}
		} else {
			ls.SetUserInput(input)
		}
	}
}

func Trigger(s string) bool {
	re := regexp.MustCompile(triggerTerm)
	if re.FindString(s) != "" {
		return true
	}

	return false
}

func (ls *Listener) IsSessionActive() bool {
	return ls.sessionState
}

func (ls *Listener) StartSession() {
	ls.sessionState = true
	t := time.Now().Format("20060102150405")
	ls.sessionId = t
}

func (ls *Listener) GetSessionId() string {
	return ls.sessionId
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