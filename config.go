package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	AddressListen  string `json:"addressListen"`
	AddressQr      string `json:"addressQr"`
	Salvation      string `json:"salvation"`
	PlaceholderImg string `json:"placeholderImg"`
}

var config Config

func init() {
	content, err := os.ReadFile("./config.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("using address listen", config.AddressListen)
	log.Println("using address qr", config.AddressQr)
	log.Printf("using salvation '%s'", config.Salvation)
}
