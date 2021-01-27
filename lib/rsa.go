package lib

import (
	"github.com/wenzhenxi/gorsa"
)

//content：内容
//privateKey：私钥
//signType：签名类型
//签名
func rsaSign(content, privateKey, signType string) (string, error) {
	if signType == SIGN_TYPE_RSA {
		return gorsa.SignSha1WithRsa(content, privateKey)
	} else if signType == SIGN_TYPE_RSA2 {
		return gorsa.SignSha256WithRsa(content, privateKey)
	}
	return "", icbcSignatureErr
}

//content：内容
//signature：签名数据
//privateKey：私钥
//signType：签名类型
//验签
func rsaSignVerify(content, signature, publicKey, signType string) error {
	if signType == SIGN_TYPE_RSA {
		return gorsa.VerifySignSha1WithRsa(content, signature, publicKey)
	} else if signType == SIGN_TYPE_RSA2 {
		return gorsa.VerifySignSha256WithRsa(content, signature, publicKey)
	}
	return icbcVerifySignatureFail
}
