package service

import (
	"context"

	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/wsloong/blog-service/global"

	dao2 "github.com/wsloong/blog-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao2.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao2.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
