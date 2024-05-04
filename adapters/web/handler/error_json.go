package handler

import "encoding/json"

func jsonError(msg string) []byte {
	error := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}

	r, err := json.Marshal(error)
	if err != nil {
		return []byte(`{"message":"Error marshalling error message"}`)
	}

	return r
}
