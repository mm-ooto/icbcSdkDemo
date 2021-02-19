package example

import (
	"fmt"
	"github.com/mm-ooto/base/common/config"
	"github.com/mm-ooto/icbcSdkDemo/lib"
)

//生成二维码
func APICardbusinessQrcodeQrgenerate() {
	baseParams := &lib.Base{
		AppId:          config.AppConfig.String("appId"),
		PrivateKey:     config.AppConfig.String("myPrivateKey"),
		SignType:       "RSA2",
		Charset:        "UTF-8",
		Format:         "json",
		IcbcPublickKey: config.AppConfig.String("apiGwPublicKey"),
		EncryptKey:     "*******",
		EncryptType:    "",
		Ca:             "",
		Password:       "",
		IcbcHost:       config.AppConfig.String("host"),
		IsNeedEncrypt:  false,
	}
	cli, err := lib.NewIcbcClient(baseParams)
	if err != nil {
		fmt.Printf("newIcbcClient error:%s\n", err.Error())
		return
	}
	bizContentMap := map[string]string{
		"mer_id":            config.AppConfig.String("offlineMerId"), //商户线下档案编号,特约商户12位，特约部门15位
		"out_trade_no":      "O1111",                                 //商户系统订单号
		"order_amt":         "1200",                                  //订单总金额，单位分
		"trade_date":        "*******",                               //商户订单生成日期，格式：yyyyMMdd
		"trade_time":        "*******",                               //商户订单生成时间，格式：HHmmss
		"attach":            "",                                      //商户附加数据，最多21个汉字字符
		"pay_expire":        "1200",                                  //二维码有效期，单位：秒，必须小于24小时1200
		"notify_url":        config.AppConfig.String("notifyUrl"),    //商户接收支付成功通知消息URL，当notify_flag为1时必输
		"tporder_create_ip": "127.0.0.1",                               //商户订单生成的机器IP
		"sp_flag":           "0",                                     //扫码后是否需要跳转分行，0：否，1：是，默认值0
		"notify_flag":       "1",                                     //商户是否开启通知接口，0-否；1-是，默认值0
	}

	gjsonBody, err := cli.Execute(bizContentMap, lib.UrlQrcodeGenerate, lib.METHOD_TYPE_POST)
	if err != nil {
		fmt.Printf("cli execute error:%s\n", err.Error())
		return
	}
	fmt.Printf("respose result:%s", gjsonBody.Raw)
}
