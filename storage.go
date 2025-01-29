package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var choreFile = "chores.json"

func loadChores() ChoreManager {
	data, err := ioutil.ReadFile(choreFile)
	if err != nil {
		if os.IsNotExist(err) {
			return ChoreManager{Chores: []Chore{}}
		}
		log.Fatalf("Failed to read chores: %v", err)
	}
	var manager ChoreManager
	if err := json.Unmarshal(data, &manager); err != nil {
		log.Fatalf("Failed to unmarshal chores: %v", err)
	}
	return manager
}

func saveChores(manager ChoreManager) {
	data, err := json.MarshalIndent(manager, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal chores: %v", err)
	}
	if err := ioutil.WriteFile(choreFile, data, 0644); err != nil {
		log.Fatalf("Failed to save chores: %v", err)
	}
}
