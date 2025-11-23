package middlewares

import (
	"avito-internship/internal/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger(logger *zap.SugaredLogger) *Logger {
	return &Logger{logger: logger}
}

func (l *Logger) LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(logger.WithLogger(c.Request.Context(), l.logger))
		c.Next()
	}
}
