package jimsdk

import (
	"encoding/json"
	"io/ioutil"
  "path/filepath"

	"github.com/parnurzeal/gorequest"
)

type UploadAvatarResponse struct {
  URL string `json:"head-pic-url"`
  Message string `json:"message"`
  Error *ResponseError
}

func (c *Client) SendUploadAvatar(file string) (*UploadAvatarResponse) {
  requestAgent := c.getRequestAgent().Post(c.ClusterURL + UploadAvatarRouter).
                                      Type("multipart").
                                      Set("JIM-APP-SIGN", c.getJimAppSign()).
                                      Set("JIM-APP-ID", c.JimAppID)
                                       
  if pathToFile, err := filepath.Abs(file); err != nil {
		requestAgent.Errors = append(requestAgent.Errors, err)
	} else {
    if data, err := ioutil.ReadFile(file); err != nil {
      requestAgent.Errors = append(requestAgent.Errors, err)
    } else {
      requestAgent.FileData = append(requestAgent.FileData, gorequest.File{
        Filename: filepath.Base(pathToFile),
			  Fieldname: "file",
			  Data: data,
		  })
    }
  }

  resp, _, errs := requestAgent.End()

  respData := &UploadAvatarResponse{}

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

func (c *Client) SendUploadAvatarBase64(encodedStr string) (*UploadAvatarResponse) {
  var payload = struct {
    Data string `json:"data"`
  } { encodedStr }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + UploadAvatarBase64Router).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()

  respData := &UploadAvatarResponse{}

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
