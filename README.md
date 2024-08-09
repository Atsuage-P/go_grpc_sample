# gRPCのお試し
## 概要
- [zennの記事](https://zenn.dev/hsaki/books/golang-grpc-starting) を参考

## 動作確認
- サーバ起動
```
go run main.go
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
grpcurl -plaintext -d '{"name": "user"}' localhost:8080 myapp.GreetingService.Hello
```
