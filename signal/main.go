package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s := http.NewServeMux()
	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		log.Println(w, "Hello world!")
	})
	server := &http.Server{
		Addr:    ":8090",
		Handler: s,
	}
	go server.ListenAndServe()

	listenSignal(context.Background(), server)
}

func listenSignal(ctx context.Context, httpSrv *http.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-sigs:
		log.Println("notify sigs")
		httpSrv.Shutdown(ctx)
		log.Println("http shutdown")
	}
}

//func main() {
//	sigs := make(chan os.Signal)
//	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
//
//	// 监听所有信号
//	log.Println("listen sig")
//	signal.Notify(sigs)
//
//
//	// 打印进程id
//	log.Println("PID:", os.Getppid())
//	go func() {
//		for s := range sigs {
//			switch s {
//			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
//				log.Println("Program Exit...", s)
//			case syscall.SIGUSR1:
//				log.Println("usr1 signal", s)
//			case syscall.SIGUSR2:
//				log.Println("usr2 signal", s)
//			default:
//				log.Println("other signal", s)
//			}
//		}
//	}()
//
//	<-sigs
//}

//func GracefullExit() {
//	log.Println("Start Exit...")
//	log.Println("Execute Clean...")
//	log.Println("End Exit...")
//	os.Exit(0)
//}
