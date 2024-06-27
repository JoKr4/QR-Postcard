package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Address   string
	Salvation string
}

var config Config

func init() {
	content, err := os.ReadFile("./config")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
