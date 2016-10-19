package jimsdk

import (
  "encoding/json"
  
  "github.com/parnurzeal/gorequest"
)

type VerifyEmailParams struct {
  AppID int `json:"app-id"`
  Email string `json:"email"`
}

type VerifyEmailResponse struct {
  Result bool `json:"result"`
}

type VerifyEmailResponseListener interface {
	OnSuccess(respData *VerifyEmailResponse)
	OnFailure(err string)
}

func (c *Client) SendVerifyEmail(email string) (*VerifyEmailResponse) {
  payload := VerifyEmailParams{ AppID: c.AppID, Email: email }

  resp, _, errs := gorequest.New().Post(c.ClusterURL + "/v1/users/send-verify-email").
                                   Set("Content-Type", "application/json").
                                   Set("JIM-APP-ID", c.JimAppID).
                                   Set("JIM-APP-SIGN", c.getJimAppSign()).
                                   Send(payload).
                                   End()

  if errs != nil {
    return nil
  }

  respData := &VerifyEmailResponse{}

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
                      
                      respData := &VerifyEmailResponse{}
                                       
                      if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
                        listener.OnFailure("Decode failed.")
                      }
                                       
                      listener.OnSuccess(respData)
                    }
                  })
}
