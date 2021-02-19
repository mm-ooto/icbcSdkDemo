package example

import (
	"fmt"
	"github.com/mm-ooto/icbcSdkDemo/lib"
)

//线上POS聚合消费下单接口（无界面）
func CardbusinessAggregatepayB2cOnlineConsumepurchase() {
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
		"mer_id":           "*******",                //商户编号
		"out_trade_no":     "xxx09",                  //商户订单号，只能是数字、大小写字母，且在同一个商户号下唯一
		"pay_mode":         "9",                      //支付方式，9-微信；10-支付宝；13-云闪付
		"access_type":      "9",                      //收单接入方式，5-APP，7-微信公众号，8-支付宝生活号，9-微信小程序;
		"mer_prtcl_no":     "*******",                //收单产品协议编号
		"orig_date_time":   "2021-01-29T10:13:00",    //交易日期时间，格式为yyyy-MM-dd'T'HH:mm:ss
		"decive_info":      "xxxxx09090erwsfws",      //设备号？
		"body":             "body",                   //品描述，商品描述交易字段格式根据不同的应用场景按照以下格式：1）PC网站：传入浏览器打开的网站主页title名-实际商品名称 ；2）公众号：传入公众号名称-实际商品名称；3）H5：传入浏览器打开的移动网页的主页title名-实际商品名称；4）线下门店：门店品牌名-城市分店名-实际商品名称；5）APP：传入应用市场上的APP名字-实际商品名称
		"fee_type":         lib.FEE_TYPE_CNY,         //交易币种，目前工行只支持使用人民币（001）支付
		"spbill_create_ip": "*******",                //用户端IP
		"total_fee":        "100",                    //订单金额，单位为分
		"mer_url":          "http://doc.golang.ltd/", //异步通知商户URL，端口必须为443或80
		"shop_appid":       "*******",                //商户在微信开放平台注册的APPID，支付方式为微信时不能为空,如：微信appID,微信小程序appID
		"icbc_appid":       "*******",                //商户在工行API平台的APPID
		"open_id":          "*******",                //第三方用户标识，商户在微信公众号内或微信小程序内接入时必送，即access_type为7或9时，上送用户在商户APPID下的唯一标识；商户通过支付宝生活号接入时不送
		"mer_acct":         "",                       //商户账号，商户入账账号，只能交易时指定
		"expire_time":      "120",                    //订单失效时间，单位为秒，建议大于60秒
		"attach":           "",                       //附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
		"order_apd_inf":    "",                       //订单附加信息
		"detail":           "",                       //商品详细描述，对于使用单品优惠的商户，该字段必须按照规范上传。微信与支付宝的规范不同，请根据支付方式对应相应的规范上送，详细内容参考文末说明
		"notify_type":      "HS",                     //通知类型，表示在交易处理完成后把交易结果通知商户的处理模式。 取值“HS”：在交易完成后将通知信息，主动发送给商户，发送地址为mer_url指定地址； 取值“AG”：在交易完成后不通知商户
		"result_type":      "0",                      //结果发送类型，通知方式为HS时有效。取值“0”：无论支付成功或者失败，银行都向商户发送交易通知信息；取值“1”，银行只向商户发送交易成功的通知信息。默认是"0"
		"pay_limit":        "",                       //支付方式限定，上送”no_credit“表示不支持信用卡支付；上送“no_balance”表示仅支持银行卡支付；不上送或上送空表示无限制
		"return_url":       "",                       //支付成功回显页面，支付成功后，跳转至该页面显示。当access_type=5且pay_mode=10才有效
		"quit_url":         "",                       //用户付款中途退出返回商户网站的地址（仅对浏览器内支付时有效）当access_type=5且pay_mode=10才有效
	}

	gjsonBody, err := cli.Execute(bizContentMap, lib.Urlb2cOnlineConsumepurchase, lib.METHOD_TYPE_POST)
	if err != nil {
		fmt.Printf("cli execute error:%s\n", err.Error())
		return
	}
	fmt.Printf("respose result:%s", gjsonBody.Raw)

}
