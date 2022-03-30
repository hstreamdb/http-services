package main

import (
	"context"
	"github.com/hstreamdb/http-server/api"
	"github.com/hstreamdb/http-server/config"
	"github.com/hstreamdb/http-server/internal/rpcService"
	"github.com/hstreamdb/http-server/pkg/util"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func registerSignalHandler() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT)
		sig := <-sigChan
		util.Logger().Info("signal received", zap.String("signal", sig.String()))
		cancel()
	}()
	return ctx
}

func main() {
	conf := config.DefaultConfig()
	ctx := registerSignalHandler()

	client, err := rpcService.NewHStreamClient(conf.ServerUrl)
	if err != nil {
		util.Logger().Fatal("failed to connect to hstreamdb", zap.Error(err))
	}
	defer client.Close()

	router := api.InitRouter(client)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	util.Logger().Info("Start http server", zap.String("addr", server.Addr))

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			util.Logger().Error("Server error", zap.Error(err))
		}
		wg.Done()
	}()

	<-ctx.Done()
	if err := server.Shutdown(context.Background()); err != nil {
		util.Logger().Error("Server shutdown error", zap.Error(err))
	}
	wg.Wait()
}
