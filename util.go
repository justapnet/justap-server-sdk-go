package justap

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"math/big"
)

func Base64Encode(data []byte) string {

	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(s string) ([]byte, error) {
	//base64解码
	return base64.StdEncoding.DecodeString(s)
}

func JsonEncode(v interface{}) ([]byte, error) {
	return json.Marshal(&v)
}

func JsonDecode(p []byte, v interface{}) error {
	obj := json.NewDecoder(bytes.NewBuffer(p))
	obj.UseNumber()
	return obj.Decode(&v)
}

func RsaPrivateKeySign(data []byte, priv *rsa.PrivateKey) (string, error) {
	if priv == nil {
		return "", errors.New("missing private key")
	}

	hashFunc := crypto.SHA256
	h := hashFunc.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	signBytes, err := rsa.SignPKCS1v15(rand.Reader, priv, hashFunc, hashed)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signBytes), nil
}

func GenSign(data []byte, privateKey []byte) (sign string, err error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	return RsaPrivateKeySign(data, priv)
}

func Verify(data string, publicKey string, sign string) (err error) {
	signData, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)

	hashFunc := crypto.SHA256
	h := hashFunc.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed, []byte(signData))
	if err != nil {
		return err
	}

	return
}

func Md5Hash(s string) string {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	b := h.Sum(nil)
	return fmt.Sprintf("%x", b)
}

func RandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
