package main

import (
	"context"
	"github.com/yunduansing/go-study/study-grpc/proto"
	grpctool "github.com/yunduansing/gtools/grpc"
	"github.com/yunduansing/gtools/logger"
	"github.com/yunduansing/gtools/utils"
	"google.golang.org/grpc/metadata"
	"time"
)

func main() {
	createUserOrder()
}

func init() {
	logger.InitLog(logger.Config{
		Level:       "debug",
		FilePath:    "",
		LogType:     "",
		ServiceName: "go-study-grpc",
		MaxSize:     0,
		MaxAge:      0,
		BackupNum:   0,
		Compress:    false,
	})
}

type RpcClient struct {
	proto.OrderClient
}

var RpcClients RpcClient

func initOrderClient() {
	conn, err := grpctool.Init(grpctool.ClientConfig{
		Address:   "localhost",
		Port:      8080,
		ServerPem: "",
	})
	if err != nil {
		logger.Error("init grpc client conn err:", err)
		panic(err)
	}

	RpcClients.OrderClient = proto.NewOrderClient(conn)

}

func createUserOrder() {
	initOrderClient()
	in := &proto.CreateOrderReq{
		Amount: 200,
		User: &proto.CreateOrderUser{
			UserId:  10,
			Phone:   "18888888888",
			Address: "IFS",
			AreaId:  555550,
		},
		Goods: []*proto.CreateOrderGoods{
			{
				SkuId:  "1",
				Number: 1,
				Price:  100,
			},
		},
	}

	md := metadata.Pairs("requestId", utils.UUID(), "requestTime", utils.TimeToChineseStr(time.Now()), "clientIp", "110.110.110.110")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := RpcClients.OrderClient.CreateOrder(ctx, in)
	if err != nil {
		logger.Error("create order err:", err)
		return
	}
	logger.Info("create order res:", res)
}
