package function

import (
	"encoding/json"
)

// Handle a serverless request
func Handle(req []byte) string {
	var e CloudEvent
	if err := json.Unmarshal(req, &e); err != nil {
		panic(err)
	}
	if Authenticate(e.Data) {
		return "authentication successful"
	}
	return "authentication denied"
}
