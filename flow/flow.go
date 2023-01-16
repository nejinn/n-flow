package main

import (
	"flag"
	"fmt"
	"net/http"

	"nFlow/flow/internal/config"
	"nFlow/flow/internal/handler"
	"nFlow/flow/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/flow.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	//logx.MustSetup(c.LogConf)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(func(header http.Header) {
		header.Add("Access-Control-Allow-Headers", "X-Requested-With")
		header.Add("Access-Control-Allow-Headers", "ignorecanceltoken")
	}, nil, "*"))

	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
