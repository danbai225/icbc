package icbc

import (
	"os"
	"strconv"
	"testing"
	"time"
)

var (
	client *Client
)

func init() {
	var err error
	client, err = newTestClient()
	if err != nil {
		panic(err)
	}
}

func newTestClient() (*Client, error) {
	options := Options{
		Host:             "https://apipcs3.dccnet.com.cn",
		AppID:            os.Getenv("APP_ID"),
		AppPrivateKey:    os.Getenv("APP_PRIVATE_KEY"),
		GatewayPublicKey: os.Getenv("GATEWAY_PUBLIC_KEY"),
	}
	return NewClient(options)
}

func TestClient_Execute(t *testing.T) {
	msgID := strconv.FormatInt(time.Now().UnixNano(), 10)
	reqBiz := QrcodeQueryRequestV2Biz{
		MerID:      os.Getenv("MER_ID"),
		OutTradeNo: os.Getenv("OUT_TRADE_NO"),
	}
	var respBiz QrcodeQueryResponseV2Biz
	err := client.Execute(msgID, reqBiz, &respBiz)
	if err != nil {
		t.Fatal(err)
	}
}
