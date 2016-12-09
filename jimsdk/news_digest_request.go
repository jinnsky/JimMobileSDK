package jimsdk

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
)

type NewsDigestParams struct {
  AppID int `json:"app-id"`
  FromPage int `json:"from-page"`
  PageSize int `json:"page-size"`
  ThumbWidth int `json:"pic-w,omitempty"`
  ThumbHeight int `json:"pic-h,omitempty"`
  Tags string `json:"tags,omitempty"`
  Language string `json:"language,omitempty"`
}

func NewNewsDigestParams() (*NewsDigestParams) {
  return &NewsDigestParams{}
}

type NewsDigestResponse struct {
  Collection *NewsDigestCollection
  Error *ResponseError
}

type NewsDigest struct {
  Title string
  ArticleURL string
  ThumbURL string
}

type NewsDigestCollection struct {
  Items []*NewsDigest
}

func (n *NewsDigestCollection)GetSize() int {
  return len(n.Items)
}

func (n *NewsDigestCollection)GetItemAt(index int) (*NewsDigest) {
  return n.Items[index]
}

func (c *Client) generateNewsDigestResponseData(resp gorequest.Response, errs []error) (*NewsDigestResponse) {
  respData := &NewsDigestResponse{}
  
  respErr := c.processResponse(resp, errs)
  if respErr != nil {
    respData.Error = respErr
    return respData
  }

  type originalNewsDigest struct {
    Title string `json:"title"`
    Author string `json:"author"`
    Content  string `json:"content"`
    Tag string `json:"tag"`
    ScalePicURL string `json:"scale-pic-url"`
    DetailPageSubPath string `json:"detail-page-sub-path"`
    UpdateTime string `json:"update-time"`
  }

  var originals []originalNewsDigest

  if err := json.NewDecoder(resp.Body).Decode(&originals); err != nil {
    respData.Error = &ResponseError{ Key: "JSON Decode", Message: "Decode failed" }
  } else {
    respData.Collection = &NewsDigestCollection{}
    respData.Collection.Items = make([]*NewsDigest, len(originals))

    for i, original := range originals {
      newsDigest := &NewsDigest{}
      newsDigest.Title = original.Title
      newsDigest.ArticleURL = c.ClusterURL + original.DetailPageSubPath
      newsDigest.ThumbURL = original.ScalePicURL
      respData.Collection.Items[i] = newsDigest
    }
  }

  return respData
}

func (c *Client) SendNewsDigest(params *NewsDigestParams) (*NewsDigestResponse) {
  params.AppID = c.AppID

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + NewsDigestRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Send(params).
                                       End()

  return c.generateNewsDigestResponseData(resp, errs)
}

type NewsDigestResponseListener interface {
  OnSuccess(respData *NewsDigestResponse)
  OnFailure(respErr *ResponseError)
}

func (c *Client) SendNewsDigestAsync(params *NewsDigestParams, listener NewsDigestResponseListener) {
  callback := func (resp gorequest.Response, body string, errs []error) {
    if listener != nil {
      respData := c.generateNewsDigestResponseData(resp, errs)

      if respData.Error != nil {
        listener.OnFailure(respData.Error)
      } else {
        listener.OnSuccess(respData)
      }
    }
  }

  params.AppID = c.AppID

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + NewsDigestRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Send(params).
                                       End(callback)

  if listener != nil {
    respErr := c.processResponse(resp, errs)

    if respErr != nil {
      listener.OnFailure(respErr)
    }
  }
}
