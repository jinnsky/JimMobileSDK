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
  Error *ResponseError
}

type VerifyEmailResponseListener interface {
	OnSuccess(respData *VerifyEmailResponse)
	OnFailure(respErr *ResponseError)
}

func (c *Client) SendVerifyEmail(email string) (*VerifyEmailResponse) {
  payload := VerifyEmailParams{ AppID: c.AppID, Email: email }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + VerifyEmailRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()

  respData := &VerifyEmailResponse{}
  
  respErr := c.processResponse(resp, errs)
  if respErr != nil {
    respData.Error = respErr
    return respData
  }

	if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
    return nil
  }

  return respData
}

func (c *Client) SendVerifyEmailAsync(email string, listener VerifyEmailResponseListener) {
  payload := VerifyEmailParams{ AppID: c.AppID, Email: email }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + VerifyEmailRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End(func (resp gorequest.Response, body string, errs []error) {
                                         if listener != nil {
                                           respErr := c.processResponse(resp, errs)

                                           if respErr != nil {
                                             listener.OnFailure(respErr)
                                           } else {
                                             respData := &VerifyEmailResponse{}

                                             if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
                                               respErr := &ResponseError{ Key: "JSON Decode", Message: "Decode failed" }
                                               listener.OnFailure(respErr)
                                             } else {
                                               listener.OnSuccess(respData)
                                             }
                                           }
                                         }
                                       })

  if listener != nil {
    respErr := c.processResponse(resp, errs)

    if respErr != nil {
      listener.OnFailure(respErr)
    }
  }
}
