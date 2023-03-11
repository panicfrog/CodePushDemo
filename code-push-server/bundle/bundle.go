package main

import (
	"flag"
	"fmt"

	"code-push-server/bundle/internal/config"
	"code-push-server/bundle/internal/handler"
	"code-push-server/bundle/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/bundle-api.yaml", "the config file")
var logConfigFile = flag.String("l", "etc/log.yaml", "the log config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	var lc logx.LogConf
	conf.MustLoad(*logConfigFile, &lc)
	logx.MustSetup(lc)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...  \nwith \n  AK: %s\n  SK: %s\n", c.Host, c.Port, c.AccessKey, c.SecretKey)
	server.Start()
}
