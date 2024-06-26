package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type postcard struct {
	UUID        string `json:"uuid"`
	Created     string `json:"created"`
	Scanned     bool   `json:"scanned"`
	Textmessage string `json:"textmessage"`
}

type postcards struct {
	Postcards []postcard `json:"postcards"`
}

var postcardz postcards
var postcardzFile string
var pcmu sync.RWMutex

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
	pcmu.Lock()
	defer pcmu.Unlock()

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

func getPostcardByUUID(uuid string) (*postcard, error) {
	pcmu.RLock()
	defer pcmu.RUnlock()
	for i, p := range postcardz.Postcards {
		if p.UUID == uuid {
			return &postcardz.Postcards[i], nil
		}
	}
	return nil, fmt.Errorf("no postcard found with uuid %s", uuid)
}
