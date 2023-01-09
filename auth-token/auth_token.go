package auth_token

import (
	"crypto/sha1"
	"fmt"
	credentialStorage "github.com/ducknightii/api-auth/credential-storage"
	"time"
)

// AuthToken notes 负责鉴权相关计算
type AuthToken struct {
	token     string
	createdTs int64
}

func CreateAuthToken(token string, ts int64) AuthToken {
	return AuthToken{
		token:     token,
		createdTs: ts,
	}
}

func GenAuthToken(appID, baseUrl string, ts int64, storageIns credentialStorage.Storage) (AuthToken, error) {
	secret, err := storageIns.GetPassword(appID)
	if err != nil {
		return AuthToken{}, err
	}

	return AuthToken{
		token:     genToken(appID, secret, baseUrl, ts),
		createdTs: ts,
	}, nil
}

func genToken(appID, secret, baseUrl string, ts int64) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%s%d%s%s", baseUrl, ts, appID, secret))))
}

func (a AuthToken) GetToken() string {
	return a.token
}

func (a AuthToken) IsExpired(interval int64) bool {

	return time.Now().Unix()-a.createdTs < interval
}

func (a AuthToken) IsMatch(t AuthToken) bool {

	return true
}
