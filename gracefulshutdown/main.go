package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// START OMIT
func serve(ctx context.Context) (err error) {
	srv := httpServer()

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")
	<-ctx.Done()
	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {cancel()}()

	if err = srv.Shutdown(ctxShutDown); err != nil {log.Fatalf("server Shutdown Failed:%+s", err)}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {err = nil}

	return
}
// END OMIT

func httpServer() *http.Server{
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "okay")
		},
	))

	return &http.Server{
		Addr:    ":6969",
		Handler: mux,
	}
}

// START MAIN
func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	if err := serve(ctx); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}
}
// END MAIN