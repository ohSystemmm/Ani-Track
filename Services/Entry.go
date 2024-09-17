package Services

import (
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

func (entry Entry) ToJson() string {
	var item strings.Builder

	item.WriteString(fmt.Sprintf(" \"Index\":%d,", entry.Index))
	item.WriteString(fmt.Sprintf(" \"Title\":\"%s\",", entry.Title))
	item.WriteString(fmt.Sprintf(" \"Category\":\"%s\",", entry.Category))
	item.WriteString(fmt.Sprintf(" \"Rating\":\"%s\",", entry.Rating))
	item.WriteString(fmt.Sprintf(" \"Started\":\"%s\",", entry.Started))
	item.WriteString(fmt.Sprintf(" \"Finished\":\"%s\",", entry.Finished))
	item.WriteString(fmt.Sprintf(" \"Status\":\"%s\",", entry.Status))
	item.WriteString(fmt.Sprintf(" \"Progress\":\"%s\",", entry.Progress))
	item.WriteString(fmt.Sprintf(" \"Notes\":\"%s\"", entry.Notes))
	return item.String()
}
