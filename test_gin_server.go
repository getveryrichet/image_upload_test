package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//file upload를 받는 함수

func upload(c *gin.Context) {
	//C = gin.Context 부분임

	file, header, err := c.Request.FormFile("file")

	//에러 발생시 처리하는 구문
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	//받은 File name
	filename := header.Filename

	//os 모듈 사용해서 public 이라는 디렉토리에 filename wjwkd
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	//defer이 뭔지 찾아볼 것
	defer out.Close()

	//제일 첫 구문엑 ㅏ져온 file과 out을 copy함
	_, err = io.Copy(out, file)

	if err != nil {
		log.Fatal(err)
	}

	filepath := "http://localhost:8080/file/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})

}

func main() {

	router := gin.Default()
	//router에 template에 있는 html파일 읽게 함
	router.LoadHTMLGlob("template/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "select_file.html", gin.H{})
	})

	router.POST("/upload", upload)

	router.StaticFS("/file", http.Dir("public"))
	router.Run(":8080")

}
