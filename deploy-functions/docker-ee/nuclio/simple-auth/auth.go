package main

func Authenticate(authData map[string]string) bool {
	if authData["user"] == "alice" && authData["authToken"] == "0af23ed13" {
		return true
	}
	return false
}

type CloudEvent struct {
	SpecVersion string            `json:"specversion"`
	Type        string            `json:"type"`
	Source      string            `json:"source"`
	Id          string            `json:"id"`
	ContentType string            `json:"contenttype"`
	Data        map[string]string `json:"data"`
}
