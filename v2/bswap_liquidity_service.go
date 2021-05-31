package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

type OperationType string

var (
	LiqOperationAdd    OperationType = "ADD"
	LiqOperationRemove OperationType = "REMOVE"
)

// LiquidityInformationService gets liquidity information
//
// See https://binance-docs.github.io/apidocs/spot/en/#get-liquidity-information-of-a-pool-user_data
type LiquidityInformationService struct {
	c      *Client
	poolId *int
}

// PoolID set the poolId parameter
func (s *LiquidityInformationService) PoolID(id int) *LiquidityInformationService {
	s.poolId = &id
	return s
}

// Do makes request
func (s *LiquidityInformationService) Do(ctx context.Context) (ls []*PoolLiquidity, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/liquidity",
		secType:  secTypeSigned,
	}
	if s.poolId != nil {
		r.setParam("poolId", *s.poolId)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(PoolLiquidityResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	ls = *res
	return ls, nil
}

// AddLiquidityService adds liquidity to a pool
//
// See https://binance-docs.github.io/apidocs/spot/en/#add-liquidity-trade
type AddLiquidityService struct {
	c        *Client
	poolId   *int
	asset    *string
	quantity *float64
}

// PoolID set the poolId parameter
func (s *AddLiquidityService) PoolID(id int) *AddLiquidityService {
	s.poolId = &id
	return s
}

// Asset set the asset parameter
func (s *AddLiquidityService) Asset(asset string) *AddLiquidityService {
	s.asset = &asset
	return s
}

// Quantity set the quantity parameter
func (s *AddLiquidityService) Quantity(qty float64) *AddLiquidityService {
	s.quantity = &qty
	return s
}

// Do makes request
func (s *AddLiquidityService) Do(ctx context.Context) (res *LiqIdResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/bswap/liquidityAdd",
		secType:  secTypeSigned,
	}

	if s.poolId != nil {
		r.setFormParam("poolId", *s.poolId)
	}
	if s.quantity != nil {
		r.setFormParam("quantity", *s.quantity)
	}
	if s.asset != nil {
		r.setFormParam("asset", *s.asset)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(LiqIdResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// RemoveLiquidityService remove liquidity from a pool
//
// See https://binance-docs.github.io/apidocs/spot/en/#remove-liquidity-trade
type RemoveLiquidityService struct {
	c           *Client
	poolId      *int
	asset       []string
	typ         *string
	shareAmount *float64
}

// PoolID set the poolId parameter
func (s *RemoveLiquidityService) PoolID(id int) *RemoveLiquidityService {
	s.poolId = &id
	return s
}

// Asset set the asset parameter
func (s *RemoveLiquidityService) Asset(asset []string) *RemoveLiquidityService {
	s.asset = asset
	return s
}

// Type set the type parameter
func (s *RemoveLiquidityService) Type(t string) *RemoveLiquidityService {
	s.typ = &t
	return s
}

// ShareAmount set the shareAmount parameter
func (s *RemoveLiquidityService) ShareAmount(a float64) *RemoveLiquidityService {
	s.shareAmount = &a
	return s
}

// Do makes request
func (s *RemoveLiquidityService) Do(ctx context.Context) (res *LiqIdResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/bswap/liquidityRemove",
		secType:  secTypeSigned,
	}

	if s.typ != nil {
		r.setFormParam("type", *s.typ)
	}
	for _, a := range s.asset {
		r.addParam("asset", a)
	}
	if s.poolId != nil {
		r.setFormParam("poolId", *s.poolId)
	}
	if s.shareAmount != nil {
		r.setFormParam("shareAmount", *s.shareAmount)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res = new(LiqIdResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// LiquidityOperationService get liquidity operation (add/remove) records
//
// See https://binance-docs.github.io/apidocs/spot/en/#get-liquidity-operation-record-user_data
type LiquidityOperationService struct {
	c           *Client
	operationId *int
	poolId      *int
	operation   *OperationType
	startTime   *int64
	endTime     *int64
	limit       *int
}

// OperationID set hte operationId parameter
func (s *LiquidityOperationService) OperationID(id int) *LiquidityOperationService {
	s.operationId = &id
	return s
}

// PoolID set the poolId parameter
func (s *LiquidityOperationService) PoolID(id int) *LiquidityOperationService {
	s.poolId = &id
	return s
}

// Operation set the operation parameter
func (s *LiquidityOperationService) Operation(op OperationType) *LiquidityOperationService {
	s.operation = &op
	return s
}

// StartTime sets the startTime parameter
func (s *LiquidityOperationService) StartTime(id int64) *LiquidityOperationService {
	s.startTime = &id
	return s
}

// EndTime sets the endTime parameter
func (s *LiquidityOperationService) EndTime(id int64) *LiquidityOperationService {
	s.endTime = &id
	return s
}

// Limit sets the limit parameter
func (s *LiquidityOperationService) Limit(limit int) *LiquidityOperationService {
	s.limit = &limit
	return s
}

// LiqIdResponse represent response
type LiqIdResponse struct {
	OperationID int64 `json:"operationId"`
}

type PoolLiquidityResponse []*PoolLiquidity

// PoolLiquidity represent pool liquidity
type PoolLiquidity struct {
	PoolID     int            `json:"poolId"`
	PoolName   string         `json:"poolName"`
	UpdateTime int64          `json:"updateTime"`
	Liquidity  Liquidity      `json:"liquidity"`
	Share      ShareLiquidity `json:"share"`
}

// Liquidity represents liquidity
type Liquidity map[string]float64

// ShareLiquidity represents share
type ShareLiquidity struct {
	ShareAmount     int
	SharePercentage float64
	Asset           map[string]float64
}
