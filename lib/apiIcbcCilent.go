package lib

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"gopkg.in/resty.v1"
	"sort"
	"strings"
	"time"
)

type Base struct {
	AppId          string
	PrivateKey     string
	SignType       string
	Charset        string
	Format         string
	IcbcPublickKey string
	EncryptKey     string
	EncryptType    string
	Ca             string
	Password       string
	IcbcHost       string
	IsNeedEncrypt  bool
}

type IcbcClient struct {
	Base
}

func NewIcbcClient(params *Base) (apiIcbcClient *IcbcClient, err error) {
	apiIcbcClient = &IcbcClient{
		Base{
			AppId:          params.AppId,
			PrivateKey:     getPemPrivate(params.PrivateKey),
			Format:         FORMAT_JSON,
			IcbcPublickKey: getPemPublic(params.IcbcPublickKey),
			Password:       params.Password,
			IcbcHost:       params.IcbcHost,
		},
	}
	if params.IsNeedEncrypt {
		apiIcbcClient.EncryptKey = params.EncryptKey
		apiIcbcClient.EncryptType = ENCRYPT_TYPE_AES //目前工行只支持AES加密
	}
	if params.SignType == "" {
		apiIcbcClient.SignType = SIGN_TYPE_RSA
	} else {
		apiIcbcClient.SignType = params.SignType
	}
	if params.Charset == "" {
		apiIcbcClient.Charset = CHARSET_UTF8
	} else {
		apiIcbcClient.Charset = params.Charset
	}
	if params.Ca != "" {
		//去除证书数据中的空格
		params.Ca = strings.ReplaceAll(params.Ca, " ", "")
		apiIcbcClient.Ca = params.Ca
	}
	return
}

//执行请求
func (i *IcbcClient) Execute(bizContentData interface{}, action, requestType string) (*gjson.Result, error) {
	commonData := map[string]string{
		APP_ID:    i.AppId,                             // 必须 APP的编号
		MSG_ID:    getMsgId(),                          // 必须 消息通讯唯一编号
		FORMAT:    FORMAT_JSON,                         // 可选 请求参数格式，仅支持json
		CHARSET:   i.Charset,                           // 可选 字符集，缺省为UTF-8
		SIGN_TYPE: i.SignType,                          // 可选 加密方式
		TIMESTAMP: time.Now().Format(Date_Time_Format), // 必须 交易发生时间戳
	}
	params, _, err := i.prepareParams(commonData, bizContentData, action)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s%s", i.IcbcHost, action)
	//发送请求
	resBody, err := i.request(params, url, requestType)
	if err != nil {
		return nil, err
	}
	//解析响应
	result := gjson.Parse(string(resBody))
	responseBizContentResult := result.Get(RESPONSE_BIZ_CONTENT)
	responseBizContent := responseBizContentResult.String()

	if i.IsNeedEncrypt { //如果是加密数据则对返回的数据先解密，后验签
		responseBizContent = AesCFBDecrypt(responseBizContent, i.EncryptKey)
	}
	//对请求结果验签
	signData := result.Get(SIGN).String()
	signType := SIGN_TYPE_RSA //工行验签暂时只支持RSA
	if err := IcbcVerifySignature(responseBizContent, signData, signType, i.IcbcPublickKey, i.Charset, i.Password); err != nil {
		return nil, err
	}

	returnCode := responseBizContentResult.Get(RETURN_CODE).Int()
	if returnCode != 0 {
		returnMsg := responseBizContentResult.Get(RETURN_MSG).String()
		fmt.Printf("调用工行API出错，API:【%s】,错误码=【%d】,错误信息：【%s】\n", action, returnCode, returnMsg)
	}
	//返回结果
	return &result, nil
}

//准备参数
func (i *IcbcClient) prepareParams(commonData map[string]string, bizContentData interface{}, action string) (params map[string]string, strToSign string, err error) {
	params = make(map[string]string, 0)
	params = commonData
	bizContentMarshal, err := json.Marshal(bizContentData)
	if err != nil {
		return
	}
	bizContentStr := string(bizContentMarshal)
	if i.IsNeedEncrypt { //需要对bizContentStr进行加密
		bizContentStr = AesCFBEncrypt(bizContentMarshal, i.EncryptKey)
	}
	params[BIZ_CONTENT_KEY] = bizContentStr
	//待签名数据
	path := fmt.Sprintf("%s?", action)
	strToSign = getSortStr(path, params)
	//签名
	signedStr, err := IcbcSignature(strToSign, i.SignType, i.PrivateKey, i.Charset, i.Password)
	if err != nil {
		return
	}
	params[SIGN] = signedStr
	return
}

//请求
func (i *IcbcClient) request(data map[string]string, url string, requestType string) (resBody []byte, err error) {
	req := resty.R().SetHeaders(map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"charset":      "UTF-8",
	}).SetFormData(data)
	var res *resty.Response
	if requestType == REQUEST_TYPE_POST {
		res, err = req.Post(url)
		if err != nil {
			return
		}
	} else if requestType == REQUEST_TYPE_GET {
		res, err = req.Get(url)
		if err != nil {
			return
		}
	} else {
		return []byte{}, requestTypeErr
	}

	return res.Body(), nil
}

//获取排序后的字符串数据
func getSortStr(signPrex string, mapParams map[string]string) (str string) {
	keys := make([]string, 0)
	for key, _ := range mapParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	signTxt := signPrex
	for _, key := range keys {
		if value, ok := mapParams[key]; ok {
			if value == "" {
				continue
			} else {
				signTxt += key + "=" + value + "&"
			}
		}
	}
	signTxt = strings.TrimRight(signTxt, "&") //去除最后一个&
	return signTxt
}

func getMsgId() string {
	timestamp := (time.Now().UnixNano()) / 1000000
	return fmt.Sprint(timestamp)
}

func getTimestamp() string {
	return time.Now().Format(Date_Time_Format)
}
