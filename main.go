package main

import (
	"log"
	"net/http"

	. "github.com/jabin/Chatplatm/config"
	. "github.com/jabin/Chatplatm/routes"
)

func main() {
	startWebServer("8081")
}

func startWebServer(port string) {

	config := LoadConfig()
	r := NewRouter()

	// 处理静态资源文件
	assets := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r) // 通过 router.go 中定义的路由器来分发请求

	log.Println("Starting HTTP service at " + config.App.Address)
	err := http.ListenAndServe(config.App.Address, nil)
	if err != nil {
		log.Println("An error occured starting HTTP listener at " + config.App.Address)
		log.Println("Error: " + err.Error())
	}
}
