package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
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

var (
	address     = flag.String("address", "localhost:8080", "server's listening address.")
	servicesUrl = flag.String("services-url", "localhost:6580", "hstreamdb services servicesUrl, split by comma.")
	logLevel    = flag.String("log-level", "info", "log level, support debug, info, warn, error, fatal, panic")
	debugMode   = flag.Bool("debug-mode", false, "use debug mode")
)

// @title HStreamDB-Server API
// @version 0.1.0
// @description http server for HStreamDB
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	flag.Parse()
	conf := config.NewConfig(*servicesUrl, *logLevel)
	ctx := registerSignalHandler()

	client, err := rpcService.NewHStreamClient(conf.ServerUrl)
	if err != nil {
		util.Logger().Fatal("failed to connect to hstreamdb", zap.Error(err))
	}
	defer client.Close()

	if !*debugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := api.InitRouter(client)
	server := &http.Server{
		Addr:    *address,
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
