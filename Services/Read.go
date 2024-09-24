package Services

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// filepath = Data/Ani-Track.json

func Read(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	var entries []Entry
	if err := json.Unmarshal(content, &entries); err != nil {
		return "", err
	}

	entriesJson, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return "", err
	}

	return string(entriesJson), nil
}
