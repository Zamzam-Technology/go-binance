package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListDepositsService fetches deposit history.
//
// See https://binance-docs.github.io/apidocs/spot/en/#deposit-history-user_data
type ListDepositsService struct {
	c         *Client
	coin      *string
	status    *int
	startTime *int64
	endTime   *int64
	offset    *int
	limit     *int
}

// Coin sets the coin parameter.
func (s *ListDepositsService) Coin(coin string) *ListDepositsService {
	s.coin = &coin
	return s
}

// Status sets the status parameter.
func (s *ListDepositsService) Status(status int) *ListDepositsService {
	s.status = &status
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListDepositsService) StartTime(startTime int64) *ListDepositsService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListDepositsService) EndTime(endTime int64) *ListDepositsService {
	s.endTime = &endTime
	return s
}

// Offset sets the offset parameter
func (s *ListDepositsService) Offset(offset int) *ListDepositsService {
	s.offset = &offset
	return s
}

// Limit sets the limit parameter
func (s *ListDepositsService) Limit(limit int) *ListDepositsService {
	s.limit = &limit
	return s
}

// Do sends the request.
func (s *ListDepositsService) Do(ctx context.Context) (deposits []*Deposit, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/deposit/hisrec",
		secType:  secTypeSigned,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
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
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res := new(DepositHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res.Deposits, nil
}

// DepositHistoryResponse represents a response from ListDepositsService.
type DepositHistoryResponse struct {
	Success  bool       `json:"success"`
	Deposits []*Deposit `json:"depositList"`
}

// Deposit represents a single deposit entry.
type Deposit struct {
	InsertTime int64   `json:"insertTime"`
	Amount     float64 `json:"amount"`
	Asset      string  `json:"asset"`
	Address    string  `json:"address"`
	AddressTag string  `json:"addressTag"`
	TxID       string  `json:"txId"`
	Status     int     `json:"status"`
}

// GetDepositsAddressService retrieves the details of a deposit address.
//
// See https://binance-docs.github.io/apidocs/spot/en/#deposit-address-supporting-network-user_data
type GetDepositsAddressService struct {
	c       *Client
	asset   string
	network string
}

// Asset sets the asset parameter (MANDATORY).
func (s *GetDepositsAddressService) Asset(v string) *GetDepositsAddressService {
	s.asset = v
	return s
}

// Network sets the status parameter.
func (s *GetDepositsAddressService) Network(net string) *GetDepositsAddressService {
	s.network = net
	return s
}

// Do sends the request.
func (s *GetDepositsAddressService) Do(ctx context.Context) (*GetDepositAddressResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/deposit/address",
		secType:  secTypeSigned,
	}
	r.setParam("coin", s.asset)
	if s.network != "" {
		r.setParam("network", s.network)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &GetDepositAddressResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetDepositAddressResponse represents a response from GetDepositsAddressService.
type GetDepositAddressResponse struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}
