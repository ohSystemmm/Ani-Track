package Services

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
