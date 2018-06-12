package models

type Info struct {
	Version string `json:"version"`

	Build struct {
		NodeVersion string `json:"nodeVersion"`
		Arch        string `json:"arch"`
		Platform    string `json:"platform"`
		Cpus        int    `json:"cpus"`
	} `json:"build"`

	Commit struct {
		Hash    string `json:"hash"`
		Date    string `json:"date"`
		Author  string `json:"author"`
		Subject string `json:"subject"`
		Tag     string `json:"tag"`
		Branch  string `json:"branch"`
	} `json:"commit"`
}
