package control

import (
	"context"
	"fmt"
	"net/http"
	"tiktok/pkg/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Publish(c *gin.Context) {
	useID := int64(1)
	title := c.PostForm("title")
	data, _ := c.FormFile("data")

	videoName := data.Filename
	if err := c.SaveUploadedFile(data, "/home/sakura/projects/tiktok/video/" + videoName); err != nil {
		fmt.Println(err)
		return
	}

	in := &pb.PublishReq{Title: title, UserID: useID, Filepath: "/home/sakura/projects/tiktok/video/" + videoName}

	conn, err := grpc.Dial("0.0.0.0:10086", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client := pb.NewTestClient(conn)
	_, err = client.Publish(context.Background(), in)
	if err != nil {
		fmt.Printf("err:%v", err)
		c.JSON(http.StatusOK, gin.H {
			"success": "NOT",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H {
		"success": "OK",
	})
}
