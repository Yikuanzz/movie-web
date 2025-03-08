package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/yikuanzz/distributed-system/registry"
)

func main() {
	// 注册服务逻辑
	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort

	// 启动服务器
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	// 优雅关闭
	go func() {
		fmt.Printf("registry service started. Press any key to stop. \n")
		var s string
		fmt.Scanln(&s)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("shutting down registry service")
}
