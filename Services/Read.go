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

	result.WriteString("[\n")

	for i, entry := range entries {

		result.WriteString(fmt.Sprintf("  {"))
		result.WriteString(entry.ToJson())

		if i != len(entries)-1 {
			result.WriteString(" },\n")
		} else {
			result.WriteString(" }\n")
		}
	}
	result.WriteString("]")
	return result.String(), nil
}
