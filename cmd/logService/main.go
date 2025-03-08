package main

import (
	"context"
	"fmt"
	stlog "log"
	"path"

	"github.com/yikuanzz/distributed-system/log"
	"github.com/yikuanzz/distributed-system/registry"
	"github.com/yikuanzz/distributed-system/service"
)

func main() {
	logPath := path.Join("..", "..", "runtime", "logs", "distributed.log")
	log.Run(logPath)
	host, port := "localhost", "8080"

	r := registry.Registration{
		ServiceName: "LogService",
		ServiceURL:  fmt.Sprintf("http://%s:%s", host, port),
	}

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		func() {
			log.RegisterHandler()
		},
	)

	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down log service")
}
