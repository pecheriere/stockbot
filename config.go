package main

import (
	"os"
	"encoding/json"
	"fmt"
	"reflect"
	"go/types"
)

type Config struct {
	Auth struct {
		Username string
		Password string
	}
}

func Configuration() Config {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}
	return config
}