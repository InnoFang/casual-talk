package main

import (
	"net/http"
	"fmt"
)

/**
回调函数：从 Request 中提取相关信息，通过 ResponseWriter 接口将响应返回给客户端
 */
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
