package main

import (
	"fmt"
	"log"
	"net/http"
)

//Request：来自用户的请求信息，包括 post、get、Cookie、url 等。
//Response：服务器返回给客户端的信息。
//Connect：用户的每次的请求连接
//Handler：处理请求和生成返回信息的处理逻辑
//
//
//1. 创建 Listen Socket，监听指定端口，等待客户端请求。
//2. Listen Socket 接受客户端请求，得到 Client Socket，接下来通过 Client Socket与客户端通信。
//3. 处理客户端请求。首先从 Client Socket 获取 HTTP 请求数据，然后交给相应的 handler 处理请求，handler 处理完毕后再通过 Client Socket 返回给客户端。

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PATH: ", r.URL.Path)
	fmt.Println("METHOD: ", r.Method)
	fmt.Println()
	_, _ = fmt.Fprintf(w, "Hello Wrold!") //这个写入到w的是输出到客户端的

}

func main() {
	http.HandleFunc("/", HandleIndex)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}
