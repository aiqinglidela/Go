package main

import (
	"TianCheng/model"
	"TianCheng/router"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

func main() {
	// 连接数据库
	err := model.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer model.Close()  // 程序退出关闭数据库连接
	// 注册路由
	r1:=router.Router01()
	server01 := &http.Server{
		Addr:         ":8081",
		Handler:      r1,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	r2:=router.Router02()
	server02 := &http.Server{
		Addr:         ":8082",
		Handler:      r2,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	r3:=router.Router03()
	server03 := &http.Server{
		Addr:         ":8083",
		Handler:      r3,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	r4:=router.Router04()
	server04 := &http.Server{
		Addr:         ":8084",
		Handler:      r4,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})
	g.Go(func() error {
		return server03.ListenAndServe()
	})
	g.Go(func() error {
		return server04.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
