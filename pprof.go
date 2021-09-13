package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
)

type Config struct {
	Host     string
	Port     int
	Disabled bool
}

// RunServer 启动 pprof 性能分析服务器
func RunServer(conf *Config) {
	if conf.Disabled {
		return
	}
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	server := http.Server{Addr: addr, Handler: mux}
	go func() {
		log.Println(fmt.Sprintf("pprof服务器已启动: %v/debug/pprof", addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(err.Error())
		}
	}()
}
