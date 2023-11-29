package minioio

import (
	"context"
	"fmt"
	"log"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	Endpoint    string
	VideoBucket string
	CoverBucket string
	client *minio.Client
}

const (
	endpoint        string = "192.168.3.205:9000"
	accessKeyID     string = "admin123"
	secretAccessKey string = "admin123"
	useSSL          bool   = false
)

var (
	minioServer *Minio
	err error
)

func Init()  {
	minioServer = &Minio{}
	minioServer.Endpoint = endpoint
	minioServer.client, err = minio.New(minioServer.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL})
	if err != nil {
		log.Fatalln("minio连接错误: ", err)
	}
	// log.Printf("%#v\n", client)
}

func (m *Minio) CreateBucket (bucketName string) {
	exists, _ := m.client.BucketExists(context.Background(), bucketName)
	if exists {
		log.Printf("bucket: %s已经存在", bucketName)
		return 
	}
	
	err = m.client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "cn-south-1", ObjectLocking: false})
	if err != nil {
		log.Println("创建bucket错误: ", err)
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
}

func (m *Minio) ListBuckets () {
	buckets, _ := m.client.ListBuckets(context.Background())
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
}

func (m *Minio) FileUploader(bucketName, objectName, filePath, contextType string) {
	object, err := m.client.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contextType})
	if err != nil {
		log.Println("上传失败：", err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", objectName, object.Size)
}

func (m *Minio) FileGet(bucketName, objectName, fileSavePath string) {
	err = m.client.FGetObject(context.Background(), bucketName, objectName, fileSavePath, minio.GetObjectOptions{})
	if err != nil {
		log.Println("下载错误: ", err)
	}
}

func (m *Minio) FilesDelete(bucketName, objectName string) {
	//删除一个文件
	_ = m.client.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{GovernanceBypass: true})

	//批量删除文件
	objectsCh := make(chan minio.ObjectInfo)
	go func() {
		defer close(objectsCh)
		options := minio.ListObjectsOptions{Prefix: "test", Recursive: true}
		for object := range m.client.ListObjects(context.Background(), bucketName, options) {
			if object.Err != nil {
				log.Println(object.Err)
			}
			objectsCh <- object
		}
	}()
	m.client.RemoveObjects(context.Background(), objectName, objectsCh, minio.RemoveObjectsOptions{})
}


func GetMinio() *Minio {
	return minioServer 
}