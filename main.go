package main

import (
	"tiktok/routes"
)

func main() {
	// lis, err := net.Listen("tcp", "0.0.0.0:10086")
	// if err != nil {
	// 	fmt.Println("连接失败", err.Error())
	// 	return
	// }
	// server := &service.TestServer{}
	// s := grpc.NewServer()
	// pb.RegisterTestServer(s, server)
	// fmt.Println("成功")
	// if err := s.Serve(lis); err != nil {
	// 	fmt.Println("失败", err.Error())
	// 	return
	// }
	// fmt.Println("成功")

	// conn, err := grpc.Dial("0.0.0.0:10086", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer conn.Close()
	// client := pb.NewTestClient(conn)
	// _, err = client.Publish(context.Background(), &pb.PublishReq{Title: "nihao", Filepath: "/home/sakura/projects/tiktok/video/IMG_0822.mp4"})
	// if err != nil {
	// 	fmt.Printf("err:%v", err)
	// 	return
	// }
	routes.Init()

}
