package justap

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
)

func CreateChargeExample() {
	auth := context.WithValue(context.Background(), ContextAPIKey, APIKey{
		Key:    "Test_abcdefg",
		Prefix: "",
	})

	c := ALIPAY_QR_V1Channel
	cfg := NewConfiguration()
	err := cfg.SetMerchantRsaPrivateKeyByPath("./private_key.pem")
	if err != nil {
		return
	}

	client := NewAPIClient(cfg)
	r, response, err := client.TradeServiceApi.TradeServiceCharges(auth, V1CreateChargeRequest{
		AppId:           "app_z61w0gk524g2lyqv8oxp",
		Channel:         &c,
		Amount:          0.01,
		ClientIp:        "112.74.39.50",
		MerchantTradeId: "test_" + String(10),
		Subject:         "subject",
		Body:            "body",
	})

	fmt.Println(r, response, err)
	fmt.Printf("%+v\n", r)
}

func main() {
	CreateChargeExample()
}

func String(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
