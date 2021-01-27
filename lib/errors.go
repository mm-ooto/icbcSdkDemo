package lib

import "errors"

var (
	icbcSignatureErr        = errors.New("Only support CA/RSA signature!")
	icbcVerifySignatureErr  = errors.New("Only support CA or RSA signature verify!")
	requestTypeErr          = errors.New("Only support GET or POST http method!")
	icbcVerifySignatureFail = errors.New("Icbc sign verify not passed!")
	icbcEncryptErr          = errors.New("Only support AES encrypt!")
	icbcDecryptErr          = errors.New("Only support AES decrypt!")
	icbcConfigInvalid       = errors.New("Invalid configuration data!")
	bizContentIsNilErr      = errors.New("Request params bizContent is Empty!")
	noData                  = errors.New("No Data!")
	systemErr               = errors.New("System error")
)
