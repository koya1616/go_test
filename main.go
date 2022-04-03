package main

import "github.com/gin-gonic/gin"

import "net/http"

func main() {
    engine:= gin.Default()
    // htmlのディレクトリを指定
    engine.LoadHTMLGlob("templates/*")
    engine.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
             // htmlに渡す変数を定義
            "message": "hello gin",
        })
    })
    engine.Run(":3010")
}