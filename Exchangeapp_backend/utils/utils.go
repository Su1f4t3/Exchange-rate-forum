package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// 加密密码
func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash), err
}

// 生成 JWT Token
func GerenateJWT(username string) (string, error) {
	// 创建一个新的 JWT Token，使用 HS256 签名算法，并设置 Payload（Claims）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	// JWT 的 Signature 部分，使用密钥 "secret" 签名的前两部分（Header 和 Payload）
	signedToken, err := token.SignedString([]byte("secret"))
	return "Bearer " + signedToken, err
}

// 验证密码
func CheckPasswordHash(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

// 解析 JWT Token，获取用户名
func ParseJWT(tokenString string) (string, error) {

	// 移除 Bearer 前缀
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// 解析并验证 JWT Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// 验证签名方法是否为 HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return "", err
	}

	// 从解析后的 Token 中提取 Claims，并获取用户名
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)

		if !ok {
			return "", errors.New("username not found in token")
		}

		return username, nil
	}

	return "", err
}
