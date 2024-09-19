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

func (entry Entry) ToString() string {
	var entity strings.Builder
	entity.WriteString(fmt.Sprintf("Index: %v, Title: %v, Category: %v, Rating: %v, Started: %v, Finished: %v, Status: %v, Progress: %v, Notes: %v",
		entry.Index, entry.Title, entry.Category, entry.Rating, entry.Started, entry.Finished, entry.Status, entry.Progress, entry.Notes))
	return entity.String()
}
