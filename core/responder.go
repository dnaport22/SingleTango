package core

import (
	"strings"
	"encoding/json"
)

type Responder struct {
	typeOf string
	response string
}

func (rs *Responder) Response(r []byte) (interface{}, error) {
	var jsonRes map[string]map[string]interface{}
	json.Unmarshal(r, &jsonRes)

	return jsonRes["result"]["speech"], nil
}

func (rs *Responder) GetResponse() {
	if rs.typeOf == typeOne {
		StdPrint(rs.response)
	}
}

func (rs *Responder) SetResponse(r string) {
	rs.response = r
}

func (rs *Responder) SetType(t string) {
	rs.typeOf = strings.ToLower(t)
}

