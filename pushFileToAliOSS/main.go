package main

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"mime/multipart"
	"net/http"
	"strings"
)

var bucket *oss.Bucket

func InitOSS() error {
	client, err := oss.New(
		viper.GetString("oss.endPoint"),
		viper.GetString("oss.accessKeyId"),
		viper.GetString("oss.accessKeySecret"))
	if err != nil {
		return err
	}

	bucket, err = client.Bucket(viper.GetString("oss.bucket"))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	InitConfig()

	err := InitOSS()
	if err != nil {
		logrus.Fatalln("OSS启动失败", err)
	}

	router := gin.Default()
	router.POST("/douyin/publish/action/", func(c *gin.Context) {
		var data, _ = c.FormFile("data")
		SaveVideo(data)
		c.JSON(http.StatusOK, gin.H{"msg": "挺好的"})
	})

	router.Run()
}

// SaveVideo 将视频流上传到oss，并返回视频地址
func SaveVideo(data *multipart.FileHeader) (string, error) {
	var file, _ = data.Open()
	var videoKey = strings.Replace(uuid.NewString(), "-", "", -1)

	if err := bucket.PutObject(videoKey, file); err != nil {
		return "", err
	}
	return viper.GetString("oss.root") + videoKey, nil
}

func InitConfig() {
	//workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./pushFileToAliOSS") // 获取配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalln(err)
		return
	}
}
