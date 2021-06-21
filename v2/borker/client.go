package broker

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/Zamzam-Technology/go-binance/v2/common"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"

	// Endpoints
	baseAPIMainURL    = "https://api.binance.com"
	baseAPITestnetURL = "https://testnet.binance.vision"

	defaultTimeout = 5 * time.Second
)

// UseTestnet switch all the API endpoints from production to the testnet
var UseTestnet = false

type doFunc func(req *http.Request) (*http.Response, error)

type Client struct {
	HTTPClient *http.Client
	apiKey     string
	secretKey  string
	BaseURL    string
	TimeOffset int64
	Logger     *log.Logger
	Debug      bool
	do         doFunc
}

// NewClient creates new broker client
func NewClient(apiKey, secretKey string, writer io.Writer) *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: defaultTimeout,
		},
		apiKey:    apiKey,
		secretKey: secretKey,
		BaseURL:   getAPIEndpoint(),
		Logger:    log.New(writer, "Binance-Broker", log.LstdFlags),
	}
}

// CreateSubAccount creates new sub account
func (c *Client) CreateSubAccount(ctx context.Context, opts ...RequestOption) (res *SubAccount, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/subAccount",
		secType:  secTypeSigned,
	}
	data, err := c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccount)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// EnableFuturesSubAccount make request
func (c *Client) EnableFuturesSubAccount(ctx context.Context, subAccountId int, opts ...RequestOption) (res *EnableFuturesToSubAccountResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/subAccount/futures",
		secType:  secTypeSigned,
	}

	r.setParam("subAccountId", subAccountId)
	r.setParam("futures", true)

	data, err := c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EnableFuturesToSubAccountResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// EnableMarginSubAccount make request
func (c *Client) EnableMarginSubAccount(ctx context.Context, subAccountId int, opts ...RequestOption) (res *EnableFuturesToSubAccountResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/subAccount/margin",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", subAccountId)
	r.setParam("margin", true)

	data, err := c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EnableFuturesToSubAccountResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateApiKeyForSub make request
func (c *Client) CreateApiKeyForSub(ctx context.Context, req CreateApiKeyRequest, opts ...RequestOption) (res *CreateApiKeyResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/subAccountApi",
		secType:  secTypeSigned,
	}
	if req.SubAccountID != "" {
		r.setParam("subAccountId", req.SubAccountID)
	} else {
		return nil, common.CreateErrorMandatoryField("subAccountId")
	}
	r.setParam("canTrade", req.CanTrade)
	r.setParam("marginTrade", req.MarginTrade)
	r.setParam("futuresTrade", req.FuturesTrade)

	data, err := c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateApiKeyResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteSubApiKey deletes sub account api key
func (c *Client) DeleteSubApiKey(ctx context.Context, req DeleteSubApiKeyRequest, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/broker/subAccountApi",
		secType:  secTypeSigned,
	}
	if req.SubAccountID != "" {
		r.setParam("subAccountId", req.SubAccountID)
	} else {
		return common.CreateErrorMandatoryField("subAccountId")
	}
	if req.ApiKey == "" {
		return common.CreateErrorMandatoryField("subAccountApiKey")
	}
	r.setParam("subAccountApiKey", req.ApiKey)

	_, err := c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ChangeSubAccountApiPermission make request
func (c *Client) ChangeSubAccountApiPermission(ctx context.Context, req ChangeApiPermissionRequest, opts ...RequestOption) (res *ChangeApiPermissionResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/subAccountApi/permission",
		secType:  secTypeSigned,
	}

	if req.SubAccountID != "" {
		r.setParam("subAccountId", req.SubAccountID)
	} else {
		return nil, common.CreateErrorMandatoryField("subAccountId")
	}

	if req.SubAccountApiKey == "" {
		return nil, common.CreateErrorMandatoryField("subAccountApiKey")
	}
	r.setParam("subAccountApiKey", req.SubAccountApiKey)

	r.setParam("canTrade", req.CatTrade)
	r.setParam("marginTrade", req.MarginTrade)
	r.setParam("futuresTrade", req.FuturesTrade)

	data, err := c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ChangeApiPermissionResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ChangeSubAccountCommission make request
func (c *Client) ChangeSubAccountCommission(ctx context.Context, req ChangeSubAccountCommissionRequest, opts ...RequestOption) (res *ChangeSubAccountCommissionResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/subAccountApi/commission",
		secType:  secTypeSigned,
	}

	if req.SubAccountID != "" {
		r.setParam("subAccountId", req.SubAccountID)
	} else {
		return nil, common.CreateErrorMandatoryField("subAccountId")
	}
	r.setParam("makerCommission", req.MakerCommission)
	r.setParam("takerCommission", req.TakerCommission)
	r.setParam("marginMakerCommission", req.MarginMakerCommission)
	r.setParam("marginTakerCommission", req.MarginTakerCommission)

	data, err := c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ChangeSubAccountCommissionResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubAccountTransfer makes request
func (c *Client) SubAccountTransfer(ctx context.Context, req SubAccountTransferRequest, opts ...RequestOption) (res *SubAccountTransferResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/transfer",
		secType:  secTypeSigned,
	}

	if req.FromID == "" && req.ToID == "" {
		return nil, common.CreateErrorMandatoryField("fromID or toID")
	}
	if req.FromID != "" && req.ToID == "" {
		r.setParam("fromId", req.FromID)
	} else if req.ToID != "" && req.FromID == "" {
		r.setParam("toId", req.ToID)
	} else {
		r.setParam("fromId", req.FromID)
		r.setParam("toId", req.ToID)
	}
	if req.Asset == "" {
		return nil, common.CreateErrorMandatoryField("asset")
	} else {
		r.setParam("asset", req.Asset)
	}

	if req.ClientTransferID != "" {
		r.setParam("clientTranId", req.ClientTransferID)
	}

	if req.Amount > 0 {
		r.setParam("amount", req.Amount)
	} else {
		return nil, common.CreateErrorMandatoryField("amount")
	}

	data, err := c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(SubAccountTransferResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// TransferHistory makes request
func (c *Client) TransferHistory(ctx context.Context, req SubAccountTransferHistoryRequest, opts ...RequestOption) (transfers []*Transfer, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/transfer",
		secType:  secTypeSigned,
	}
	if req.FromID != "" {
		r.setParam("fromId", req.FromID)
	}
	if req.ToID != "" {
		r.setParam("toId", req.ToID)
	}
	if req.ClientTransferID != "" {
		r.setParam("clientTranId", req.ClientTransferID)
	}
	if req.StartTime > 0 {
		r.setParam("startTime", req.StartTime)
	}
	if req.EndTime > 0 {
		r.setParam("endTime", req.EndTime)
	}
	if req.Limit > 0 {
		r.setParam("limit", req.Limit)
	}
	if req.Page > 0 {
		r.setParam("page", req.Page)
	}

	data, err := c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(TransferHistoryResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	transfers = *res
	return transfers, nil
}

// callAPI makes API call
func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= 400 {
		apiErr := new(common.APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}
	return data, nil
}

// parseRequest parses given request
func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	if r.recvWindow > 0 {
		r.setParam(recvWindowKey, r.recvWindow)
	}
	if r.secType == secTypeSigned {
		r.setParam(timestampKey, currentTimestamp()-c.TimeOffset)
	}
	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}
	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}
	if r.secType == secTypeAPIKey || r.secType == secTypeSigned {
		header.Set("X-MBX-APIKEY", c.apiKey)
	}

	if r.secType == secTypeSigned {
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		mac := hmac.New(sha256.New, []byte(c.secretKey))
		_, err = mac.Write([]byte(raw))
		if err != nil {
			return err
		}
		v := url.Values{}
		v.Set(signatureKey, fmt.Sprintf("%x", mac.Sum(nil)))
		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s", fullURL, bodyString)

	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func currentTimestamp() int64 {
	return FormatTimestamp(time.Now())
}

// FormatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Binance.
func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// getAPIEndpoint return the base endpoint of the Rest API according the UseTestnet flag
func getAPIEndpoint() string {
	if UseTestnet {
		return baseAPITestnetURL
	}
	return baseAPIMainURL
}
