package main

import (
	"encoding/json"
	"log"
	"os"
)

type postcard struct {
	UUID        string `json:"uuid"`
	Created     string `json:"created"`
	Textmessage string `json:"textmessage"`
}

type postcards struct {
	Postcards []postcard `json:"postcards"`
}

var postcardz postcards
var postcardzFile string

func readPostcards() error {
	data, err := os.ReadFile(postcardzFile)
	if err != nil {
		log.Print(err)
		return err
	}
	err = json.Unmarshal(data, &postcardz)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func safePostcards() error {
	bytez, err := json.MarshalIndent(postcardz, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(postcardzFile, bytez, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (pc postcard) HasContent() bool {
	return pc.Textmessage != ""
}
