package example

import (
	"fmt"
	"github.com/mm-ooto/base/common/config"
	"github.com/mm-ooto/icbcSdkDemo/lib"
)

//回调通知示例
func Notify(data string) {
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
		NotifyUrl: config.AppConfig.String("notifyUrl"),
	}
	cli, err := lib.NewIcbcClient(baseParams)
	if err != nil {
		fmt.Printf("newIcbcClient error:%s\n", err.Error())
		return
	}
	gjsons, res, err := cli.DisposeNotifyData(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
	fmt.Println(gjsons.String())
}
