package lib

//签名
func IcbcSignature(strToSign, signType, privateKey, charset, password string) (signStr string, err error) {
	switch signType {
	case SIGN_TYPE_CA:
		break
	case SIGN_TYPE_RSA, SIGN_TYPE_RSA2:
		signStr, err = rsaSign(strToSign, privateKey, signType)
		break
	default:
		err = icbcSignatureErr
	}
	return
}

//验签
func IcbcVerifySignature(strToSign, signedData, signType, publicKey, charset, password string) (err error) {
	switch signType {
	case SIGN_TYPE_CA:
		break
	case SIGN_TYPE_RSA, SIGN_TYPE_RSA2:
		err = rsaSignVerify(strToSign, signedData, publicKey, signType)
		break
	default:
		err = icbcVerifySignatureErr
	}
	return
}
