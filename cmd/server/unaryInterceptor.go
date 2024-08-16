package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// インターセプタ(gRPCのミドルウェア)
func myUnaryServerInterceptor1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	//　前処理
	log.Println("[pre] my unary server interceptor 1: ", info.FullMethod, req)
	// メインの処理
	res, err := handler(ctx, req)
	// 後処理
	log.Println("[post] my unary server interceptor 1: ", res)

	return res, err
}

func myUnaryServerInterceptor2(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("[pre] my unary server interceptor 2: ", info.FullMethod, req)
	res, err := handler(ctx, req) // 本来の処理
	log.Println("[post] my unary server interceptor 2: ", res)
	return res, err
}
