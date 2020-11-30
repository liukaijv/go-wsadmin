package main

import (
	"flag"
	"fmt"
	"go-wsadmin/common"
	"go-wsadmin/config"
	"go-wsadmin/models"
	"go-wsadmin/pkg/win"
	"go-wsadmin/router"
	"log"
	"net/http"
)

func main() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	var addr = flag.String("addr", ":8080", "http service address")
	var configFile = flag.String("config", "./config.json", "config file")
	var err error

	if err = config.Init(*configFile); err != nil {
		panic(fmt.Sprintf("加载配置文件错误：%s", err))
	}

	if err = models.InitDb(); err != nil {
		panic(fmt.Sprintf("数据库初始化错误：%s", err))
	}

	w := win.NewServer()

	w.SetNotFoundHandler(func(ctx win.Context) {
		ctx.ReplyError(common.StatusCodeNotFound, "")
	})

	w.AddHandler("/ping", func(ctx win.Context) {
		ctx.Reply("pong")
	})

	// 初始化路由
	router.InitRouter(w)

	http.HandleFunc("/ws", w.Serve)

	log.Println("server addr: " + *addr)
	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
