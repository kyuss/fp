package main

import "encoding/json"

type Result struct {
	Url string `json:"url"`
}

func unmarshall(payload []byte, result *Result) {
	json.Unmarshal(payload, &result)
}
