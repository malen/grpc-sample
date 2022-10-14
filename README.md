# grpc-sample

# 事前准备
下载protoc 可执行文件，用来生成接口库(pb文件)
https://github.com/protocolbuffers/protobuf/releases

项目中用到以下pacakge
go get google.golang.org/protobuf/cmd/protoc-gen-go   
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

## 生成pb文件(可以直接通过make compile来生成)
protoc --go_out=./user --go_opt=paths=source_relative --go-grpc_out=./user --go-grpc_opt=paths=source_relative helloworld.proto
protoc --go_out=./user --go_opt=paths=source_relative --go-grpc_out=./user --go-grpc_opt=paths=source_relative user.proto

# 执行
## 生成证书
make cert

## 生成文件
make compile

## 启动服务端
make server

## 启动客户端
make client

# 参考
https://segmentfault.com/a/1190000019216566
https://github.com/grpc/grpc-go/blob/master/examples/helloworld/helloworld/helloworld_grpc.pb.go