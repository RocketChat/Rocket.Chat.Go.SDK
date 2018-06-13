package models

import "time"

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

type Directory struct {
	Result []struct {
		ID        string    `json:"_id"`
		CreatedAt time.Time `json:"createdAt"`
		Emails    []struct {
			Address  string `json:"address"`
			Verified bool   `json:"verified"`
		} `json:"emails"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"result"`

	Count  int `json:"count"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type Spotlight struct {
	Users []User    `json:"users"`
	Rooms []Channel `json:"rooms"`
}
