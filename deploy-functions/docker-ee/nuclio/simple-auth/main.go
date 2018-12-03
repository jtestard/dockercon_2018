package main

import (
	"encoding/json"

	"github.com/nuclio/nuclio-sdk-go"
)

// Handle a serverless request
func Handler(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
	var e CloudEvent
	if err := json.Unmarshal(event.GetBody(), &e); err != nil {
		return nil, err
	}
	if Authenticate(e.Data) {
		return "authentication successful", nil
	}
	return "authentication denied", nil
}
