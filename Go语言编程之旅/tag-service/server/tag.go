package server

import (
	"context"
	"encoding/json"
	"github.com/wsloong/tag-service/pkg/bapi"
	"github.com/wsloong/tag-service/pkg/errcode"
	pb "github.com/wsloong/tag-service/proto"

)

type TagServer struct{}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}
	tagList := pb.GetTagListReply{}
	if err := json.Unmarshal(body,&tagList); err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return &tagList, nil
}
