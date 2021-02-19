package lib

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

var (
	headers = map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"charset":      "UTF-8",
	}
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
	NotifyUrl      string
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
			NotifyUrl:      params.NotifyUrl,
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
	if apiIcbcClient.SignType == SIGN_TYPE_CA {
		if len(params.Ca) == 0 || len(apiIcbcClient.Password) == 0 {
			return nil, caOrPasswordEmptyErr
		}
		//去除证书数据中的空格
		params.Ca = strings.ReplaceAll(params.Ca, " ", "")
		apiIcbcClient.Ca = params.Ca
	}
	return
}

//执行请求
func (i *IcbcClient) Execute(bizContentData interface{}, action, method string) (*gjson.Result, error) {
	commonData := map[string]string{
		APP_ID:    i.AppId,                             // 必须 APP的编号
		MSG_ID:    getMsgId(),                          // 必须 消息通讯唯一编号
		FORMAT:    FORMAT_JSON,                         // 可选 请求参数格式，仅支持json
		CHARSET:   i.Charset,                           // 可选 字符集，缺省为UTF-8
		SIGN_TYPE: i.SignType,                          // 可选 加密方式
		TIMESTAMP: time.Now().Format(Date_Time_Format), // 必须 交易发生时间戳
	}
	if i.SignType == SIGN_TYPE_CA {
		commonData[CA] = i.Ca
	}
	params, _, err := i.prepareParams(commonData, bizContentData, action)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s%s", i.IcbcHost, action)
	//发送请求
	resBody, err := i.execRequest(params, url, method, &headers)
	if err != nil {
		return nil, err
	}
	//解析响应
	result := gjson.Parse(resBody)
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

//执行请求
func (i *IcbcClient) execRequest(data map[string]string, requestUrl, methodType string, headers *map[string]string) (resBody string, err error) {
	urlValue := url.Values{}
	for k, v := range data {
		urlValue.Add(k, v)
	}
	req, err := http.NewRequest(methodType, requestUrl, strings.NewReader(urlValue.Encode()))
	if err != nil {
		return "", err
	}
	if headers != nil {
		for k, v := range *headers {
			req.Header.Add(k, v)
		}
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(resByte), nil
}

//处理异步通知数据
func (i *IcbcClient) DisposeNotifyData(notifyData string) (*gjson.Result, *NotifyResponseReturn, error) {
	var (
		gjsonData  gjson.Result
		returnCode int64
		returnMsg  string
	)

	params := make(map[string]string)
	notifyDataRes, err := url.ParseQuery(notifyData)
	if err != nil {
		return nil, nil, err
	}
	params[FROM] = notifyDataRes.Get(FROM)
	params[API] = notifyDataRes.Get(API)
	params[APP_ID] = notifyDataRes.Get(APP_ID)
	params[FORMAT] = notifyDataRes.Get(FORMAT)
	params[CHARSET] = notifyDataRes.Get(CHARSET)
	params[ENCRYPT_TYPE] = notifyDataRes.Get(ENCRYPT_TYPE)
	params[TIMESTAMP] = notifyDataRes.Get(TIMESTAMP)
	params[SIGN_TYPE] = notifyDataRes.Get(SIGN_TYPE)
	responseBizContent := notifyDataRes.Get(BIZ_CONTENT_KEY)
	if i.IsNeedEncrypt { //如果是加密数据则对返回的数据先解密，后验签
		responseBizContent = AesCFBDecrypt(responseBizContent, i.EncryptKey)
	}
	params[BIZ_CONTENT_KEY] = responseBizContent
	//对请求结果验签
	signTxt := getSortStr(getPathByNotifyUrl(i.NotifyUrl)+"?", params)
	signData := notifyDataRes.Get(SIGN)
	signType := SIGN_TYPE_RSA //工行验签暂时只支持RSA
	if err := IcbcVerifySignature(signTxt, signData, signType, i.IcbcPublickKey, i.Charset, i.Password); err != nil {
		return nil, nil, err
	}

	params[SIGN] = signData
	gjsonData = gjson.Parse(responseBizContent)
	returnCode = gjsonData.Get(RETURN_CODE).Int()
	returnMsg = gjsonData.Get(RETURN_MSG).String()
	if returnCode != 0 {
		fmt.Printf("异步回调通知失败，错误码：【%d】,错误信息：【%s】\n", returnCode, returnMsg)
	}

	//异步回调通知响应参数组装
	notifyResponseReturn := NotifyResponseReturn{
		ResponseBizContent: ResponseBizContentM{
			ReturnCode: returnCode,
			ReturnMsg:  returnMsg,
			MsgId:      getMsgId(),
		},
		SignType: i.SignType,
	}
	//待签名字符串
	strToSign, _ := json.Marshal(notifyResponseReturn)
	signedStr, err := IcbcSignature(string(strToSign), i.SignType, i.PrivateKey, i.Charset, i.Password)
	if err != nil {
		return nil, nil, err
	}
	notifyResponseReturn.Sign = signedStr
	return &gjsonData, &notifyResponseReturn, nil
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

//https://studygolang.com/articles ---> /articles
//截取通知地址
func getPathByNotifyUrl(notifyUrl string) string {
	notifyUrls := strings.Split(notifyUrl, "//")
	if len(notifyUrls) != 2 {
		return ""
	}
	notifyUrl = notifyUrls[1]
	index := strings.Index(notifyUrls[1], "/")
	return notifyUrl[index:]
}