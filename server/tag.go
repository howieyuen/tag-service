package server

import (
	"context"
	"encoding/json"

	"github.com/howieyuen/tag-service/pkg/bapi"
	"github.com/howieyuen/tag-service/pkg/errcode"
	pb "github.com/howieyuen/tag-service/proto"
)

type TagServer struct {
	*pb.UnimplementedTagServiceServer
}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (s *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListResponse, error) {
	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}

	var tagList *pb.GetTagListResponse
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, err
	}
	return tagList, nil
}
