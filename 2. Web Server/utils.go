package main


import (
	"encoding/json"
	"log"
)

func toJSON(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonData)
}