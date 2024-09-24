package Services

import (
	"io/ioutil"
	"os"
)

func Save(content string, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = ioutil.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}
