package main

import (
	"context"
	"fmt"
	stlog "log"
	"path"

	"github.com/yikuanzz/distributed-system/log"
	"github.com/yikuanzz/distributed-system/service"
)

func main() {
	logPath := path.Join("..", "..", "runtime", "logs", "distributed.log")
	log.Run(logPath)
	host, port := "localhost", "8080"

	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host,
		port,
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
