package main

import (
	"context"
	"flag"
	"fmt"
	"kit-clean-app/app/order"
	"kit-clean-app/app/product"
	"kit-clean-app/db"
	"kit-clean-app/pkg/external/exchangerate"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const defaultPort = "8080"

func main() {
	var httpAddr = flag.String("http.addr", ":"+defaultPort, "HTTP listen address")

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	if err := godotenv.Load(".env"); err != nil {
		level.Error(logger).Log("err", err)
		return
	}

	idb, ctx := db.New()

	exchangeRateAPI, err := exchangerate.New(os.Getenv("EXCHANGE_RATE_API_BASE_URL"), os.Getenv("EXCHANGE_RATE_API_KEY"))
	if err != nil {
		level.Error(logger).Log("err", err)
		return
	}

	productStore := product.NewStore(idb.Client)
	productSvc := product.NewService(productStore, exchangeRateAPI)
	productSvc = product.NewLoggingService(log.With(logger, "component", "product"), productSvc)

	orderStore := order.NewStore(idb.Client)
	orderSvc := order.NewService(idb, orderStore, productStore)
	orderSvc = order.NewLoggingService(log.With(logger, "component", "order"), orderSvc)

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()
	mux.Handle("/v1/products/", product.MakeHandler(productSvc, httpLogger))
	mux.Handle("/v1/orders/", order.MakeHandler(orderSvc, httpLogger))

	http.Handle("/", accessControl(mux, ctx))
	http.Handle("/metrics", promhttp.Handler())

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal, 1)
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
