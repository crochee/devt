package middleware

import (
	"github.com/crochee/lirity/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RequestLogger 设置请求日志
func RequestLogger(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		c.Request = c.Request.WithContext(logger.With(ctx, log))
		c.Next()
	}
}
