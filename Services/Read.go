package Services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Read() (string, error) {
	file, err := os.Open("Data/Ani-Track.json")
	if err != nil {
		return "", err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	var entries []Entry
	if err := json.Unmarshal(bytes, &entries); err != nil {
		return "", err
	}

	var result strings.Builder
	
	for _, entry := range entries {
		result.WriteString(fmt.Sprintf("Index: %d\n", entry.Index))
		result.WriteString(fmt.Sprintf("Title: %s\n", entry.Title))
		result.WriteString(fmt.Sprintf("Category: %s\n", entry.Category))
		result.WriteString(fmt.Sprintf("Rating: %s\n", entry.Rating))
		result.WriteString(fmt.Sprintf("Started: %s\n", entry.Started))
		result.WriteString(fmt.Sprintf("Finished: %s\n", entry.Finished))
		result.WriteString(fmt.Sprintf("Status: %s\n", entry.Status))
		result.WriteString(fmt.Sprintf("Notes: %s\n\n", entry.Notes))
	}

	return result.String(), nil
}
