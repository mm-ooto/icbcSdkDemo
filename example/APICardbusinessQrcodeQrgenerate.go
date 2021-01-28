
package example

import (
	"fmt"
	"github.com/mm-ooto/icbcSdkDemo/lib"
)

//生成二维码
func APICardbusinessQrcodeQrgenerate()  {
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
	cli, err := lib.NewIcbcClient(baseParams)
	if err != nil {
		fmt.Printf("newIcbcClient error:%s\n", err.Error())
		return
	}
	bizContentMap := map[string]string{
		"mer_id":            "*******",                                      //商户线下档案编号,特约商户12位，特约部门15位
		"out_trade_no":      "O1111",                              //商户系统订单号
		"order_amt":         "1200", //订单总金额，单位分
		"trade_date":        "*******",             //商户订单生成日期，格式：yyyyMMdd
		"trade_time":        "*******",             //商户订单生成时间，格式：HHmmss
		"attach":            "",                                       //商户附加数据，最多21个汉字字符
		"pay_expire":        "1200",                                 //二维码有效期，单位：秒，必须小于24小时1200
		"notify_url":        "*******",                                  //商户接收支付成功通知消息URL，当notify_flag为1时必输
		"tporder_create_ip": "*******",                             //商户订单生成的机器IP
		"sp_flag":           "0",                                          //扫码后是否需要跳转分行，0：否，1：是，默认值0
		"notify_flag":       "1",                                          //商户是否开启通知接口，0-否；1-是，默认值0
	}
	action := "/api/cardbusiness/qrcode/qrgenerate/V1"
	gjsonBody, err := cli.Execute(bizContentMap, action, lib.METHOD_TYPE_POST)
	if err != nil {
		fmt.Printf("cli execute error:%s\n", err.Error())
		return
	}
	fmt.Printf("respose result:%s", gjsonBody.Raw)
}
