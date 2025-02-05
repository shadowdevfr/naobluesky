package mastodon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (m Mastodon) Post(text string) (post MastodonPost) {
	params := url.Values{}
	params.Set("status", text)
	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/api/v1/statuses", m.Instance), bytes.NewBufferString(params.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+m.Token)

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
	err = json.Unmarshal(respstr, &post)
	if err != nil {
		panic(err)
	}

	return post
}
