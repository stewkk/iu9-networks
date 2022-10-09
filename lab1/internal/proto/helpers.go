package proto

import (
	"encoding/json"
	"io"
)

func MakeRequest(w io.Writer, command string, payload any) error {
	var data json.RawMessage
	data, _ = json.Marshal(payload)
	return json.NewEncoder(w).Encode(&Request{
		Command: command,
		Data:    &data,
	})
}

func MakeResponse(w io.Writer, status string, payload any) error {
	var data json.RawMessage
	data, _ = json.Marshal(payload)
	return json.NewEncoder(w).Encode(&Response{
		Status: status,
		Data:   &data,
	})
}
