package bapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wsloong/tag-service/pkg/errcode"
	"io/ioutil"
	"net/http"
)

const (
	APPKEY = "eddycjy"
	APPSECRET = "go-programming-tour-book"
)

type API struct {
	URL string
}

func NewAPI(url string) *API {
	return &API{URL:url}
}

type AccessToken struct {
	Token string `json:"token"`
}

// 获取token，这个可以缓存起来
func (a *API) getAccessToken(ctx context.Context) (string, error) {
	body, err := a.httpGet(ctx, fmt.Sprintf("%s?app_key=%s&app_secret=%s", "auth", APPKEY, APPSECRET))
	if err != nil {
		return "", err
	}

	var accessToken AccessToken
	_ = json.Unmarshal(body, &accessToken)
	return accessToken.Token, nil
}

func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", a.URL, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {
	token, err := a.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	body, err := a.httpGet(ctx, fmt.Sprintf("%s?token=%s&name=%s", "api/v1/tags", token, name))
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
		//return nil, err
	}
	return body, err
}