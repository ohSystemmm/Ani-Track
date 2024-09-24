package Services

import (
	"io/ioutil"
	"os"
)

// Save takes JSON content as a string and a file path, and writes the content to the file.
func Save(content string, filepath string) error {
	// Create or overwrite the file
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the content to the file
	err = ioutil.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}
