package bapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	APP_KEY    = "howieyuen"
	APP_SECRET = "yuanhao"
)

type AccessToken struct {
	Token string `json:"token"`
}

type API struct {
	URL string
}

func NewAPI(url string) *API {
	return &API{URL: url}
}

func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {
	token, err := a.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	body, err := a.httpGet(ctx, fmt.Sprintf("%s?token=%s&name=%s", "api/v1/tags", token, name))
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	path = fmt.Sprintf("%s/%s", a.URL, path)
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return body, nil
}

func (a *API) httpPost(ctx context.Context, path string, data url.Values) ([]byte, error) {
	path = fmt.Sprintf("%s/%s", a.URL, path)

	resp, err := http.PostForm(path, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (a *API) getAccessToken(ctx context.Context) (string, error) {
	values := url.Values{
		"app_key":    {APP_KEY},
		"app_secret": {APP_SECRET},
	}
	bytes, err := a.httpPost(ctx, "auth", values)
	if err != nil {
		return "", err
	}
	fmt.Println(string(bytes))
	var accessToken = AccessToken{}
	err = json.Unmarshal(bytes, &accessToken)
	if err != nil {
		return "", err
	}
	return accessToken.Token, nil
}
