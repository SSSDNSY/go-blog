package test

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestHttpServer2(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/", &helloHandler{})
	server := http.Server{
		Addr:         ":9090",
		Handler:      mux,
		WriteTimeout: 2 * time.Second, //设置响应超时时间
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		if err := server.Shutdown(context.Background()); nil != err {
			log.Fatal("shutdown server error:", err)
		}
	}()

	server.ListenAndServe()
	log.Println("Starting HTTP server...")
	log.Fatal(http.ListenAndServe("localhost:4000", mux))

}

type helloHandler struct{}

func (*helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Method)
	w.Write([]byte("Hello world!"))
}
