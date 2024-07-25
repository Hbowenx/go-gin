package oss

import (
	"cell_phone_store/structure"
	"context"
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"os"
)

//腾讯云oss文件上传
func UploadFile(file io.Reader, fileName string)(fileUrl string,err error) {
	var conf structure.Config
	if _, err := toml.DecodeFile(structure.ConfigPath, &conf); err != nil {
		return "",errors.New("文件上传配置文件读取失败！")
	}
	u, _ := url.Parse(conf.Oss.Url)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(conf.Oss.SecretId),
			SecretKey: os.Getenv(conf.Oss.SecretKey),
		},
	})
	// 使用提供的文件流和文件名来上传文件
	_, err = c.Object.Put(context.Background(), fileName, file, nil)
	if err != nil {
		return "",err
	}
	return conf.Oss.Url+"/"+fileName,nil
}