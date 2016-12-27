package jimsdk

import (
  "github.com/antonholmquist/jason"
)

type LinkedInUserParams struct {
  AppID int `json:"app-id"`
  AccessToken string `json:"access-token"`
}

type LinkedInUserResponse struct {
  Result bool
  InfoID string
  InfoName string
  InfoEmail string
  InfoGender string
  Error *ResponseError
}

func (c *Client) SendLinkedInUser(accessToken string) (*LinkedInUserResponse) {
  payload := LinkedInUserParams{ AppID: c.AppID, AccessToken: accessToken }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + LinkedInUserRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()

  respData := &LinkedInUserResponse{}
  
  respErr := c.processResponse(resp, errs)
  if respErr != nil {
    respData.Error = respErr
    return respData
  }

  v, _ := jason.NewObjectFromReader(resp.Body)

  result, _ := v.GetBoolean("result")
  id, _ := v.GetString("thirdpart-user-info", "id")
  name, _ := v.GetString("thirdpart-user-info", "name")
  email, _ := v.GetString("thirdpart-user-info", "email")
  gender, _ := v.GetString("thirdpart-user-info", "gender")

  respData.Result = result
  respData.InfoID = id
  respData.InfoName = name
  respData.InfoEmail = email
  respData.InfoGender = gender

  return respData
}
