package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/4nth0/goaway/internal/config"
	"github.com/4nth0/goaway/redirect"
	"github.com/4nth0/goaway/stats"
	"github.com/4nth0/goaway/store/fs"
	log "github.com/sirupsen/logrus"
)

type NopStatsWriter struct{}

func (w *NopStatsWriter) WriteLine(line string) error {
	log.Info("New line: ", line)
	return nil
}

func (w *NopStatsWriter) Close() {}

func main() {
	ctx := context.Background()
	configureLog()
	// @TODO Gracefull shutdown

	writer := &NopStatsWriter{}
	stats := stats.NewClient(writer)
	go stats.Collect(ctx)

	// @TODO Prometheus metrics
	// @TODO Healthcheck

	cfg := config.GetConfig()
	// @TODO Configurable log time zone

	store, err := NewStore(ctx, cfg.Store)
	if err != nil {
		log.Fatal(err)
	}

	client := redirect.NewClient(store)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		id := strings.TrimPrefix(r.URL.Path, "/")
		path := client.GetRedirectPath(id, r)

		broadcastInboundRequest(stats, id, r)

		if path != "" {
			w.Write([]byte(path))
			return
		}

		w.Write([]byte("Nop!"))
	})

	log.Fatal(http.ListenAndServe(cfg.ServerPort(), nil))
}

func broadcastInboundRequest(s *stats.Client, id string, req *http.Request) {
	if s.Hits != nil {
		inbound := stats.Hit{
			URL:     req.URL,
			Method:  req.Method,
			Headers: req.Header,
		}

		if req.Body != nil {
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Error("UNABLE_TO_READ_BODY")
			}
			req.Body.Close()
			inbound.Body = string(body)

			req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}

		s.Hits <- inbound
	}
}

func NewStore(ctx context.Context, cfg config.StoreConfig) (redirect.DBClient, error) {
	switch cfg.DBType {
	case "fs":
		return fs.NewClient(cfg.DBPath)
	}

	return nil, fmt.Errorf("UNKNOWN_DB_TYPE: %s", cfg.DBType)
}

func configureLog() {
	log.SetOutput(os.Stdout)

	logLevel := "info"

	if value, ok := os.LookupEnv("LOG_LEVEL"); ok {
		logLevel = value
	}

	logrusLevel, errLogLevel := log.ParseLevel(logLevel)

	if errLogLevel != nil {
		log.Fatalf("ENV LOG_LEVEL provided is not a viable option, can be either: panic, fatal, error, warn, info, debug, trace")
	}
	log.Printf("Set log level to: %s", logLevel)
	log.SetLevel(logrusLevel)
}
