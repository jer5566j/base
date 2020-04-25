package token

import (
	"base/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// 生成TOKEN
func GenerateToken(data interface{}, secret string) (string, error) {
	jsonData, err := json.Marshal(data)
	if nil != err {
		return "", err
	}

	payload := base64.StdEncoding.EncodeToString(jsonData)
	sign, ok := utils.SHA256(payload, secret)
	if false == ok {
		return "", err
	}

	token := fmt.Sprintf("%s.%s", payload, sign)
	return token, err
}

// Token校验
func VerifyToken(data Token, token string, secret string) bool {
	// 按 "." 截取TOKEN
	comma := strings.Split(token, ".")
	if len(comma) < 2 {
		return false
	}

	// "." 之前为负载
	payload := comma[0]
	// "." 之后为签名
	sign := comma[1]

	// 比对签名
	verifyStr, ok := utils.SHA256(payload, secret)

	if (false == ok) || (sign != verifyStr) {
		return false
	}

	// 负载BASE64解码
	decodeByte, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		utils.Log("error", fmt.Sprintf("BASE64解码失败: %s", err.Error()))
		return false
	}

	err = json.Unmarshal(decodeByte, data)
	if err != nil {
		utils.Log("error", fmt.Sprintf("JSON解码错误: %s", err.Error()))
		return false
	}

	// 判断是否超时
	nowTime := time.Now().Unix()
	if nowTime > data.Expires() {
		return false
	}

	// 返回数据
	return true
}
