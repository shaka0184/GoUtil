package zoomUtil

import (
	"encoding/base64"
	"encoding/json"
	"github.com/pkg/errors"
	httpUtil "github.com/shaka0184/GoUtil/pkg/http"
	"net/http"
	"os"
)

const baseUrl = "https://api.zoom.us/v2"

type OauthRequest struct {
	ClientId     string
	ClientSecret string
	AccountId    string
}

type OauthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

type Client struct {
	accessToken string

	clientId     string
	clientSecret string
	accountId    string
}

// NewClient 環境変数からパラメーターを取得する
func NewClient() *Client {
	return &Client{
		clientId:     os.Getenv("ZoomClientId"),
		clientSecret: os.Getenv("ZoomClientSecret"),
		accountId:    os.Getenv("ZoomAccountId"),
	}
}

func NewClientRequestParam(oauthRequest OauthRequest) *Client {
	return &Client{
		clientId:     oauthRequest.ClientId,
		clientSecret: oauthRequest.ClientSecret,
		accountId:    oauthRequest.AccountId,
	}
}

func NewClientToken(oauthRequest *OauthRequest) (*Client, error) {
	var client *Client

	if oauthRequest != nil {
		client = NewClientRequestParam(*oauthRequest)
	} else {
		client = NewClient()
	}

	res, err := client.GetAccessToken()
	if err != nil {
		return nil, err
	}

	if res != nil {
		client.accessToken = res.AccessToken
	}

	return client, nil
}

func (c Client) GetAccessToken() (*OauthResponse, error) {
	url := "https://zoom.us/oauth/token?grant_type=account_credentials&account_id=" + c.accountId

	src := []byte(c.clientId + ":" + c.clientSecret)

	enc := base64.StdEncoding.EncodeToString(src)
	header := []httpUtil.Header{{Key: "Authorization", Value: "Basic " + enc}}

	res, err := httpUtil.Request(http.MethodPost, url, nil, header)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res)
	ret := new(OauthResponse)

	err = decoder.Decode(res)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return ret, nil
}
