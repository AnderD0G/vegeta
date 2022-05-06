package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vegeta/pkg"
)

const AuthUserKey = "user"

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware(level int) func(c *gin.Context) {
	abort := func(c *gin.Context) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		code := c.Request.Header.Get("code")
		if code == "" {
			if authHeader == "" {
				abort(c)
			}
			//拿到auth 进行解析
			mc, err := pkg.ParseToken(authHeader)
			if err != nil {
				if authHeader == "" {
					abort(c)
				}
			}
			// 将当前请求的username信息保存到请求的上下文c上
			c.Set(AuthUserKey, mc.Id)
			c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		}

	}
}
