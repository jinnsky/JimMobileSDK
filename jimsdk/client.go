package jimsdk

import (
  "encoding/json"
	"errors"
  "github.com/parnurzeal/gorequest"
)

type Client struct {
  ClusterURL string
  AppID string
  JimAppID string
  JimAppSecret string
  ServerTimestamp int64
}

func NewClient(clusterURL string, appID string, jimAppID string, jimAppSecret string) (*Client, error) {
  client := &Client {
    ClusterURL: clusterURL,
    AppID: appID,
    JimAppID: jimAppID,
    JimAppSecret: jimAppSecret,
  }

  if clusterURL == "" {
    return client, errors.New("ClusterURL must be indicated")
  }

  resp, _, errs := gorequest.New().Post(clusterURL + "/v1/system/base-info").Set("Content-Type", "application/json").End()

  if errs != nil {
    return client, errors.New(clusterURL + " isn't reachable")
  }

  type apiResultType struct {
		Time int64 `json:"time"`
	}

  apiResult := &apiResultType{}
	err := json.NewDecoder(resp.Body).Decode(apiResult)

  if err != nil {
    return client, errors.New(clusterURL + " can't response basic info: " + err.Error())
  }

  client.ServerTimestamp = apiResult.Time

	return client, nil
}
