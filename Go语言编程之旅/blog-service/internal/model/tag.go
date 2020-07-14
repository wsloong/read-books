package model

import "github.com/wsloong/blog-service/pkg/app"

// 针对 swagger 的注释
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
