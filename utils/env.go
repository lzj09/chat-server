package utils

import "os"

// GetEnv 获取环境变量
func GetEnv(key string, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return def
}
