package jimsdk

type FaqListParams struct {
  AppID int `json:"app-id"`
  Language string `json:"language"`
  From int `json:"from,omitempty"`
  Size int `json:"size,omitempty"`
}

func NewFaqListParams() (*FaqListParams) {
  return &FaqListParams{}
}

type FaqListResponse struct {
  Text string
  Error *ResponseError
}

func (c *Client) SendFaqList(params *FaqListParams) (*FaqListResponse) {
  params.AppID = c.AppID

  resp, body, errs := c.getRequestAgent().Post(c.ClusterURL + FaqListRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(params).
                                       End()

  respData := &FaqListResponse{}
  
  respErr := c.processResponse(resp, errs)
  if respErr != nil {
    respData.Error = respErr
    return respData
  }

  respData.Text = body

  return respData
}
