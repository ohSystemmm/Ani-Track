package Services

import (
	"os"
	"strings"
)

func Download(content string, filepath string) error {
	if strings.HasPrefix(filepath, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		filepath = strings.Replace(filepath, "~", homeDir, 1)
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}
