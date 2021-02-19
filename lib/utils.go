package
lib

// 公钥转换
func getPemPublic(public_key string) string {
	res := "-----BEGIN PUBLIC KEY-----\n"
	strlen := len(public_key)
	for i := 0; i < strlen; i += 64 {
		if i+64 >= strlen {
			res += public_key[i:] + "\n"
		} else {
			res += public_key[i:i+64] + "\n"
		}
	}
	res += "-----END PUBLIC KEY-----"
	return res
}

// 私钥转换
func getPemPrivate(private_key string) string {
	res := "-----BEGIN RSA PRIVATE KEY-----\n"
	strlen := len(private_key)
	for i := 0; i < strlen; i += 64 {
		if i+64 >= strlen {
			res += private_key[i:] + "\n"
		} else {
			res += private_key[i:i+64] + "\n"
		}
	}
	res += "-----END RSA PRIVATE KEY-----"
	return res
}


