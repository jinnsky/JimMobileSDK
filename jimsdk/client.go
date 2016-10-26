package jimsdk

import (
  "crypto/md5"
	"encoding/hex"
  "encoding/json"
	"errors"
  "fmt"
  "math"
	"strconv"
  "time"

  "github.com/parnurzeal/gorequest"
)

const (
  BindEmailRouter = "/v1/users/bind-email"
  BindPhoneRouter = "/v1/users/bind-phone"
  ChangePasswordRouter = "/v1/users/change-password"
  FacebookUserRouter = "/v1/users/is-has-facebook-user-by-access-token"
  LinkedInUserRouter = "/v1/users/is-has-linkin-user"
  LoginRouter = "/v1/users/login"
  QqUserRouter = "/v1/users/is-has-qq-user"
  RegisterRouter = "/v1/users/register"
  TwitterUserRouter = "/v1/users/is-has-twitter-user"
  VerifyEmailRouter = "/v1/users/send-verify-email"
  VerifySmsRouter = "/v1/users/send-verify-sms"
  WeiboUserRouter = "/v1/users/is-has-sina-weibo-user"
  WeixinUserRouter = "/v1/users/is-has-weixin-user"
)

type Client struct {
  ClusterURL string
  AppID int
  JimAppID string
  JimAppSecret string
  ServerTimestampDiff int64
  RequestTimeout int
  requestAgent *gorequest.SuperAgent
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

  request := gorequest.New().Set("Content-Type", "application/json")
  resp, _, errs := request.Post(clusterURL + "/v1/system/base-info").End()

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

  client.requestAgent = request.Set("JIM-APP-ID", jimAppID)
  client.ServerTimestampDiff = apiResult.Time - time.Now().UnixNano() / (1000 * 1000)

	return client, nil
}

func (c *Client) getRequestAgent() *gorequest.SuperAgent {
  if c.RequestTimeout >= 0 {
    return c.requestAgent.Timeout(time.Duration(c.RequestTimeout) * time.Millisecond)
  }

  return c.requestAgent.Timeout(time.Duration(math.MaxUint32) * time.Millisecond)
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
