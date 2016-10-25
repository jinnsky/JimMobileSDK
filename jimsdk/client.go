package jimsdk

import (
  "crypto/md5"
	"encoding/hex"
  "encoding/json"
	"errors"
  "fmt"
	"strconv"
  "time"

  "github.com/parnurzeal/gorequest"
)

const (
  BindEmailRouter = "/v1/users/bind-email"
  BindPhoneRouter = "/v1/users/bind-phone"
  ChangePasswordRouter = "/v1/users/change-password"
  LoginRouter = "/v1/users/login"
  RegisterRouter = "/v1/users/register"
  VerifyEmailRouter = "/v1/users/send-verify-email"
  VerifySmsRouter = "/v1/users/send-verify-sms"
)

type Client struct {
  ClusterURL string
  AppID int
  JimAppID string
  JimAppSecret string
  ServerTimestampDiff int64
  RequestTimeout int
}

func (c *Client) getJimAppSign() (string) {
  serverTime := strconv.FormatInt(time.Now().UnixNano() / (1000 * 1000) + c.ServerTimestampDiff, 10)
  hasher := md5.New()
  hasher.Write([]byte(c.JimAppSecret + serverTime))

  return hex.EncodeToString(hasher.Sum(nil)) + "," + serverTime
}

func NewClient(clusterURL string, appID int, jimAppID string, jimAppSecret string) (*Client, error) {
  client := &Client {
    ClusterURL: clusterURL,
    AppID: appID,
    JimAppID: jimAppID,
    JimAppSecret: jimAppSecret,
  }

  if clusterURL == "" {
    return client, errors.New("ClusterURL must be indicated")
  }

  resp, _, errs := gorequest.New().
                             Post(clusterURL + "/v1/system/base-info").
                             Set("Content-Type", "application/json").
                             End()

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

  client.ServerTimestampDiff = apiResult.Time - time.Now().UnixNano() / (1000 * 1000)

	return client, nil
}

func (c *Client) getRequest() *gorequest.SuperAgent {
  request := gorequest.New().Set("Content-Type", "application/json").
                             Set("JIM-APP-ID", c.JimAppID)

  if c.RequestTimeout > 0 {
    request.Timeout(time.Duration(c.RequestTimeout) * time.Millisecond)
  }                           

  return request
}

type ResponseError struct {
  Key string `json:"key"`
  Message string `json:"message"`
}

func CatchResponseError(respError *ResponseError) bool {
    return (respError != nil) 
}

func (c *Client) processResponse(resp gorequest.Response, errs []error) *ResponseError {
  respError := &ResponseError{}

  if errs != nil {
    respError.Key = "Unexpected errors"
    respError.Message = fmt.Sprint(errs)

    return respError
  }

  if resp.StatusCode == 422 {
    if err := json.NewDecoder(resp.Body).Decode(respError); err == nil {
      return respError
    }
  }

  return nil
}
