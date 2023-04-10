package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/your-username/did-example/services"
)

// GenerateDID 生成DID的路由处理函数
func GenerateDID(c *gin.Context) {
	// 调用DID服务生成DID
	did, err := services.GenerateDID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回生成的DID
	c.JSON(http.StatusOK, gin.H{"did": did})
}

// VerifyDID 验证DID的路由处理函数
func VerifyDID(c *gin.Context) {
	// 获取请求参数
	var req struct {
		DID string `json:"did"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// 调用DID服务验证DID
	isValid := services.VerifyDID(req.DID)

	// 返回DID是否有效
	c.JSON(http.StatusOK, gin.H{"is_valid": isValid})
}
