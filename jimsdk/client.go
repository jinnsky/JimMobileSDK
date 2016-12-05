package jimsdk

import (
  "crypto/md5"
	"encoding/hex"
  "encoding/json"
	"errors"
  "fmt"
  "math"
	"net/url"
	"strconv"
  "time"

	"github.com/antonholmquist/jason"
  "github.com/juju/persistent-cookiejar"
  "github.com/parnurzeal/gorequest"
)

const (
  BindEmailRouter = "/v1/users/bind-email"
  BindPhoneRouter = "/v1/users/bind-phone"
  ChangePasswordRouter = "/v1/users/change-password"
  FacebookUserRouter = "/v1/users/is-has-facebook-user-by-access-token"
  LinkedInUserRouter = "/v1/users/is-has-linkin-user"
  LoginRouter = "/v1/users/login"
  QqUserRouter = "/v1/users/is-has-qq-user"
  RegisterRouter = "/v1/users/register"
  RegisterInfoRouter = "/v1/users/register-with-info"
  ResetPasswordEmailRouter = "/v1/users/reset-password-email"
  ResetPasswordRouter = "/v1/users/reset-password"
  ResetPasswordSmsRouter = "/v1/users/reset-password-sms"
  TwitterUserRouter = "/v1/users/is-has-twitter-user"
  UserInfoRouter = "/v1/users/info"
  UpdateBindEmailRouter = "/v1/users/send-bind-email-email"
  UpdateBindPhoneRouter = "/v1/users/send-bind-phone-sms"
  UpdateUserRouter = "/v1/users/update"
  UploadAvatarRouter = "/v1/avatar/upload"
  UploadAvatarBase64Router = "/v1/avatar/upload-base64"
  VerifyEmailRouter = "/v1/users/send-verify-email"
  VerifySmsRouter = "/v1/users/send-verify-sms"
  WeiboUserRouter = "/v1/users/is-has-sina-weibo-user"
  WeixinUserRouter = "/v1/users/is-has-weixin-user"
)

const (
  FeedbackSubmitRouter = "/v1/feedback/submit"
)

type Client struct {
  ClusterURL string
  AppID int
  JimAppID string
  JimAppSecret string
  RequestTimeout int
  serverTimestampDiff int64
  requestAgent *gorequest.SuperAgent
  cookiejar *cookiejar.Jar
}

func (c *Client) getJimAppSign() (string) {
  serverTime := strconv.FormatInt(time.Now().UnixNano() / (1000 * 1000) + c.serverTimestampDiff, 10)
  hasher := md5.New()
  hasher.Write([]byte(c.JimAppSecret + serverTime))

  return hex.EncodeToString(hasher.Sum(nil)) + "," + serverTime
}

func NewClient(clusterURL string, appID int, jimAppID string, jimAppSecret string, cookieFilePath string) (*Client, error) {
  client := &Client {
    ClusterURL: clusterURL,
    AppID: appID,
    JimAppID: jimAppID,
    JimAppSecret: jimAppSecret,
    RequestTimeout: 3 * 60 * 1000,
  }

  if clusterURL == "" {
    return client, errors.New("ClusterURL must be indicated")
  }

  request := gorequest.New().Set("Content-Type", "application/json")
  resp, _, errs := request.Post(clusterURL + "/v1/system/base-info").End()

  if errs != nil {
    return client, errors.New(clusterURL + " isn't reachable")
  }

  type apiResultType struct {
		Time int64 `json:"time"`
	}

  apiResult := &apiResultType{}
	err := json.NewDecoder(resp.Body).Decode(apiResult)

  if err != nil {
    return client, errors.New(clusterURL + " can't response basic info: " + err.Error())
  }

  if len(cookieFilePath) > 0 {
    if jar, err := cookiejar.New(&cookiejar.Options{Filename: cookieFilePath}); err == nil {
      request.Client.Jar = jar
      client.cookiejar = jar
    }
  }

  client.requestAgent = request.Set("JIM-APP-ID", jimAppID)
  client.serverTimestampDiff = apiResult.Time - time.Now().UnixNano() / (1000 * 1000)

	return client, nil
}

func (c *Client) getRequestAgent() *gorequest.SuperAgent {
  if c.RequestTimeout >= 0 {
    return c.requestAgent.Timeout(time.Duration(c.RequestTimeout) * time.Millisecond)
  }

  return c.requestAgent.Timeout(time.Duration(math.MaxUint32) * time.Millisecond)
}

func (c *Client) saveCookieJar() {
  if c.cookiejar != nil {
    domain, _ := url.Parse(c.ClusterURL)
    cookies := c.requestAgent.Client.Jar.Cookies(domain)
    
    for _, cookie := range cookies {
      if cookie.Name == "ring-session" {
        cookie.MaxAge = 365 * 24 * 60 * 60
        break
      }
    }

    c.cookiejar.SetCookies(domain, cookies)
    c.cookiejar.Save()
  }
}

func (c *Client) removeCookieJar() {
  if c.cookiejar != nil {
    domain, _ := url.Parse(c.ClusterURL)
    cookies := c.requestAgent.Client.Jar.Cookies(domain)

    for _, cookie := range cookies {
      c.cookiejar.RemoveCookie(cookie)
    }

    c.cookiejar.SetCookies(domain, cookies)
    c.cookiejar.Save()
  }
}

type UserInfoResponse struct {
  ID int64
  Username string
  RegisterTime int64
  AvatarURL string
  Email string
  EmailChecked bool
  Phone string
  PhoneChecked bool
  InfoBirthday string
  InfoCaseHistory string
  InfoNickname string
  InfoHeight int
  InfoWeight int
  InfoGender int
  Error *ResponseError
}

func (c *Client) decodeUserInfoObject(obj *jason.Object) (id int64, 
                                                          username string, 
                                                          registerTime int64, 
                                                          avatarURL string,
                                                          email string, 
                                                          emailChecked bool,
                                                          phone string,
                                                          phoneChecked bool,
                                                          birthday string,
                                                          caseHistory string,
                                                          nickname string,
                                                          height int,
                                                          weight int,
                                                          gender int) {
  id, _ = obj.GetInt64("id")
  username, _ = obj.GetString("username")
  registerTime, _ = obj.GetInt64("register-time")
  avatarURL, _ = obj.GetString("head-pic-url")
  emailChecked, _ = obj.GetBoolean("email", "checked")
  email, _ = obj.GetString("email", "email")
  phoneChecked, _ = obj.GetBoolean("phone", "checked")
  phone, _ = obj.GetString("phone", "phone")
  birthday, _ = obj.GetString("info", "birthday")
  caseHistory, _ = obj.GetString("info", "case-history")
  nickname, _ = obj.GetString("info", "nickname")

  height64, _ := obj.GetInt64("info", "height")
  weight64, _ := obj.GetInt64("info", "weight")
  gender64, _ := obj.GetInt64("info", "sex")

  height = int(height64)
  weight = int(weight64)
  gender = int(gender64)

  return
}

func (c *Client) HasValidSession() bool {
  if c.cookiejar != nil {
    domain, _ := url.Parse(c.ClusterURL)
    
    for _, cookie := range c.cookiejar.Cookies(domain) {
      if cookie.Name == "ring-session" && len(cookie.Value) > 0 {
        return true
      }    
    }
  }

  return false
}

type ResponseError struct {
  Key string `json:"key"`
  Message string `json:"message"`
}

func CatchResponseError(respError *ResponseError) bool {
    return (respError != nil) 
}

func (c *Client) processResponse(resp gorequest.Response, errs []error) *ResponseError {
  respError := &ResponseError{}

  if errs != nil {
    respError.Key = "Unexpected errors"
    respError.Message = fmt.Sprint(errs)

    return respError
  }

  if resp.StatusCode == 422 {
    if err := json.NewDecoder(resp.Body).Decode(respError); err == nil {
      return respError
    }
  }

  return nil
}
