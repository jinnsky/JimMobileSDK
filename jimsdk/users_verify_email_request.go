package jimsdk

import (
  "encoding/json"
  
  "github.com/parnurzeal/gorequest"
)

type VerifyEmailResponseData struct {
  Result bool `json:"result"`
}

type VerifyEmailResponseListener interface {
	OnSuccess(respData *VerifyEmailResponseData)
	OnFailure(err string)
}

func (c *Client) SendVerifyEmail(email string) (*VerifyEmailResponseData) {
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

  respData := &VerifyEmailResponseData{}

	if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
    return nil
  }

  return respData
}

func (c *Client) SendVerifyEmailAsync(email string, listener VerifyEmailResponseListener) {
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
                      
                      respData := &VerifyEmailResponseData{}
                                       
                      if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
                        listener.OnFailure("Decode failed.")
                      }
                                       
                      listener.OnSuccess(respData)
                    }
                  })
}
