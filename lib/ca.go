package lib

//证书签名
func icbcCaSign(content, privateKey, pasword string) (string, error) {
	if len(content) <= 0 {
		return "", caOrPasswordEmptyErr
	}
	return "", nil
}

//证书验签
func icbcVerifySign(ontent, privateKey, pasword string) {

}
