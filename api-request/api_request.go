package api_request

import (
	"fmt"
	"net/url"
	"strconv"
)

// ApiRequest 负责解析url 生成需要元素
type ApiRequest struct {
	baseUrl string
	token   string
	appID   string
	ts      int64
}

func Parse2ApiRequest(str string) (req ApiRequest, err error) {
	urlIns, err := url.Parse(str)
	if err != nil {
		return ApiRequest{}, err
	}

	path := urlIns.Path
	query := urlIns.RawQuery
	q, _ := url.ParseQuery(query)
	appID := q.Get("app_id")
	_ts := q.Get("ts")
	ts, _ := strconv.Atoi(_ts)
	token := q.Get("token")
	q.Del("app_id")
	q.Del("ts")
	q.Del("token")
	baseUrl := fmt.Sprintf("%s?%s", path, q.Encode())

	return ApiRequest{
		baseUrl: baseUrl,
		token:   token,
		appID:   appID,
		ts:      int64(ts),
	}, nil
}

func (a ApiRequest) GetToken() string {
	return a.token
}

func (a ApiRequest) GetTs() int64 {
	return a.ts
}

func (a ApiRequest) GetAppID() string {
	return a.appID
}

func (a ApiRequest) GetBaseUrl() string {
	return a.baseUrl
}
