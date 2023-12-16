package arsa

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "io/ioutil"
    "os"
    "sync"
)

var genLock sync.RWMutex

//RSA公钥私钥产生
func GenRsaKey(bits int, password ...string) (private string, public string, err error) {
    genLock.Lock()
    defer genLock.Unlock()

    privatePath := "tempPrivate.pem"
    publicPath := "tempPublic.pem"
    defer os.RemoveAll(privatePath)
    defer os.RemoveAll(publicPath)

    err = GenRsaKeyPath(bits, privatePath, publicPath, password...)
    if err != nil {
        return "", "", err
    }

    privateBytes, err := ioutil.ReadFile(privatePath)
    if err != nil {
        return "", "", err
    }
    private = string(privateBytes)

    publicBytes, err := ioutil.ReadFile(publicPath)
    if err != nil {
        return "", "", err
    }
    public = string(publicBytes)

    return private, public, nil
}

//RSA公钥私钥产生
func GenRsaKeyPath(bits int, privatePath string, publicPath string, password ...string) error {
    if bits <= 0 {
        bits = 2048
    }
    if privatePath == "" {
        privatePath = "tempPrivate.pem"
    }
    if publicPath == "" {
        publicPath = "tempPublic.pem"
    }

    // 生成私钥文件
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return err
    }
    X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
    privateBlock := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: X509PrivateKey,
    }
    if len(password) > 0 && password[0] != "" {
        privateBlock, err = x509.EncryptPEMBlock(rand.Reader, "RSA Private Key", X509PrivateKey, []byte(password[0]), x509.PEMCipherAES256)
        if err != nil {
            return err
        }
    }
    privateFile, err := os.Create(privatePath)
    if err != nil {
        return err
    }
    defer privateFile.Close()
    err = pem.Encode(privateFile, privateBlock)
    if err != nil {
        return err
    }

    // 生成公钥文件
    publicKey := &privateKey.PublicKey
    X509PublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
    if err != nil {
        return err
    }
    publicBlock := &pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: X509PublicKey,
    }
    publicFile, err := os.Create(publicPath)
    if err != nil {
        return err
    }
    defer publicFile.Close()
    err = pem.Encode(publicFile, publicBlock)
    if err != nil {
        return err
    }
    return nil
}
