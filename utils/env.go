package utils

import (
	"os"
	"strconv"
)

const JWTPayloadKey = "JWT_PAYLOAD"

var AccessSecret string
var AccessExpire int64

func init() {
	AccessSecret = GetEnv("ACCESS_SECRET", "lzj09-chat-server-@#$%")
	expire := GetEnv("ACCESS_EXPIRE", "86400")
	accessExpire, err := strconv.ParseInt(expire, 10, 64)
	if err != nil {
		panic(err)
	}
	AccessExpire = accessExpire
}

// GetEnv 获取环境变量
func GetEnv(key string, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}
