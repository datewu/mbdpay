package mbdpay

import (
	"errors"
	"strconv"
)

// WxJSReq for weixin jsapi request
type WxJSReq struct {
	OpenID      string `json:"openid"`
	AppID       string `json:"app_id"`
	Description string `json:"description"`
	AmountTotal int    `json:"amount_total"`
	OutTradeNo  string `json:"out_trade_no"`
	CallbackURL string `json:"callback_url"`
	Sign        string `json:"sign"`
}

func (r WxJSReq) toParams() map[string]string {
	p := map[string]string{
		"openid":       r.OpenID,
		"app_id":       r.AppID,
		"description":  r.Description,
		"amount_total": strconv.Itoa(r.AmountTotal),
		"callback_url": r.CallbackURL,
	}
	if r.OutTradeNo != "" {
		p["out_trade_no"] = r.OutTradeNo
	}
	return p
}

// WxJSRes a weixin jsapi response container
type WxJSRes struct {
	AppID     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
	Error     string `json:"error,omitempty"`
}

// GetWxOpenID must be implemented on wexin builtin browser
func (c Client) GetWxOpenID(url string) (string, error) {
	// must open in weixin builtin browser
	// https://mbd.pub/openid?app_id=1234567890&target_url=http://www.example.com/abc?uid=32
	msg := "cannot implement on server,try using client cookie"
	return "", errors.New(msg)
}

// WxJS make weixin jsapi call
func (c Client) WxJS(req *WxJSReq) (*WxJSRes, error) {
	const path = "/release/wx/prepay"
	req.AppID = c.id
	hashString := c.sign(req.toParams())
	req.Sign = hashString
	res := new(WxJSRes)
	err := postJSON(apiAddress+path, req, res)
	if err != nil {
		return nil, err
	}
	if res.Error != "" {
		return nil, errors.New(res.Error)
	}
	return res, nil
}

// WxH5Req for weixin H5 request
type WxH5Req struct {
	Channel     string `json:"channel"`
	AppID       string `json:"app_id"`
	Description string `json:"description"`
	AmountTotal int    `json:"amount_total"`
	OutTradeNo  string `json:"out_trade_no"`
	Sign        string `json:"sign"`
}

func (r WxH5Req) toParams() map[string]string {
	p := map[string]string{
		"channel":      r.Channel,
		"app_id":       r.AppID,
		"description":  r.Description,
		"amount_total": strconv.Itoa(r.AmountTotal),
	}
	if r.OutTradeNo != "" {
		p["out_trade_no"] = r.OutTradeNo
	}
	return p
}

// WxH5Res a weixin H5 response container
type WxH5Res struct {
	URL   string `json:"h5_url"`
	Error string `json:"error,omitempty"`
}

// WxH5 make weixin H5 api call
func (c Client) WxH5(req *WxH5Req) (*WxH5Res, error) {
	const path = "/release/wx/prepay"
	req.Channel = "h5"
	req.AppID = c.id
	hashString := c.sign(req.toParams())
	req.Sign = hashString
	res := new(WxH5Res)
	err := postJSON(apiAddress+path, req, res)
	if err != nil {
		return nil, err
	}
	if res.Error != "" {
		return nil, errors.New(res.Error)
	}
	return res, nil
}
