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
	//c.~~ 데이터타입 뭐시기를 반환 하도록 돼있음.
	//그래서 만약에 에러가 발생하면 file errr ~~ 가 반환됨.
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

	//defer= 지연 실행임
	//특정 문장 ㅎ고은 함수를 나중에 실행하게 한다.
	//Java의 finally block 처럼 마지막에 clean-up을 위해 쓰임
	//여기서는 file을 out으로 옮기고난 뒤에 마지막에 close를 함으로써, 저장 시키는 역할을 하는 것처럼 보임.
	//아마 JSON을 return하고나면 끝이니까. 한 것. 실험해보니까 맞음.
	defer out.Close()

	//제일 첫 구문엑 ㅏ져온 file과 out을 copy함
	//file에 담긴 내용을 out 변수에 넣어줌.
	_, err = io.Copy(out, file)

	if err != nil {
		log.Fatal(err)
	}

	filepath := "http://localhost:8080/file/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})

}
func main_page(c *gin.Context) {
	//select_file.html이라는 디렉토리에 있는 html파일 넘겨줌.
	c.HTML(http.StatusOK, "select_file.html", gin.H{})
}

func main() {

	router := gin.Default()
	//router에 template에 있는 html파일 읽게 함
	router.LoadHTMLGlob("template/*")

	//처음에 / endpoint로 들어오게 되면, select_file.html을 받게 함.

	router.GET("/", main_page)

	//upload 라는 엔드포인트가 들어오게되면 위에서 정의한 upload 함수 실행
	router.POST("/upload", upload)

	// /file이라는 endpoint/ 뒤에 이미지 이름이 들어오면
	// 내 로컬 파일 시스템에 저장 돼 있는 디렉토리에 해당 파일이 있을 경우 연결해줌
	// 이 케이스는 public이라는 디렉토리가 있어서 가능
	router.StaticFS("/file", http.Dir("public"))
	router.Run(":8080")

}
