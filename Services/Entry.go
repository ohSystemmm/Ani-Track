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
	var entity strings.Builder
	entity.WriteString(fmt.Sprintf(` "Index":%d, "Title":"%s", "Category":"%s", "Rating":"%s", "Started":"%s", "Finished":"%s", "Status":"%s", "Progress":"%s", "Notes":"%s"`,
		entry.Index, entry.Title, entry.Category, entry.Rating, entry.Started, entry.Finished, entry.Status, entry.Progress, entry.Notes))
	return entity.String()
}
