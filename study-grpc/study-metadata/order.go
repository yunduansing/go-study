package study_metadata

import (
	"context"
	"github.com/yunduansing/go-study/study-grpc/proto"
	"github.com/yunduansing/gtools/logger"
	"github.com/yunduansing/gtools/utils"
	"time"
)

type OrderGrpcServer struct {
	proto.OrderServer
}

func (s *OrderGrpcServer) CreateOrder(ctx context.Context, in *proto.CreateOrderReq) (*proto.CreateOrderRes, error) {
	logger.Info("收到rpc create请求,in=", in)

	res := &proto.CreateOrderRes{
		OrderId:     "1",
		CreatedTime: utils.TimeToChineseStr(time.Now()),
		State:       0,
	}
	return res, nil
}
