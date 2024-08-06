package crypto

import (
	"crypto/rand"
	"encoding/hex"
)

// 生成SM4密钥并返回十六进制表示
// 返回: 原始key, 十六进制表示, 错误
func GenerateSM4Key() ([]byte, string, error) {
	key := make([]byte, 16) // SM4密钥长度为16字节（128位）
	_, err := rand.Read(key)
	if err != nil {
		return nil, "", err
	}
	return key, hex.EncodeToString(key), nil
}
