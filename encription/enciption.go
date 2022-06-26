package encription

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(message string) string {
	hash := md5.Sum([]byte(message))
	return hex.EncodeToString(hash[:])
}
