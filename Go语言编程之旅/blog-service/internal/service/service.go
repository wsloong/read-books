package service

import (
	"context"

	"github.com/wsloong/blog-service/global"

	dao2 "github.com/wsloong/blog-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao2.Dao
}

func New(ctx context.Context) Service {
	svc := Service{
		ctx: ctx,
		dao: dao2.New(global.DBEngine),
	}
	return svc
}
