package Services

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Entry struct {
	Index    uint64 `json:"Index"`
	Title    string `json:"Title"`
	Category string `json:"Category"`
	Rating   string `json:"Rating"`
	Started  string `json:"Started"`
	Finished string `json:"Finished"`
	Status   string `json:"Status"`
	Progress string `json:"Progress"`
	Notes    string `json:"Notes"`
}

func (entry Entry) ToJson() (string, error) {
	jsonData, err := json.Marshal(entry)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// Not Finished Yet

func (entry Entry) ToString() string {
	var entity strings.Builder
	entity.WriteString(fmt.Sprintf(""))
	return entity.String()
}
