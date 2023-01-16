package main

import (
	"flag"
	"gitee.com/chunanyong/zorm"
	"github.com/zeromicro/go-zero/core/conf"
	"nFlow/common/database"
	"nFlow/flow/internal/config"
)

var configFile = flag.String("f", "../etc/flow.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	database.InitZorm(c.Mysql.DataSource)
	zorm.NewFinder()
}
