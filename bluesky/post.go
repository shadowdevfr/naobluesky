package bluesky

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Post(session createSessionResponse, text string) createSessionResponse {
	requestBody, err := json.Marshal(createRecordRequest{
		Repo:       session.Did,
		Collection: "app.bsky.feed.post",
		Record: record{
			Type:      "app.bsky.feed.post",
			CreatedAt: time.Now().Format(time.RFC3339),
			Text:      text,
		},
	})
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", "https://bsky.social/xrpc/com.atproto.repo.createRecord", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+session.AccessJwt)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	respstr, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(respstr))
	var response createSessionResponse
	err = json.Unmarshal(respstr, &response)
	if err != nil {
		panic(err)
	}

	return response
}
