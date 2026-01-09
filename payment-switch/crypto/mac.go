package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func VerifyMAC(secret, data, mac string) bool {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil)) == mac
}
