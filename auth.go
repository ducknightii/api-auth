package api_auth

import (
	"fmt"
	api_request "github.com/ducknightii/api-auth/api-request"
	auth_token "github.com/ducknightii/api-auth/auth-token"
	credential_storage "github.com/ducknightii/api-auth/credential-storage"
	"sync"
	"time"
)

type Auth struct {
	s         credential_storage.Storage
	intervalS int64
}

var once sync.Once

func Init(intervalS int64, s credential_storage.Storage) Auth {
	var a Auth

	once.Do(func() {
		if s == nil {
			s = credential_storage.ConfStorageInit(map[string]string{
				"a": "aaaa",
			})
		}
		a.s = s
		a.intervalS = intervalS
	})

	return a
}

func (a Auth) GenTokenUrl(rawUrl, appID string) (string, error) {
	ts := time.Now().Unix()
	authToken, err := auth_token.GenAuthToken(appID, rawUrl, ts, a.s)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s&app_id=%s&ts=%d&token=%s", rawUrl, appID, ts, authToken.GetToken()), nil
}

func (a Auth) AuthCheck(rawTokenUrl string) (bool, error) {
	apiReq, err := api_request.Parse2ApiRequest(rawTokenUrl)
	if err != nil {
		return false, err
	}

	a1 := auth_token.CreateAuthToken(apiReq.GetToken(), apiReq.GetTs())
	if !a1.IsExpired(a.intervalS) {
		return false, nil
	}
	a2, err := auth_token.GenAuthToken(apiReq.GetAppID(), apiReq.GetBaseUrl(), apiReq.GetTs(), a.s)
	if err != nil {
		return false, err
	}

	return a1.IsMatch(a2), nil
}
