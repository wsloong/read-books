// 用于声明具体的错误
package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务器内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败, 找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败, Token错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败, Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败, Token生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")

	ErrorGetTagList = NewError(20010001, "获取标签列表失败")
	ErrorCreateTag  = NewError(20010002, "创建标签失败")
	ErrorUpdateTag  = NewError(20010003, "更新标签失败")
	ErrorDeleteTag  = NewError(20010004, "删除标签失败")
	ErrorCountTag   = NewError(20010005, "统计标签失败")

	ErrorGetArticle    = NewError(30010001, "获取标单个文章失败")
	ErrorGetArticles   = NewError(30010002, "获取文章列表")
	ErrorCreateArticle = NewError(30010003, "创建文章失败")
	ErrorUpdateArticle = NewError(30010004, "更新文章失败")
	ErrorDeleteArticle = NewError(30010005, "删除文章失败")
)
