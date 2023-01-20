package main

import (
	"context"
	"flag"
	"fmt"
	"gitee.com/chunanyong/zorm"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"nFlow/common/database"
	"nFlow/flow/internal/config"
)

//var configFile = flag.String("f", "../../etc/flow.yaml", "the config file")

func main() {
	var configFile = flag.String("f", "../ect/flow.yaml", "the config file")
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	database.InitZorm(c.Mysql.DataSource)
	finder := zorm.NewFinder()
	findesql := `select id from n_flow_user where user_code in (?)`

	finder.Append(findesql, []string{"nFlow", "nFlow1"})
	fmt.Println(finder.GetSQL())
	var ctx = context.Background()

	var ids []string
	err := zorm.Query(ctx, finder, &ids, nil)

	logx.Infof("err: %v", err)

}
