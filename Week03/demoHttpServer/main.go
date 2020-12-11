package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error{
		return processHttp(ctx)
	})
	eg.Go(func()error{
		return processorSig(ctx)
	})
	if err := eg.Wait(); err!=nil {
		fmt.Println("eg got error:", err)
	}
}

func processHttp(ctx context.Context) error {
	myHandler := MyHttpHandler{}
	server := http.Server{
		Addr:    ":8081",
		Handler: &myHandler,
	}
	go func(){
		select {
		case <- ctx.Done():
			fmt.Println("Http server quit by the other routine")
			server.Shutdown(context.Background())
		}
	}()
	return server.ListenAndServe()
}

func processorSig(ctx context.Context) error {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)
	select {
	case <- ctx.Done():
		fmt.Println("Sig monitor quit by other routine")
		return fmt.Errorf("OtherRoutineErr")
	case <-ch:
		fmt.Println("Sig monitor quit by sig")
		return fmt.Errorf("SigErr")
	}
}

type MyHttpHandler struct{
}

func (h *MyHttpHandler)ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}