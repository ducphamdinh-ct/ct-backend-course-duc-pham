package section2

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	BaseUrl = "https://gateway.chotot.com/v1/public/ad-listing"
	CateVeh = "2000"
	CatePty = "1000"
)

func NewClient(baseUrl string, retryTimes int, log *log.Logger) *Client {
	// TODO #4 refactor NewClient using functional options
	return &Client{
		httpClient: http.DefaultClient,
		baseUrl:    baseUrl,
		retryTimes: retryTimes,
		logger:     log,
	}
}

type Client struct {
	httpClient *http.Client
	baseUrl    string
	retryTimes int
	logger     *log.Logger
}

func (c *Client) GetAdByCate() (*AdsResponse, error) {
	now := time.Now()
	defer func() {
		c.logger.Printf("GetAdByCate Request - Cate %v, Duration: %v", CateVeh, time.Since(now).String())
	}()

	url := fmt.Sprintf("%v?cg=%v&limit=10", BaseUrl, CateVeh)

	// TODO #3 implement retry if StatusCode = 5xx
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	//b, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Printf("\nResponse %v", string(b))

	var adResp AdsResponse
	// TODO #2 unmarshal json
	err = json.NewDecoder(resp.Body).Decode(adResp)
	if err != nil {
		return nil, err
	}
	return &adResp, nil
}

type AdsResponse struct {
	Total int  `json:"total"`
	Ads   []Ad `json:"ads"`
}

type Ad struct {
	AdId        int    `json:"ad_id"`
	ListId      int    `json:"list_id"`
	AccountName string `json:"account_name"`
	Subject     string `json:"subject"`
	ListTime    string `json:"list_time"`
}
