package middleware

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"blog-service/pkg/limiter"
	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			if count := bucket.TakeAvailable(1); count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponseList(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
