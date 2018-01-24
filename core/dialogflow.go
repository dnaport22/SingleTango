package core

import (
	"net/http"
	"bytes"
	"os"
	"log"
	"io/ioutil"
)

type DialogFlow struct {
	query string
	lang string
	sessionId string
}

const endpoint  = "https://api.api.ai/v1/query"
const envKeyVar = "SINGLE_TANGO_KEY"

func (df *DialogFlow) MakeRequest() ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(df.GetIntent()))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "bearer" + df.GetClientAccessToken())
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return ResponseBody(resp)
}

func ResponseBody(res *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (df *DialogFlow) GetIntent() []byte {
	return []byte(`{"query": "` + df.query + `", "lang": "` + df.lang + `", "sessionId": "` + df.sessionId + `"}`)
}

func (df *DialogFlow) GetClientAccessToken() string {
	return os.Getenv(envKeyVar)
}

func (df *DialogFlow) SetQuery(q string) {
	df.query = q
}

func (df *DialogFlow) SetSessionId(sid string) {
	df.sessionId = sid
}

func (df *DialogFlow) SetLang(lang string) {
	df.lang = lang
}