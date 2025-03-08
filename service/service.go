package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, serviceName, host, port string, registerHandlerFunc func()) (context.Context, error) {
	registerHandlerFunc()
	ctx = startService(ctx, serviceName, host, port)
	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = host + ":" + port

	// 	启动服务器
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	// 优雅关闭
	go func() {
		fmt.Printf("%v service started. Press any key to stop. \n", serviceName)
		var s string
		fmt.Scanln(&s)
		cancel()
	}()

	return ctx
}
