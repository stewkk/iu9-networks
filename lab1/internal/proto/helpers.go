package proto

import (
	"encoding/json"
	"io"
)

func MakeRequest(w io.Writer, command string, payload any) {
	var data json.RawMessage
	data, _ = json.Marshal(payload)
	json.NewEncoder(w).Encode(&Request{
		Command: command,
		Data:    &data,
	})
}

func MakeResponse(w io.Writer, status string, payload any) {
	var data json.RawMessage
	data, _ = json.Marshal(payload)
	json.NewEncoder(w).Encode(&Response{
		Status: status,
		Data:   &data,
	})
}
