package middlewares

import (
	"exchangeapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头中获取 Authorization token
		token := ctx.GetHeader("Authorization")

		// 检查 token 是否存在
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			ctx.Abort() // 停止处理请求
			return
		}

		// 解析 JWT token，获取用户名
		username, err := utils.ParseJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			ctx.Abort() // 停止处理请求
			return
		}

		// 将用户名存储在上下文中，以便后续处理使用
		ctx.Set("username", username)
		ctx.Next() // 继续处理请求
	}
}
