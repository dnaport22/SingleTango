package main

import (
	"testing"
	"SingleTango/core"
)

func TestSingleTangoListener(t *testing.T) {
	listener := core.Listener{}
	listener.SetType("std")
	listener.SetType("en")
	listener.SetUserInput("hello")

	if listener.GetUserInput() != "hello" {
		t.Errorf("Error while setting up the intent.")
	}
}

func TestSingleTangoResponder(t *testing.T) {
	responder := core.Responder{}
	res := []byte(`{"result": {"speech":"hi"}}`)
	responder.SetType("std")
	r, err := responder.Response(res)

	if err != nil {
		t.Errorf(err.Error())
	}

	if r != "hi" {
		t.Errorf("Error while setting up the response.")
	}
}

func TestSingleTangoDialogFlow(t *testing.T) {
	listener := core.Listener{}
	responder := core.Responder{}
	df := core.DialogFlow{}

	listener.SetType("std")
	responder.SetType("std")

	df.SetSessionId("1234")
	df.SetLang("en")

	listener.Listen()
	df.SetQuery(listener.GetUserInput())
	if df.GetQuery() != listener.GetUserInput() {
		t.Errorf("Mixed intent, given: %s and processed: %s", listener.GetUserInput(), df.GetQuery())
	}
}