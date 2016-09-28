package jimsdk

import (
	"errors"
	"encoding/json"
	"github.com/parnurzeal/gorequest"
)

type Client struct {
	Origin string `json:"origin"`
	URL string `json:"url"`
}

func NewClient() (*Client, error) {
	resp, _, errs := gorequest.New().Get("http://httpbin.org/get").End()

	if errs != nil {
		client := &Client{ Origin: "", URL: "" }
		return client, errors.New("Request failed.")
	}

	var client = new(Client)
	json.NewDecoder(resp.Body).Decode(&client)

	return client, nil
}
