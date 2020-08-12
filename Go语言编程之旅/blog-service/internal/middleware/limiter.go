// 限流的中间件
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wsloong/blog-service/pkg/app"
	"github.com/wsloong/blog-service/pkg/errcode"
	"github.com/wsloong/blog-service/pkg/limiter"
)

func RageLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			// 占用存储桶中立即可用的令牌数量，返回值为删除的令牌数
			// 如果没有令牌了，返回0
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
