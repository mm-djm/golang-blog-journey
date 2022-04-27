package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"golang-blog-journey/util/log"
)

//Use AES-128
var PwdKey = []byte("ABCDABCDABCDABCD")

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		log.Errorf("pkcs7UnPadding data error")
		return nil, errors.New("data error")
	}
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Errorf("[AesEncrypt] %v", err)
		return nil, err
	}
	blockSize := block.BlockSize()
	encryptBytes := pkcs7Padding(data, blockSize)
	crypted := make([]byte, len(encryptBytes))
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Errorf("[AesDecrypt] %v", err)
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		log.Errorf("[pkcs7UnPadding] %v", err)
		return nil, err
	}
	return crypted, nil
}

func EncryptByAes(data []byte) (string, error) {
	res, err := AesEncrypt(data, PwdKey)
	if err != nil {
		log.Errorf("[EncryptByAes] %v", err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

func DecryptByAes(data string) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Errorf("[DecryptByAes] %v", err)
		return nil, err
	}
	return AesDecrypt(dataByte, PwdKey)
}
