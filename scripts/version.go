package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"path/filepath"
)

func main() {
	getVersion()
}

type PackageJson struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func getVersion() {
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		log.Fatalf("Unable to get working dir: %v", wdErr)
	}

	filePath := filepath.Join(wd, "npm/package.json")

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Unable to read package.json at filepath: %s\nError received: %v", filePath, err)
	}

	var pkgJson PackageJson
	jsonErr := json.Unmarshal(bytes, &pkgJson)
	if jsonErr != nil {
		fmt.Printf("JSON Unmarshal Error: %v", jsonErr)
	}

	fmt.Print(pkgJson.Version)
}
