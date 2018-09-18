package test

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"
)

// RunOneShotHTTPServer executes an http server and serves on localhost at
// the specified port using the provided handlefunc. 'path' is the
// absolute path on the server where the handler will be maped, usually "/".The
// server will exit after one request or when the timeout expires which ever
// occures first. There is also a shutdown  path at /api/killserver that
// will stop the server when it recieves a GET request.
func RunOneShotHTTPServer(port int, timeout time.Duration, path string, handler http.HandlerFunc) {
	mux := http.NewServeMux()
	idleConnsClosed := make(chan struct{})

	// kill server on request
	mux.HandleFunc("/api/killserver", func(w http.ResponseWriter, r *http.Request) {
		close(idleConnsClosed)
	})

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, reply)
		handler(w, r)
		close(idleConnsClosed)
	})

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: mux,
	}

	go func() {
		<-idleConnsClosed

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	go func() {
		time.Sleep(timeout)
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

// RunTimedHTTPServer executes an http server and serves on localhost at the
// specified port using the provided handlefunc. The server will exit when the
// timeout expires.
func RunTimedHTTPServer(port int, timeout time.Duration, handler http.HandlerFunc) {
	mux := http.NewServeMux()
	idleConnsClosed := make(chan struct{})

	// kill server on request
	mux.HandleFunc("/api/killserver", func(w http.ResponseWriter, r *http.Request) {
		close(idleConnsClosed)
	})

	mux.HandleFunc("/api/project_analyses/search", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, reply)
		handler(w, r)
	})

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: mux,
	}

	go func() {
		<-idleConnsClosed

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	go func() {
		time.Sleep(timeout)
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

// RunTimedHTTPHandler executes an http server and serves on localhost at the
// specified port using the provided handler. The server will exit when the
// timeout expires.
func RunTimedHTTPHandler(port int, timeout time.Duration, handler http.Handler) {
	mux := http.NewServeMux()
	idleConnsClosed := make(chan struct{})

	// kill server on request
	mux.HandleFunc("/api/killserver", func(w http.ResponseWriter, r *http.Request) {
		close(idleConnsClosed)
	})

	mux.Handle("/api/project_analyses/search", handler)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: mux,
	}

	go func() {
		<-idleConnsClosed

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	go func() {
		time.Sleep(timeout)
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
