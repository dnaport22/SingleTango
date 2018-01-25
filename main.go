package main

import (
	"SingleTango/core"
	"fmt"
)

func main() {
	agent_input := core.Listener{}
	agent_output := core.Responder{}
	df := core.DialogFlow{}

	agent_input.SetType("std")
	agent_output.SetType("std")

	df.SetSessionId("1234")
	df.SetLang("en")

	for {
		agent_input.Listen()
		df.SetQuery(agent_input.GetUserInput())
		res, _ := df.MakeRequest()
		reply, err := agent_output.Response(res)
		if err != nil {
			fmt.Println(reply)
		}
		fmt.Println(reply)
	}
}