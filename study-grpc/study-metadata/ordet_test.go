package study_metadata

import (
	"github.com/yunduansing/go-study/study-grpc/proto"
	grpctool "github.com/yunduansing/gtools/grpc"
	"github.com/yunduansing/gtools/logger"
	"google.golang.org/grpc"
	"testing"
)

func TestOrderGrpcServer_CreateOrder(t *testing.T) {
	startGrpc()
}

func startGrpc() {
	var servers = []grpctool.GrpcServerHandler{
		func(server *grpc.Server) {
			proto.RegisterOrderServer(server, &OrderGrpcServer{})
		},
	}

	err := grpctool.Run(grpctool.ServerConfig{
		Port: 8080,
	}, servers...)
	if err != nil {
		logger.Error("start grpc err:", err)
		panic(err)
	}
}
