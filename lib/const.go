package lib

const (
	//工行测试环境host
	ICBC_TEST_HOST="https://apipcs3.dccnet.com.cn"
)

//公共字段定义
const (
	FROM                 = "from"
	API                  = "api"
	APP_ID               = "app_id"
	FORMAT               = "format"
	CA                   = "ca"
	CA_Encrypt_Key       = "encrypt_key"
	CA_Password          = "password"
	TIMESTAMP            = "timestamp"
	SIGN_TYPE            = "sign_type"
	SIGN                 = "sign"
	CHARSET              = "charset"
	MSG_ID               = "msg_id"
	BIZ_CONTENT_KEY      = "biz_content"
	RESPONSE_BIZ_CONTENT = "response_biz_content"
	NOTIFY_URL           = "notify_url"
	RETURN_URL           = "return_url"
	ENCRYPT_TYPE         = "encrypt_type"
	RETURN_CODE          = "return_code"
	RETURN_MSG           = "return_msg"
)

const (
	//加密方式
	ENCRYPT_TYPE_AES = "AES"
	//签名类型
	SIGN_TYPE_RSA  = "RSA"
	SIGN_TYPE_RSA2 = "RSA2"
	SIGN_TYPE_CA   = "CA"
	//交易币种
	FEE_TYPE_CNY = "001"
	//字符集
	CHARSET_UTF8 = "UTF-8"
	CHARSET_GBK  = "GBK"
	//数据格式
	FORMAT_JSON = "json"
	//通知类型，表示在交易处理完成后把交易结果通知商户的处理模式。 取值“HS”：在交易完成后将通知信息，主动发送给商户，发送地址为mer_url指定地址； 取值“AG”：在交易完成后不通知商户
	NOTIFY_TYPE_HS = "HS"
	NOTIFY_TYPE_AG = "AG"
	//结果发送类型，通知方式为HS时有效。取值“0”：无论支付成功或者失败，银行都向商户发送交易通知信息；取值“1”，银行只向商户发送交易成功的通知信息。默认是"0"
	RESULT_TYPE_0 = "0"
	RESULT_TYPE_1 = "1"
	//请求方式
	METHOD_TYPE_POST = "POST"
	METHOD_TYPE_GET  = "GET"
)

const (
	//时间或日期格式
	Date_Format      = "20060102"
	Time_Format      = "150203"
	Date_Time_Format = "2006-01-02 15:02:03"
)

