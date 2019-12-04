package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	ClustalW string
}

func LoadConfig() *Config {
	file, err := os.Open("./config/config.json")

	if err != nil {
		fmt.Println("Error opening config file.")
	}

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var config Config

	json.Unmarshal(byteValue, &config)

	return &config
}
