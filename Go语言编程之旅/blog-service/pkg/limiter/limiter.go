package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 不同的接口可能使用不同的限流器
// 因此这里定义为接口类型
type LimiterIface interface {
	Key(c *gin.Context) string                          // 对应限流器的键值对名称
	GetBucket(key string) (*ratelimit.Bucket, bool)     // 获取令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface // 新增多个令牌桶
}

// 用于存储令牌桶与键值对名称的映射关系
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

// 定义令牌桶的一些响应规则属性
type LimiterBucketRule struct {
	Key          string        // 键值对名称
	FillInterval time.Duration // 间隔多久时间放N个令牌
	Capacity     int64         // 令牌桶的容量
	Quantum      int64         // 每次到达间隔时间后所放的具体令牌数量
}
