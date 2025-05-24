package file

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveToFile(m map[string]string) error {

	// Convert to JSON
	jsonData, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	// Write to file
	err = os.WriteFile("address.json", jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadFromFile() (map[string]string, error) {
	fileContent, err := os.ReadFile("address.json")
	if err != nil {
		panic(err)
	}
	hashedData := make(map[string]string)
	err = json.Unmarshal(fileContent, &hashedData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("hashedData: %s", hashedData)
	return hashedData, nil
}
