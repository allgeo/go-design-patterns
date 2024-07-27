package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

const port = ":4000"

type application struct {
	// storing template
	templateMap map[string]*template.Template
	config      appConfig
}

type appConfig struct {
	useCache bool
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.Parse()

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting server on ", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
