package user

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lzj09/chat-server/user"
	"github.com/lzj09/chat-server/utils"
	"k8s.io/klog/v2"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	userService := utils.Obtain(new(user.DefaultUserService)).(*user.DefaultUserService)
	user, err := userService.Login(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// 生成token
	now := time.Now()
	token, err := genToken(user.ID, now)
	if err != nil {
		klog.Errorf("gen token error: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "expires": now.Add(time.Duration(utils.AccessExpire) * time.Second)})
}

func Info(c *gin.Context) {
	// 解析用户id
	id := utils.UserInfo(c)

	userService := utils.Obtain(new(user.DefaultUserService)).(*user.DefaultUserService)
	u, err := userService.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	u.Password = ""
	c.JSON(http.StatusOK, u)
}

func genToken(id string, now time.Time) (string, error) {
	claim := TokenClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(utils.AccessExpire) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(utils.AccessSecret))
}
