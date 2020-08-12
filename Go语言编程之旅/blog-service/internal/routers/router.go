package routers

import (
	"net/http"
	"time"

	"github.com/wsloong/blog-service/pkg/limiter"

	"github.com/wsloong/blog-service/internal/routers/api"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/wsloong/blog-service/docs"
	"github.com/wsloong/blog-service/global"
	"github.com/wsloong/blog-service/internal/middleware"
	v1 "github.com/wsloong/blog-service/internal/routers/api/v1"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.RageLimiter(methodLimiters))                             // 注册限流中间件
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout)) // 注册超时控制中间件
	r.Use(middleware.Translations())                                          // 注册翻译的中间件

	// swagger 文档的路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tag := v1.NewTag()
	article := v1.NewArticle()
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile) // 文件上传
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.GET("/auth", api.GetAuth) // 认证

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
