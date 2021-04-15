package mbdpay

import (
	"errors"
	"strconv"
)

// AliReq for alipay request
type AliReq struct {
	URL         string `json:"url"`
	AppID       string `json:"app_id"`
	Description string `json:"description"`
	AmountTotal int    `json:"amount_total"`
	OutTradeNo  string `json:"out_trade_no"`
	CallbackURL string `json:"callback_url"`
	Sign        string `json:"sign"`
}

func (r AliReq) toParams() map[string]string {
	p := map[string]string{
		"url":          r.URL,
		"app_id":       r.AppID,
		"description":  r.Description,
		"amount_total": strconv.Itoa(r.AmountTotal),
	}
	if r.OutTradeNo != "" {
		p["out_trade_no"] = r.OutTradeNo
	}
	if r.CallbackURL != "" {
		p["callback_url"] = r.CallbackURL
	}
	return p
}

// AliRes a alipay response container
type AliRes struct {
	Body  string `json:"body"`
	Error string `json:"error,omitempty"`
}

// AliPay make alipay api call
func (c Client) AliPay(req *AliReq) (*AliRes, error) {
	const path = "/release/alipay/pay"
	req.AppID = c.id
	hashString := c.sign(req.toParams())
	req.Sign = hashString
	res := new(AliRes)
	err := postJSON(apiAddress+path, req, res)
	if err != nil {
		return nil, err
	}
	if res.Error != "" {
		return nil, errors.New(res.Error)
	}
	return res, nil
}
