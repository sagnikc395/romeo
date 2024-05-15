package controllers

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

// password hash
func HashPassword(password string) string {
	return ""
}

// function to verfiy password
func VerifyPassword(userPass string, providePass string) (bool, string) {
	return false, ""
}
