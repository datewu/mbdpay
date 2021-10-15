package mbdpay

import (
	"context"
	"errors"

	"github.com/datewu/fetch"
)

// Webhook for mbd wehhook
type Webhook struct {
	Type string   `json:"type"`
	Data dataWrap `json:"data"`
}

// Concrete to charge or complain hook
func (w Webhook) Concrete() (*ChargeHook, *ComplainHook) {
	switch w.Type {
	case "charge_succeeded":
		c := &ChargeHook{
			Description: w.Data.Description,
			OutTradeNo:  w.Data.OutTradeNo,
			Amount:      w.Data.Amount,
			OpenID:      w.Data.OpenID,
			ChargeID:    w.Data.ChargeID,
			Payway:      w.Data.Payway,
		}
		return c, nil
	case "complaint":
		c := &ComplainHook{
			OutTradeNo: w.Data.OutTradeNo,
			Detail:     w.Data.Detail,
			Amount:     w.Data.Amount,
			Phone:      w.Data.Phone,
		}
		return nil, c
	}
	return nil, nil
}

type dataWrap struct {
	Description string `json:"description"`
	OutTradeNo  string `json:"out_trade_no"`
	Amount      int    `json:"amount"`
	OpenID      string `json:"openid,omitempty"`
	ChargeID    string `json:"charge_id"`
	Payway      int    `json:"payway"`

	//OutTradeNo string `json:"out_trade_no"`
	Detail string `json:"complaint_detail"`
	//Amount     int `json:"amount"`
	Phone string `json:"payer_phone"`
}

// ChargeHook a charge type webhook handle type
type ChargeHook struct {
	Description string `json:"description"`
	OutTradeNo  string `json:"out_trade_no"`
	Amount      int    `json:"amount"`
	OpenID      string `json:"openid,omitempty"`
	ChargeID    string `json:"charge_id"`
	Payway      int    `json:"payway"`
}

// ComplainHook a complaint type webhook handle type
type ComplainHook struct {
	OutTradeNo string `json:"out_trade_no"`
	Detail     string `json:"complaint_detail"`
	Amount     int    `json:"amount"`
	Phone      string `json:"payer_phone"`
}

// RefundReq for mbd refund request
type RefundReq struct {
	OrderID string `json:"order_id"`
	AppID   string `json:"app_id"`
	Sign    string `json:"sign"`
}

func (r RefundReq) toParams() map[string]string {
	p := map[string]string{
		"order_id": r.OrderID,
		"app_id":   r.AppID,
	}
	return p
}

// RefundRes a refund response container
type RefundRes struct {
	// according to docs the code field should be int
	//	Code  int    `json:"code"`
	Code  string `json:"code"`
	Info  string `json:"info"`
	Error string `json:"error,omitempty"`
}

// Refund make refund api call
func (c Client) Refund(req *RefundReq) (*RefundRes, error) {
	const path = "/release/main/refund"
	req.AppID = c.id
	hashString := c.sign(req.toParams())
	req.Sign = hashString
	res := new(RefundRes)
	cli := fetch.DefaultClient(context.Background())
	err := cli.Post(apiAddress+path, req, res)
	if err != nil {
		return nil, err
	}
	if res.Error != "" {
		return nil, errors.New(res.Error)
	}
	return res, nil
}

// SearchReq for search request
type SearchReq struct {
	OutTradeNo string `json:"out_trade_no"`
	AppID      string `json:"app_id"`
	Sign       string `json:"sign"`
}

func (r SearchReq) toParams() map[string]string {
	p := map[string]string{
		"out_trade_no": r.OutTradeNo,
		"app_id":       r.AppID,
	}
	return p
}

// SearchRes a search response container
type SearchRes struct {
	OrderID      string `json:"order_id"`
	ChargeID     string `json:"charge_id"`
	Description  string `json:"description"`
	ShareID      string `json:"share_id"`
	ShareState   string `json:"share_state"`
	Amount       string `json:"amount"`
	State        string `json:"state"`
	CreateTime   string `json:"create_time"`
	Payway       string `json:"payway"`
	RefundState  string `json:"refund_state"`
	RefundAmount string `json:"refund_amount"`
	Error        string `json:"error,omitempty"`
}

// Search make search api call
func (c Client) Search(req *SearchReq) (*SearchRes, error) {
	const path = "/release/main/search_order"
	req.AppID = c.id
	hashString := c.sign(req.toParams())
	req.Sign = hashString
	res := new(SearchRes)
	cli := fetch.DefaultClient(context.Background())
	err := cli.Post(apiAddress+path, req, res)
	if err != nil {
		return nil, err
	}
	if res.Error != "" {
		return nil, errors.New(res.Error)
	}
	return res, nil
}
