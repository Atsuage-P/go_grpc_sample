# gRPCのお試し
## 概要
- [zennの記事](https://zenn.dev/hsaki/books/golang-grpc-starting) を参考

## 動作確認
```
# サーバ側
go run ./cmd/server/main.go

# クライアント側
go run ./cmd/client/main.go
```

- `gRPCurl` による動作確認
```
# サーバ内のサービス一覧表示
grpcurl -plaintext localhost:8080 list

# 特定のサービスのメソッド一覧表示
grpcurl -plaintext localhost:8080 list <サービス名>
grpcurl -plaintext localhost:8080 list myapp.GreetingService

# メソッドの呼び出し
grpcurl -plaintext -d '{"name": "user"}' localhost:8080 <メソッド名>

# Unary RPC(1リクエスト-1レスポンス)
grpcurl -plaintext -d '{"name": "user"}' localhost:8080 myapp.GreetingService.Hello

# Server Streaming RPC(1リクエスト-複数レスポンス)
grpcurl -plaintext -d '{"name": "user"}' localhost:8080 myapp.GreetingService.HelloServerStream

# Client Streaming RPC(複数リクエスト-1レスポンス)
grpcurl -plaintext -d '{"name": "user1"}{"name": "user2"}{"name": "user3"}{"name": "user4"}{"name": "user5"}' localhost:8080 myapp.GreetingService.HelloClientStream
```
