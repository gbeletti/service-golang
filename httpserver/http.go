package httpserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gbeletti/service-golang/cryptoasset"
	"github.com/gbeletti/service-golang/queuerabbit"
	"github.com/go-chi/chi/v5"
)

var srv *http.Server

// Start starts the http server
func Start() {
	createServer()
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

// Shutdown shuts down the http server
func Shutdown(ctx context.Context) (done chan struct{}) {
	done = make(chan struct{})
	go func() {
		defer close(done)
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("couldnt shutdown server error [%s]\n", err)
		}
	}()
	return
}

func createServer() {
	mux := chi.NewRouter()
	mux.Get("/", helloWorld)
	mux.Get("/bitcoin/startdate/{startDate:[0-9]{4}-[0-9]{2}-[0-9]{2}}/enddate/{endDate:[0-9]{4}-[0-9]{2}-[0-9]{2}}", bitcoinVariation)
	srv = &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	queuerabbit.PublishTest(context.Background(), msg)
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	}
}

func bitcoinVariation(w http.ResponseWriter, r *http.Request) {
	startDate := chi.URLParam(r, "startDate")
	endDate := chi.URLParam(r, "endDate")
	variation, err := cryptoasset.GetBitcoinVariation(startDate, endDate)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
	}
	res, err := json.Marshal(variation)
	if err != nil {
		log.Printf("couldnt marshal the response to json error [%s]\n", err)
	}
	_, err = w.Write(res)
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	}
}
