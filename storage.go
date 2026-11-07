package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadServers() Servers {

	var onloading Servers

	file, err := os.Open("servers.json")

	if os.IsNotExist(err) {
		return onloading
	}
	if err != nil {
		fmt.Println(err)
		return onloading
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&onloading)
	if err != nil {
		fmt.Println("Error decoding array:", err)
		return onloading
	}

	return onloading
}
