package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListSwapPoolsService fetch pools
//
// See https://binance-docs.github.io/apidocs/spot/en/#list-all-swap-pools-market_data
type ListSwapPoolsService struct {
	c *Client
}

// Do makes request
func (s *ListSwapPoolsService) Do(ctx context.Context) (pools []*Pool, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/bswap/pools",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res := new(ListPoolsResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	pools = *res

	return pools, nil
}

// RequestQuoteService request a quote for swap quote asset (selling asset) for base asset (buying asset), essentially price/exchange rates.
//
// See https://binance-docs.github.io/apidocs/spot/en/#request-quote-user_data
type RequestQuoteService struct {
	c          *Client
	quoteAsset *string
	baseAsset  *string
	quoteQty   *float64
}

// QuoteAsset sets the quoteAsset parameter
func (s *RequestQuoteService) QuoteAsset(asset string) *RequestQuoteService {
	s.quoteAsset = &asset
	return s
}

// BaseAsset sets the baseAsset parameter
func (s *RequestQuoteService) BaseAsset(asset string) *RequestQuoteService {
	s.baseAsset = &asset
	return s
}

// Quantity sets the quoteQty parameter
func (s *RequestQuoteService) Quantity(qty float64) *RequestQuoteService {
	s.quoteQty = &qty
	return s
}

// Do makes request
func (s *RequestQuoteService) Do(ctx context.Context) (res *RequestQuoteResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/quote",
		secType:  secTypeSigned,
	}
	if s.baseAsset != nil {
		r.setParam("baseAsset", *s.baseAsset)
	}
	if s.quoteQty != nil {
		r.setParam("quoteQty", *s.quoteQty)
	}
	if s.quoteAsset != nil {
		r.setParam("quoteAsset", *s.quoteAsset)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(RequestQuoteResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MakeSwapService makes swap
//
// See https://binance-docs.github.io/apidocs/spot/en/#swap-trade
type MakeSwapService struct {
	c          *Client
	quoteAsset *string
	baseAsset  *string
	quoteQty   *float64
}

// QuoteAsset sets the quoteAsset parameter
func (s *MakeSwapService) QuoteAsset(asset string) *MakeSwapService {
	s.quoteAsset = &asset
	return s
}

// BaseAsset sets the baseAsset parameter
func (s *MakeSwapService) BaseAsset(asset string) *MakeSwapService {
	s.baseAsset = &asset
	return s
}

// Quantity sets the quoteQty parameter
func (s *MakeSwapService) Quantity(qty float64) *MakeSwapService {
	s.quoteQty = &qty
	return s
}

// Do makes request
func (s *MakeSwapService) Do(ctx context.Context) (res *MakeSwapResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/bswap/swap",
		secType:  secTypeSigned,
	}
	if s.baseAsset != nil {
		r.setFormParam("baseAsset", *s.baseAsset)
	}
	if s.quoteQty != nil {
		r.setFormParam("quoteQty", *s.quoteQty)
	}
	if s.quoteAsset != nil {
		r.setFormParam("quoteAsset", *s.quoteAsset)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(MakeSwapResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SwapHistoryService fetch swap history
//
// See https://binance-docs.github.io/apidocs/spot/en/#get-swap-history-user_data
type SwapHistoryService struct {
	c          *Client
	swapId     *int64
	startTime  *int64
	endTime    *int64
	status     *int
	quoteAsset *string
	baseAsset  *string
	limit      *int
}

// QuoteAsset sets the quoteAsset parameter
func (s *SwapHistoryService) QuoteAsset(asset string) *SwapHistoryService {
	s.quoteAsset = &asset
	return s
}

// BaseAsset sets the baseAsset parameter
func (s *SwapHistoryService) BaseAsset(asset string) *SwapHistoryService {
	s.baseAsset = &asset
	return s
}

// StartTime sets the startTime parameter
func (s *SwapHistoryService) StartTime(id int64) *SwapHistoryService {
	s.startTime = &id
	return s
}

// EndTime sets the endTime parameter
func (s *SwapHistoryService) EndTime(id int64) *SwapHistoryService {
	s.endTime = &id
	return s
}

// Status sets the status parameter
func (s *SwapHistoryService) Status(status int) *SwapHistoryService {
	s.status = &status
	return s
}

// Limit sets the limit parameter
func (s *SwapHistoryService) Limit(limit int) *SwapHistoryService {
	s.limit = &limit
	return s
}

// Do makes request
func (s *SwapHistoryService) Do(ctx context.Context) (swaps []*Swap, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/swap",
		secType:  secTypeSigned,
	}
	if s.baseAsset != nil {
		r.setParam("baseAsset", *s.baseAsset)
	}
	if s.quoteAsset != nil {
		r.setParam("quoteAsset", *s.quoteAsset)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.swapId != nil {
		r.setParam("status", *s.swapId)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(GetSwapResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	swaps = *res
	return swaps, nil
}

// GetSwapResponse represent history of swaps response
type GetSwapResponse []*Swap

// Swap represent swap history data
type Swap struct {
	QuoteAsset string  `json:"quoteAsset"`
	BaseAsset  string  `json:"baseAsset"`
	QuoteQty   float64 `json:"quoteQty"`
	BaseQty    float64 `json:"baseQty"`
	Price      float64 `json:"price"`
	Slippage   float64 `json:"slippage"`
	Fee        float64 `json:"fee"`
}

// MakeSwapResponse represent swap response
type MakeSwapResponse struct {
	SwapID int64 `json:"swapId"`
}

// RequestQuoteResponse represent quote request response
type RequestQuoteResponse struct {
	QuoteAsset string  `json:"quoteAsset"`
	BaseAsset  string  `json:"baseAsset"`
	QuoteQty   float64 `json:"quoteQty"`
	BaseQty    float64 `json:"baseQty"`
	Price      float64 `json:"price"`
	Slippage   float64 `json:"slippage"`
	Fee        float64 `json:"fee"`
}

// ListPoolsResponse represent pools response
type ListPoolsResponse []*Pool

// Pool represent a swap pool
type Pool struct {
	PoolID   int      `json:"poolId"`
	PoolName string   `json:"poolName"`
	Assets   []string `json:"assets"`
}
