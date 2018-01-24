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

	df.SetQuery("hello")
	df.SetSessionId("1234")
	df.SetLang("en")

	for {
		agent_input.Listen()
		agent_output.SetResponse(agent_input.GetUserInput())
		res, _ := df.MakeRequest()
		fmt.Println(agent_output.Response(res))
	}
}