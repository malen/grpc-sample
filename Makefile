BINARY_SERVER=grpc_server
BINARY_CLIENT=grpc_client
# -w 去掉调试信息
# -s 去掉符号表
GOBUILD_SERVER=go build -ldflags '-s -w' -o ${BINARY_SERVER}
GOBUILD_CLIENT=go build -ldflags '-s -w' -o ${BINARY_CLIENT}
GOCLEAN=go clean

# 伪目标的作用：https://blog.csdn.net/tilblackout/article/details/114766598
.PHONY: compile
compile:
	./protoc --go_out=. --go-grpc_out=. proto/*.proto

# 启动客户端
.PHONY: client
client:
	go run --race client/main.go

# 启动服务端
.PHONY: server
server:
	go run --race server/main.go

# 需要开启SSL/TLS的时候，需要生成证书。查看证书细节：openssl x509 -in cert/server.crt -text
.PHONY: cert
cert:
	openssl genrsa -out cert/server.key 2048
	openssl req -nodes -new -x509 -sha256 -days 1825 -config cert/cert.conf -extensions 'req_ext' -key cert/server.key -out cert/server.crt

# 编译服务端和客户端
.PHONY: clean
clean:
	$(GOCLEAN)

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD_SERVER).exe ./server/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD_CLIENT).exe ./client/main.go
