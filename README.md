# mbdpay
### description
A client/sdk for [mbdpay api](https://doc.mbd.pub/)

## usage
import as a normal go package

```go
import pay "github.com/datewu/mbdpay"

```

### demo

```golang
package main

import (
	"fmt"
	"log"
	"os"

	pay "github.com/datewu/mbdpay"
)

func main() {
	demoApiID := os.Getenv("APP_ID")
	demoApiKEY := os.Getenv("APP_KEY")
	cli := pay.New(demoApiID, demoApiKEY)

	wxjsDemo(cli)
	wxH5Demo(cli)
	aliDemo(cli)
	refundDemo(cli)
	searchDemo(cli)
}

func wxjsDemo(cli *pay.Client) {
	const demoOpenid = "orZlx1uViEPNW5B8iTPwgaYLAw-w"
	req := &pay.WxJSReq{
		OpenID:      demoOpenid,
		Description: "测试jsapi",
		AmountTotal: 40, // 40 fen
		OutTradeNo:  "xdlke11244testodder_number_xx88",
		CallbackURL: "http://wutuofu.com",
	}

	res, err := cli.WxJS(req)
	if err != nil {
		log.Println("error:", err)
		return
	}
	fmt.Println("wxjs response:")
	fmt.Println(res)
}

func wxH5Demo(cli *pay.Client) {
	req := &pay.WxH5Req{
		Description: "test商品",
		AmountTotal: 45,
		OutTradeNo:  "h5_order_xxx_yyy",
	}
	res, err := cli.WxH5(req)
	if err != nil {
		log.Println("wxH5 error:", err)
		return
	}
	fmt.Println("wxH5 response:")
	fmt.Println(res)
}

func aliDemo(cli *pay.Client) {
	req := &pay.AliReq{
		URL:         "https://wutuofu.com",
		Description: "ali test商品",
		AmountTotal: 35,
		OutTradeNo:  "ali_order-32423-xxx",
		CallbackURL: "https://wutuofu.com/ali/redirect",
	}

	res, err := cli.AliPay(req)
	if err != nil {
		log.Println("ali error:", err)
		return
	}
	fmt.Println("ali response:")
	fmt.Println(res)
}

func refundDemo(cli *pay.Client) {
	req := &pay.RefundReq{
		OrderID: "xdlke11244testodder_number_xx88",
	}

	res, err := cli.Refund(req)
	if err != nil {
		log.Println("refund error:", err)
		return
	}
	fmt.Println("refund response:")
	fmt.Println(res)
}

func searchDemo(cli *pay.Client) {
	req := &pay.SearchReq{
		OutTradeNo: "xdlke11244testodder_number_xx88",
	}

	res, err := cli.Search(req)
	if err != nil {
		log.Println("search error:", err)
		return
	}
	fmt.Println("search order response:")
	fmt.Println(res)
}




```

```shell
export APP_ID=your_app_id
export APP_KEY=your_app_key
go run main.go
```
