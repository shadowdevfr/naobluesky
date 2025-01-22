package bluesky

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func CreateSession(username string, password string) createSessionResponse {
	requestBody, err := json.Marshal(createSessionRequest{
		Identifier: username,
		Password:   password,
	})
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("https://bsky.social/xrpc/com.atproto.server.createSession", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	respstr, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response createSessionResponse
	err = json.Unmarshal(respstr, &response)
	if err != nil {
		panic(err)
	}

	return response
}
