package jimsdk

import (
  "encoding/json"
)

type UpdateBindEmailParams struct {
  UserID int `json:"users-id"`
  Email string `json:"email"`
}

type UpdateBindEmailResponse struct {
  Result bool `json:"result"`
  Error *ResponseError
}

func (c *Client) SendUpdateBindEmail(userID int, email string) (*UpdateBindEmailResponse) {
  payload := UpdateBindEmailParams{ UserID: userID, Email: email }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + UpdateBindEmailRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()

  respData := &UpdateBindEmailResponse{}

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
