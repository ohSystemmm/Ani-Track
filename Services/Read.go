package Services

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

func Read() (string, error) {
	filepath := "~/.local/state/ani-track/Ani-Track.json"
	if strings.HasPrefix(filepath, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		filepath = strings.Replace(filepath, "~", homeDir, 1)
	}

	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
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
