package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/yikuanzz/distributed-system/registry"
)

func Start(ctx context.Context, host, port string, reg registry.Registration, registerHandlerFunc func()) (context.Context, error) {
	registerHandlerFunc()
	ctx = startService(ctx, reg.ServiceName, host, port)

	// 注册服务到注册中心
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
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
