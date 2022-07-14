package utils

import (
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

// UserInfo 解析用户信息
func UserInfo(c *gin.Context) string {
	payload, ok := c.Get(JWTPayloadKey)
	if !ok {
		klog.Errorf("jwt payload not found")
		return ""
	}

	principle, ok := payload.(string)
	if !ok {
		klog.Error("bad jwt payload")
		return ""
	}

	return principle
}
