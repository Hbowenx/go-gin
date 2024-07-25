package controller

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"context"
)


func HttpGet(url string, params url.Values) ([]byte, error) {
	// 将参数添加到URL中
	if len(params) > 0 {
		url += "?" + params.Encode()
	}

	// 发送GET请求
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func RandString(n int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func HandleError(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  err.Error(),
	})
}

func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"code": statusCode,
		"msg":  "success",
		"data": data,
	})
}