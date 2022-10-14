package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	user "aoisoft/user"
)

const (
	address = "192.168.11.88:50051"
)

func main() {

	// ################################################
	// 建立链接方案1：不使用SSL/TLS
	// ################################################
	// conn, err := grpc.Dial(address,
	// 	// コネクションでSSL/TLSを使用しない
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// 	// コネクションが確立されるまで待機する(同期処理をする)
	// 	grpc.WithBlock())

	// ################################################
	// 建立链接方案2：使用SSL/TLS
	// ################################################
	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "localhost")
	if err != nil {
		log.Fatalln(err)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// ↑↑↑　建立链接完成

	defer conn.Close()

	userClient := user.NewUserClient(conn)

	// 设定请求超时时间 3s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// UserIndex 请求
	userIndexReponse, err := userClient.UserIndex(ctx, &user.UserIndexRequest{Page: 1, PageSize: 12})
	if err != nil {
		log.Printf("user index could not greet: %v", err)
	}

	if userIndexReponse.Err == 0 {
		log.Printf("user index success: %s", userIndexReponse.Msg)
		// 包含 UserEntity 的数组列表
		userEntityList := userIndexReponse.Data
		for _, row := range userEntityList {
			fmt.Println(row.Name, row.Age)
		}
	} else {
		log.Printf("user index error: %d", userIndexReponse.Err)
	}

	// UserView 请求
	userViewResponse, err := userClient.UserView(ctx, &user.UserViewRequest{Uid: 1})
	if err != nil {
		log.Printf("user view could not greet: %v", err)
	}

	if userViewResponse.Err == 0 {
		log.Printf("user view success: %s", userViewResponse.Msg)
		userEntity := userViewResponse.Data
		fmt.Println(userEntity.Name, userEntity.Age)
	} else {
		log.Printf("user view error: %d", userViewResponse.Err)
	}

	// UserPost 请求
	userPostReponse, err := userClient.UserPost(ctx, &user.UserPostRequest{Name: "malen", Password: "666", Age: 29})
	if err != nil {
		log.Printf("user post could not greet: %v", err)
	}

	if userPostReponse.Err == 0 {
		log.Printf("user post success: %s", userPostReponse.Msg)
	} else {
		log.Printf("user post error: %d", userPostReponse.Err)
	}

	// UserDelete 请求
	userDeleteReponse, err := userClient.UserDelete(ctx, &user.UserDeleteRequest{Uid: 1})
	if err != nil {
		log.Printf("user delete could not greet: %v", err)
	}

	if userDeleteReponse.Err == 0 {
		log.Printf("user delete success: %s", userDeleteReponse.Msg)
	} else {
		log.Printf("user delete error: %d", userDeleteReponse.Err)
	}
}
