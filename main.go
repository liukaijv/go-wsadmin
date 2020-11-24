package main

import (
	"flag"
	"go-wsadmin/models"
	"go-wsadmin/pkg/win"
	"go-wsadmin/router"
	"go-wsadmin/utils"
	"log"
	"net/http"
)

func main() {

	var addr = flag.String("addr", ":8080", "http service address")

	models.InitDb()

	w := win.NewServer()
	
	w.SetNotFoundHandler(func(ctx win.Context) {
		ctx.ReplyError(utils.StatusCodeNotFound, "")
	})

	w.AddHandler("/ping", func(ctx win.Context) {
		ctx.Reply("pong")
	})

	router.InitRouter(w)

	http.HandleFunc("/ws", w.Serve)

	log.Println("server addr: " + *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
