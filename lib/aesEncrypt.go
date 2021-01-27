package lib

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

//plainText:原始数据
//pwdKeyStr:密钥key
//Aes加密
func AesCBCEncrypt(plainText []byte, pwdKeyStr string) string {
	pwdKey := []byte(pwdKeyStr)
	cipherBlock, err := aes.NewCipher(pwdKey)
	if err != nil {
		return ""
	}
	blockSize := cipherBlock.BlockSize()
	srcByte := pKCS7Padding(plainText, blockSize) //填充字符串后的数据
	blockMode := cipher.NewCBCEncrypter(cipherBlock, pwdKey[:blockSize])
	dstByte := make([]byte, len(srcByte))
	blockMode.CryptBlocks(dstByte, srcByte)
	encStr := base64.StdEncoding.EncodeToString(dstByte)
	return encStr
}

//decryptStr:加密数据
//pwdKeyStr:密钥key
//Aes解密
func AesCBCDecrypt(decryptStr string, pwdKeyStr string) string {
	pwdKey := []byte(pwdKeyStr)
	srcByte, err := base64.StdEncoding.DecodeString(decryptStr)
	if err != nil {
		return ""
	}
	cipherBlock, err := aes.NewCipher(pwdKey)
	if err != nil {
		return ""
	}
	blockSize := cipherBlock.BlockSize()
	blockMode := cipher.NewCBCDecrypter(cipherBlock, pwdKey[:blockSize])
	dstByte := make([]byte, len(srcByte))
	blockMode.CryptBlocks(dstByte, srcByte)
	srcData, err := pKCS7UnPadding(dstByte)
	if err != nil {
		return ""
	}
	return srcData
}

//移除填充的字符串
func pKCS7UnPadding(encryptData []byte) (string, error) {
	len := len(encryptData)
	if len == 0 {
		return "", errors.New("encryptData is empty!")
	}
	paddingCount := int(encryptData[(len - 1)]) //最后一个数据，同时也是填充物的个数
	data := encryptData[:(len - paddingCount)]
	return string(data), nil
}

//填充字符串
func pKCS7Padding(plainText []byte, blockSize int) []byte {
	paddingCount := blockSize - len(plainText)%blockSize                  //需要填充字符串的量
	bytesRepeat := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount) //返回paddingCount个[]byte{byte(paddingCount)}串联形成的新的切片。
	return append(plainText, bytesRepeat...)
}

//AES加墨
func AesCFBEncrypt(plainText []byte, pwdKeyStr string) string {
	pwdKey := []byte(pwdKeyStr)
	cBlock, err := aes.NewCipher(pwdKey)
	if err != nil {
		return ""
	}
	blockSize := cBlock.BlockSize()
	stream := cipher.NewCFBEncrypter(cBlock, pwdKey[:blockSize])
	dstByte := make([]byte, len(plainText))
	stream.XORKeyStream(dstByte, plainText)
	dstStr := base64.StdEncoding.EncodeToString(dstByte)
	return dstStr
}

//AES解密
func AesCFBDecrypt(decryptStr string, pwdKeyStr string) string {
	srcStr, err := base64.StdEncoding.DecodeString(decryptStr)
	if err != nil {
		return ""
	}
	pwdKey := []byte(pwdKeyStr)
	cBlock, err := aes.NewCipher(pwdKey)
	if err != nil {
		return ""
	}
	blockSize := cBlock.BlockSize()
	stream := cipher.NewCFBDecrypter(cBlock, pwdKey[:blockSize])
	dstByte := make([]byte, len(srcStr))
	stream.XORKeyStream(dstByte, srcStr)
	return string(dstByte)
}
