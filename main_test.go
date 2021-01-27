package icbc_api_sdk

import (
	"github.com/mm-ooto/icbcSdkDemo/lib"
	"testing"
)

//h5在线支付
func TestH5Pay(t *testing.T) {
	baseParams := &lib.Base{
		AppId:          "*******",
		PrivateKey:     "*******",
		SignType:       "RSA2",
		Charset:        "UTF-8",
		Format:         "json",
		IcbcPublickKey: "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCwFgHD4kzEVPdOj03ctKM7KV+16bWZ5BMNgvEeuEQwfQYkRVwI9HFOGkwNTMn5hiJXHnlXYCX+zp5r6R52MY0O7BsTCLT7aHaxsANsvI9ABGx3OaTVlPB59M6GPbJh0uXvio0m1r/lTW3Z60RU6Q3oid/rNhP3CiNgg0W6O3AGqwIDAQAB",
		EncryptKey:     "*******",
		EncryptType:    "",
		Ca:             "",
		Password:       "",
		IcbcHost:       "https://apipcs3.dccnet.com.cn",
		IsNeedEncrypt:  false,
	}
	action := ""
	bizContentBytes := []byte{}
	address, err := new(lib.UiIcbcClient).Execute(baseParams, action, bizContentBytes)
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(address)
}
