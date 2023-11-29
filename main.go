package main

import (
	"fmt"
	"net"
	"test/pkg/pb"
	"test/service"
	"test/utils/minioio"

	"google.golang.org/grpc"
)

func main() {
	// conn, err := grpc.Dial("0.0.0.0:10086", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer conn.Close()
	// client := pb.NewTestClient(conn)
	// resp, err := client.T(context.Background(), &pb.Req{First: "ni", Name: "sb"})
	// if err != nil {
	// 	fmt.Printf("err:%v", err)
	// 	return
	// }
	// fmt.Println(resp.Mess)
	minioio.Init()
	// Minio := minioio.GetMinio()
	// Minio.CreateBucket("mytest")
	// Minio.FileGet("mytest", "test", "./video/test1.mp4")
	// url := "http://" + Minio.Endpoint + "/" + "mytest/" + "test"
	// fmt.Println(url)

	lis, err := net.Listen("tcp", "0.0.0.0:10086")
	if err != nil {
		fmt.Println("连接失败", err.Error())
		return
	}
	server := &service.Tiktok{}
	s := grpc.NewServer()
	pb.RegisterTestServer(s, server)
	fmt.Println("成功")
	if err := s.Serve(lis); err != nil {
		fmt.Println("失败", err.Error())
		return
	}
	fmt.Println("成功")
}
