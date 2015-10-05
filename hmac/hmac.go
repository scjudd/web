package hmac

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"errors"
)

var (
	Secret         []byte
	ErrBlankSecret = errors.New("secret must not be blank")
)

func Sign(s string) (string, error) {
	if Secret == nil {
		return "", ErrBlankSecret
	}
	h := hmac.New(sha512.New, Secret)
	h.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func Equal(s, mac string) (bool, error) {
	if Secret == nil {
		return false, ErrBlankSecret
	}
	mac1 := hmac.New(sha512.New, Secret)
	mac1.Write([]byte(s))
	mac2, err := base64.StdEncoding.DecodeString(mac)
	if err != nil {
		return false, err
	}
	return hmac.Equal(mac1.Sum(nil), mac2), nil
}
