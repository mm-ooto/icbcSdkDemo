package lib

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type UiIcbcClient struct {
}

func (u *UiIcbcClient) Execute(baseParams *Base, action string, bizContentInterface interface{}) (uiAddress string, err error) {
	cli, err := NewIcbcClient(baseParams)
	if err != nil {
		return "", err
	}
	urlVlaue := url.Values{}
	urlVlaue.Add(APP_ID, cli.AppId)
	urlVlaue.Add(MSG_ID, getMsgId())
	urlVlaue.Add(FORMAT, FORMAT_JSON)
	urlVlaue.Add(CHARSET, cli.Charset)
	urlVlaue.Add(SIGN_TYPE, cli.SignType)
	urlVlaue.Add(TIMESTAMP, getTimestamp())
	if cli.SignType == SIGN_TYPE_CA {
		urlVlaue.Add(CA, cli.Ca)
	}
	bizContentBytes,err:=json.Marshal(bizContentInterface)
	if err!=nil{
		return "",err
	}
	bizContent := string(bizContentBytes)
	if len(bizContent) == 0 {
		return "", bizContentIsNilErr
	}
	if cli.IsNeedEncrypt {
		bizContent = AesCFBEncrypt(bizContentBytes, cli.EncryptKey)
	}
	urlVlaue.Add(BIZ_CONTENT_KEY, bizContent)
	path := fmt.Sprintf("%s?", action)
	strToSign := getSortStr2(path, urlVlaue)
	//签名
	signedStr, err := IcbcSignature(strToSign, cli.SignType, cli.PrivateKey, cli.Charset, cli.Password)
	if err != nil {
		return
	}
	urlVlaue.Add(SIGN, signedStr)
	//访问地址
	uiAddress = fmt.Sprintf("%s%s%s", cli.IcbcHost, path, urlVlaue.Encode())
	fmt.Println("ui访问地址为：", uiAddress)
	return
}

//获取排序后的字符串数据
func getSortStr2(signPrex string, urlValues url.Values) (str string) {
	keys := make([]string, 0)
	for key, _ := range urlValues {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	str = signPrex
	for _, key := range keys {
		v := urlValues.Get(key)
		if v == "" {
			continue
		}
		str += key + "=" + v + "&"
	}
	str = strings.TrimRight(str, "&") //去除最后一个&
	return
}
