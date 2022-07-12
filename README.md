# beego-web

## grpc 的proto 協議文件的生成和安裝

```shell
brew tap grpc/grpc      
brew install  grpc
```

```shell
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative rpcprotos/go_server.proto
```
