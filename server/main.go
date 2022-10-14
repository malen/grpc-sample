package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	user "aoisoft/user"
)

const (
	port = ":50051"
)

type UserService struct {
	// 实现 User 服务的业务对象
	user.UserServer
}

// UserService 实现了 User 服务接口中声明的所有方法
func (userService *UserService) UserIndex(ctx context.Context, in *user.UserIndexRequest) (*user.UserIndexResponse, error) {
	log.Printf("receive user index request: page %d page_size %d", in.Page, in.PageSize)

	return &user.UserIndexResponse{
		Err: 0,
		Msg: "success",
		Data: []*user.UserEntity{
			{Name: "malen", Age: 28},
			{Name: "xxxxx", Age: 29},
		},
	}, nil
}

func (userService *UserService) UserView(ctx context.Context, in *user.UserViewRequest) (*user.UserViewResponse, error) {
	log.Printf("receive user view request: uid %d", in.Uid)

	return &user.UserViewResponse{
		Err:  0,
		Msg:  "success",
		Data: &user.UserEntity{Name: "test", Age: 18},
	}, nil
}

func (userService *UserService) UserPost(ctx context.Context, in *user.UserPostRequest) (*user.UserPostResponse, error) {
	log.Printf("receive user post request: name %s password %s age %d", in.Name, in.Password, in.Age)

	return &user.UserPostResponse{
		Err: 0,
		Msg: "success",
	}, nil
}

func (userService *UserService) UserDelete(ctx context.Context, in *user.UserDeleteRequest) (*user.UserDeleteResponse, error) {
	log.Printf("receive user delete request: uid %d", in.Uid)

	return &user.UserDeleteResponse{
		Err: 0,
		Msg: "success",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// ################################################
	// 创建RPC服务方案1：不使用SSL/TLS
	// ################################################
	// 创建 RPC 服务
	// grpcServer := grpc.NewServer()

	// ################################################
	// 创建RPC服务方案2：使用SSL/TLS
	// ################################################
	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatalln(err)
	}
	// 创建 RPC 服务
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	// ↑↑↑　服务创建完成

	// 为 User 服务注册业务实现 将 User 服务绑定到 RPC 服务容器上
	user.RegisterUserServer(grpcServer, &UserService{})
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
