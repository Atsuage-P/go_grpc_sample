package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func myUnaryClientInteceptor1(ctx context.Context, method string, req, res interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//　前処理
	fmt.Println("[pre] my unary client interceptor 1", method, req)
	//　メインの処理
	err := invoker(ctx, method, req, res, cc, opts...)
	//　後処理
	fmt.Println("[post] my unary client interceptor 1", res)

	return err
}

func myUnaryClientInteceptor2(ctx context.Context, method string, req, res interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//　前処理
	fmt.Println("[pre] my unary client interceptor 2", method, req)
	//　メインの処理
	err := invoker(ctx, method, req, res, cc, opts...)
	//　後処理
	fmt.Println("[post] my unary client interceptor 2", res)

	return err
}
