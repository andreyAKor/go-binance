package binance

import (
	"context"
	"encoding/json"
)

// CreateWithdrawService create withdraw
type CreateWithdrawService struct {
	c       *Client
	asset   string
	address string
	amount  string
	name    *string
}

// Asset set asset
func (s *CreateWithdrawService) Asset(asset string) *CreateWithdrawService {
	s.asset = asset
	return s
}

// Address set address
func (s *CreateWithdrawService) Address(address string) *CreateWithdrawService {
	s.address = address
	return s
}

// Amount set amount
func (s *CreateWithdrawService) Amount(amount string) *CreateWithdrawService {
	s.amount = amount
	return s
}

// Name set name
func (s *CreateWithdrawService) Name(name string) *CreateWithdrawService {
	s.name = &name
	return s
}

// Do send request
func (s *CreateWithdrawService) Do(ctx context.Context) (err error) {
	r := &request{
		method:   "POST",
		endpoint: "/wapi/v3/withdraw.html",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	r.setParam("address", s.address)
	r.setParam("amount", s.amount)
	if s.name != nil {
		r.setParam("name", *s.name)
	}
	_, err = s.c.callAPI(ctx, r)
	return err
}

// ListWithdrawsService list withdraws
type ListWithdrawsService struct {
	c         *Client
	asset     *string
	status    *int
	startTime *int64
	endTime   *int64
}

// Asset set asset
func (s *ListWithdrawsService) Asset(asset string) *ListWithdrawsService {
	s.asset = &asset
	return s
}

// Status set status
func (s *ListWithdrawsService) Status(status int) *ListWithdrawsService {
	s.status = &status
	return s
}

// StartTime set startTime
func (s *ListWithdrawsService) StartTime(startTime int64) *ListWithdrawsService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListWithdrawsService) EndTime(endTime int64) *ListWithdrawsService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *ListWithdrawsService) Do(ctx context.Context) (withdraws []*Withdraw, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/wapi/v1/getWithdrawHistory.html",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
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
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res := new(WithdrawHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res.Withdraws, nil
}

// WithdrawHistoryResponse define withdraw history response
type WithdrawHistoryResponse struct {
	Withdraws []*Withdraw `json:"withdrawList"`
	Success   bool        `json:"success"`
}

// Withdraw define withdraw info
type Withdraw struct {
	Amount    float64 `json:"amount"`
	Address   string  `json:"address"`
	Asset     string  `json:"asset"`
	TxID      string  `json:"txId"`
	ApplyTime int64   `json:"applyTime"`
	Status    int     `json:"status"`
}

// GetAssetDetailService get asset detail
type GetAssetDetailService struct {
	c *Client
}

// Do send request
func (s *GetAssetDetailService) Do(ctx context.Context, opts ...RequestOption) (*AssetDetailResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/assetDetail.html",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(AssetDetailResponse)

	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// AssetDetail asset detail
type AssetDetailResponse struct {
	Success       bool                   `json:"success"`
	AssetsDetails map[string]AssetDetail `json:"assetDetail"`
}

// Full asset detail
type AssetDetail struct {
	MinWithdrawAmount float64 `json:"minWithdrawAmount"` // min withdraw amount
	DepositStatus     bool    `json:"depositStatus"`     // deposit status
	WithdrawFee       float64 `json:"withdrawFee"`       // withdraw fee
	WithdrawStatus    bool    `json:"withdrawStatus"`    // withdraw status
	DepositTip        string  `json:"depositTip"`        // reason
}
