package akey

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

// Rsa加密
/*
```javascript
<script src="/static/common/js/jsencrypt.min.js"></script>
<script>
    var passwd = 'abc';//原始密码
    var encrypt = new JSEncrypt();
    encrypt.setPublicKey($('#rsa_public_key').val());
    var lastpwd = encrypt.encrypt(passwd);//加密密码
</script>
```
*/

// Rsa加密，密钥格式 -----BEGIN PUBLIC KEY-----
func RsaEncrypt(src string, publicKey string) (string, error) {
	return rsaEncrypt(src, []byte(publicKey))
}

// Rsa加密，密钥格式 -----BEGIN PUBLIC KEY-----
func RsaEncryptPath(src string, publicKeyPath string) (string, error) {
	publicKey, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return "", err
	}
	return rsaEncrypt(src, publicKey)
}

func rsaEncrypt(src string, publicKey []byte) (string, error) {
	if len(src) == 0 {
		return "", errors.New("src can not be empty")
	}
	srcByte := []byte(src)

	// 解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("public key error")
	}
	// 解析公钥
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pubKey := key.(*rsa.PublicKey)

	// rsa加密
	dstByte, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, srcByte)
	if err != nil {
		return "", err
	}

	// 对rsa加密结果进行base64加密
	dst := base64.StdEncoding.EncodeToString(dstByte)
	return dst, nil
}

// Rsa解密，密钥格式 -----BEGIN PRIVATE KEY-----
func RsaDecrypt(src string, privateKey string) (string, error) {
	return rsaDecrypt(src, []byte(privateKey))
}

// Rsa解密，密钥格式 -----BEGIN PRIVATE KEY-----
func RsaDecryptPath(src string, privateKeyPath string) (string, error) {
	if privateKeyPath=="" {
		privateKeyPath="config/akey/rsa_private_key.pem"
	}
	privateKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return "", err
	}
	return rsaDecrypt(src, privateKey)
}
func base64replace(guid string) string{
	//不能被4整除,就补=号(对4取余 4-余数)
	if strings.Contains(guid, "-") || strings.Contains(guid, "_") {
		//补等号
		remainder := len(guid) % 4
		if remainder != 0 {
			//该补的等号的数量
			padlen := 4 - remainder
			for i := 0; i < padlen; i++ {
				guid = guid + "="
			}
		}
		guid = strings.Replace(guid, "-", "+", -1)
		guid = strings.Replace(guid, "_", "/", -1)
	}
	return guid
}

func rsaDecrypt(src string, privateKey []byte) (string, error) {
	if len(src) == 0 {
		return "", errors.New("src can not be empty")
	}
	// 对rsa加密结果进行base64解密
	src = base64replace(src)
	srcByte, _ := base64.StdEncoding.DecodeString(src)
	// 解密pem格式的私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error")
	}
	// 解析私钥
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	//公钥解密
	//pubdecrypt, err := gorsa.RSA.PubKeyDECRYPT(bytes)
	privKey := key.(*rsa.PrivateKey)
	// rsa解密
	dstByte, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, srcByte)
	if err != nil {
		return "", err
	}
	dst := string(dstByte)
	return dst, nil
}

// Rsa解密，密钥格式 -----BEGIN RSA PRIVATE KEY-----
func RsaDecryptPKCS1(src string, privateKey string) (string, error) {
	return rsaDecryptPKCS1(src, []byte(privateKey))
}

// Rsa解密，密钥格式 -----BEGIN RSA PRIVATE KEY-----
func RsaDecryptPKCS1Path(src string, privateKeyPath string) (string, error) {
	privateKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return "", err
	}
	return rsaDecryptPKCS1(src, []byte(privateKey))
}

func rsaDecryptPKCS1(src string, privateKey []byte) (string, error) {
	if len(src) == 0 {
		return "", errors.New("src can not be empty")
	}
	// 对rsa加密结果进行base64解密
	srcByte, _ := base64.StdEncoding.DecodeString(src)

	// 解密pem格式的私钥
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error")
	}
	// 解析私钥
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// rsa解密
	dstByte, err := rsa.DecryptPKCS1v15(rand.Reader, key, srcByte)
	if err != nil {
		return "", err
	}
	dst := string(dstByte)
	return dst, nil
}

//RSA公钥私钥产生
func GenRsaKey(bits int,KeyPath string) error {
	if KeyPath=="" {
		KeyPath="config/akey/"
	}
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create(KeyPath+"rsa_private_key.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create(KeyPath+"rsa_public_key.pub")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
