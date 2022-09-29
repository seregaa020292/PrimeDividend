package secure

import (
	"crypto/md5" //nolint:gosec
	"encoding/hex"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text)) //nolint:gosec
	return hex.EncodeToString(hash[:])
}
