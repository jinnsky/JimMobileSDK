package jimsdk

import (
	"encoding/json"
)

type FeedbackSubmitParams struct {
  ContactInfo string `json:"contact-info"`
  Content string `json:"content"`
}

type FeedbackSubmitResponse struct {
  ContactInfo string `json:"contact-info"`
  Content string `json:"content"`
  Error *ResponseError
}

func (c *Client) SendFeedback(contactInfo string, content string) (*FeedbackSubmitResponse) {
  payload := FeedbackSubmitParams{ ContactInfo: contactInfo, Content: content }
  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + FeedbackSubmitRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Send(payload).
                                       End()

  respData := &FeedbackSubmitResponse{}

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
