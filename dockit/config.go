package main

import (
	"encoding/json"
	"log"
	"os"
)

var _ = log.Print // for debugging, remove

type Service struct {
	Image string
	Ports map[string]string
	Deps  []string
	Env   map[string]string
}

func parseConfig(file string) (map[string]Service, error) {
	var cfg = make(map[string]Service)

	configFile, err := os.Open(file)
	if err != nil {
		return cfg, err
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
