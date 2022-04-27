package cipher

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func RandId(len int) (string, error) {
	b := make([]byte, len)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func RandIdNotBase64(len int) (string, error) {
	b := make([]byte, len)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func RandRequestID() string {
	idList := make([]string, 3)
	idList[0], _ = RandIdNotBase64(7)
	idList[1], _ = RandIdNotBase64(7)
	idList[2], _ = RandIdNotBase64(7)
	requestID := fmt.Sprintf("%s-%s-%s", idList[0], idList[1], idList[1])
	return requestID
}
