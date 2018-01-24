package core

import (
	"strings"
	"encoding/json"
)

type Responder struct {
	typeOf string
	response string
}

func (rs *Responder) Response(r []byte) (string, error) {
	var jsonRes map[string]interface{}
	json.Unmarshal(r, &jsonRes)
	res, err := json.Marshal(jsonRes)

	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (rs *Responder) GetResponse() {
	if rs.typeOf == typeOne {
		StdPrint(rs.response)
	}
}

func (rs *Responder) SetResponse(r string) {
	rs.response = r
}

func (rs *Responder) SetType(val string) {
	rs.typeOf = strings.ToLower(val)
}

