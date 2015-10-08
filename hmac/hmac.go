package hmac

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"errors"
)

var (
	secret         []byte
	ErrBlankSecret = errors.New("Secret must not be blank!")
)

func Init(s []byte) {
	if s == nil || len(s) == 0 {
		panic(ErrBlankSecret)
	}
	secret = s
}

func Sign(s string) string {
	if secret == nil {
		panic(ErrBlankSecret)
	}
	h := hmac.New(sha512.New, secret)
	h.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func Equal(s, mac string) (bool, error) {
	if secret == nil {
		panic(ErrBlankSecret)
	}
	mac1 := hmac.New(sha512.New, secret)
	mac1.Write([]byte(s))
	mac2, err := base64.StdEncoding.DecodeString(mac)
	if err != nil {
		return false, err
	}
	return hmac.Equal(mac1.Sum(nil), mac2), nil
}
