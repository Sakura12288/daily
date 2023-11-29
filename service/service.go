package service

import (
	"context"
	"tiktok/pkg/pb"
)

type TestServer struct {
	pb.UnimplementedTestServer
}


func (t *TestServer) T(ctx context.Context,  req *pb.Req) (*pb.Rpn, error) {
	result := req.First + req.Name
	return &pb.Rpn{Mess: result}, nil 
}