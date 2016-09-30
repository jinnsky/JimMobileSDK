package jimsdk

import (
  "crypto/md5"
	"encoding/hex"
  "encoding/json"
	"strconv"
  "time"

  "github.com/parnurzeal/gorequest"
)

func (c *Client) getJimAppSign() (string) {
  serverTime := strconv.FormatInt(time.Now().UnixNano() / (1000 * 1000) + c.ServerTimestampDiff, 10)
  hasher := md5.New()
  hasher.Write([]byte(c.JimAppSecret + serverTime))

  return hex.EncodeToString(hasher.Sum(nil)) + "," + serverTime
}

type ResponseData struct {
  Result bool `json:"result"`
}

type ResponseListener interface {
	OnSuccess(respData *ResponseData)
	OnFailure(err string)
}

func (c *Client) SendVerifyEmail(email string) (*ResponseData) {
  type payloadType struct {
	  AppID int `json:"app-id"`
    Email string `json:"email"`
  }

  myPayload := payloadType{ AppID: c.AppID, Email: email }

  resp, _, errs := gorequest.New().Post(c.ClusterURL + "/v1/users/send-verify-email").
                                   Set("Content-Type", "application/json").
                                   Set("JIM-APP-ID", c.JimAppID).
                                   Set("JIM-APP-SIGN", c.getJimAppSign()).
                                   Send(myPayload).
                                   End()

  if errs != nil {
    return nil
  }

  respData := &ResponseData{}

	if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
    return nil
  }

  return respData
}

func (c *Client) SendVerifyEmailAsync(email string, listener ResponseListener) {
  type payloadType struct {
	  AppID int `json:"app-id"`
    Email string `json:"email"`
  }

  myPayload := payloadType{ AppID: c.AppID, Email: email }

  gorequest.New().Post(c.ClusterURL + "/v1/users/send-verify-email").
                  Set("Content-Type", "application/json").
                  Set("JIM-APP-ID", c.JimAppID).
                  Set("JIM-APP-SIGN", c.getJimAppSign()).
                  Send(myPayload).
                  End(func (resp gorequest.Response, body string, errs []error)  {
                    if listener != nil {
                      if errs != nil {
                        listener.OnFailure("Request failed.")
                      }
                      
                      respData := &ResponseData{}
                                       
                      if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
                        listener.OnFailure("Decode failed.")
                      }
                                       
                      listener.OnSuccess(respData)
                    }
                  })
}
