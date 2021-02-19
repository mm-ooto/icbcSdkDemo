package example

import (
	"fmt"
	"github.com/mm-ooto/icbcSdkDemo/lib"
)

//线上POS支付宝H5消费下单接口
func CardbusinessZfbh5UiH5consumption() {
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
	bizContentMap := map[string]string{
		"mer_id":          "*******",                // 必填，商户线下档案编号,特约商户12位，特约部门15位
		"mer_prtcl_no":    "*******",                // 必填，收单产品协议编号
		"order_id":        "x900990dfs0",            // 必填，商户订单号，只能是数字、大小写字母，且在同一个商户号下唯一
		"order_date_time": "2021-02-04T12:00:00",    // 必填，交易日期时间，格式为yyyy-MM-dd'T'HH:mm:ss
		"amount":          "12",                     //单位分
		"cur_type":        lib.FEE_TYPE_CNY,         // 必填，交易币种，目前工行只支持使用人民币（001）支付
		"body":            "H5在线支付",                 // 必填，商品描述
		"notify_url":      "http://doc.golang.ltd/", // 必填，异步通知商户URL，端口必须为443或80
		"icbc_appid":      "*******",                // 必填，工行API平台的APPID
		"notify_type":     lib.NOTIFY_TYPE_HS,       // 必填，通知类型，表示在交易处理完成后把交易结果通知商户的处理模式。 取值“HS”：在交易完成后将通知信息，主动发送给商户，发送地址为mer_url指定地址； 取值“AG”：在交易完成后不通知商户
		"expireTime":      "1200",                   // 可选 订单失效时间，单位为秒，建议大于60秒//异步通知商户URL，端口必须为443或80
		"result_type":     lib.RESULT_TYPE_0,        // 可选 结果发送类型
		"pay_limit":       "",                       // 可选 支付方式限定
		"return_url":      "",                       // 可选 支付成功回显页面
		"quit_url":        "",                       // 可选 用户付款中途退出返回商户网站的地址（仅对浏览器内支付时有效
		"order_apd_inf":   "",                       // 可选 订单附加信息
		"detail":          "",                       // 可选 商品详细描述
		"attach":          "",                       // 可选 附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	}

	address, err := new(lib.UiIcbcClient).Execute(baseParams, lib.Urlzfbh5uih5Consumepurchase, bizContentMap)
	if err != nil {
		fmt.Printf("uiIcbcClient execute error:%s\n", err.Error())
		return
	}
	fmt.Printf("address:%s\n", address)
}
