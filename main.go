package main

import (
	"context"
	"flag"
	"fmt"
	"kit-clean-app/app/order"
	"kit-clean-app/app/product"
	"kit-clean-app/db"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const defaultPort = "8080"

func main() {
	var httpAddr = flag.String("http.addr", ":"+defaultPort, "HTTP listen address")

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	idb, ctx := db.New()

	productRepo := product.NewRepository(idb.Client)
	productSvc := product.NewService(productRepo)
	productSvc = product.NewLoggingService(log.With(logger, "component", "product"), productSvc)

	orderRepo := order.NewRepository(idb.Client)
	orderSvc := order.NewService(idb, orderRepo, productRepo)
	orderSvc = order.NewLoggingService(log.With(logger, "component", "order"), orderSvc)

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()
	mux.Handle("/v1/products", product.MakeHandler(productSvc, httpLogger))
	mux.Handle("/v1/orders", order.MakeHandler(orderSvc, httpLogger))

	http.Handle("/", accessControl(mux, ctx))
	http.Handle("/metrics", promhttp.Handler())

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func accessControl(h http.Handler, ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
