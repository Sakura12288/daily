package service

import (
	"context"
	"fmt"
	"test/pkg/pb"
	"test/utils/ffmpeg"
	"test/utils/minioio"
)

type Tiktok struct {
	pb.UnimplementedTestServer
}

func (t *Tiktok) Publish(ctx context.Context, in *pb.PublishReq) (*pb.PublishRpn, error) {
	M := minioio.GetMinio()
	M.CreateBucket("mytest")
	M.FileUploader("mytest", "test1", in.Filepath, "video/mp4")
	picPath, _ := ffmpeg.GetSnapshot(in.Filepath,"/home/sakura/projects/tiktok/video/test1", 1)
	M.FileUploader("mytest", "nihao", picPath, "png")
	videoUrl := "http://" + M.Endpoint + "/" + "mytest/" + "test1"
    picUrl := "http://" + M.Endpoint + "/" + "mytest/" + "nihao"
	fmt.Println(videoUrl)
	fmt.Println(picUrl)
	return &pb.PublishRpn{}, nil
}