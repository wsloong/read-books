package model

import "github.com/wsloong/blog-service/pkg/app"

// 针对 swagger 的注释
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         bool   `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}
