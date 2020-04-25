package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 计算SHA256散列值
func SHA256(str string, secret string) (string, bool) {
	if str == "" || secret == "" {
		Log("error", "计算SHA256散列值错误: 参数错误")
		return "", false
	}

	hash := sha256.New()

	_, err := hash.Write([]byte(str + secret))
	if err != nil {
		Log("error", fmt.Sprintf("计算SHA256散列值错误: %s", err.Error()))
		return "", false
	}

	hashSum := hash.Sum(nil)
	hashStr := hex.EncodeToString(hashSum)

	return hashStr, true
}