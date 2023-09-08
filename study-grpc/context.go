package study_grpc

import (
	"context"
	"time"
)

type StudyContext struct {
	context.Context
	RequestId   string    `json:"requestId"`
	RequestTime time.Time `json:"requestTime"`
	ClientIp    string    `json:"clientIp"`
}
