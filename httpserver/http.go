package httpserver

import (
	"context"
	"log"
	"net/http"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	srv = &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Printf("couldnt write response error [%s]\n", err)
	}
}
