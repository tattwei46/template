package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/gin-gonic/gin"
)

func main() {
	setupConfig()
	setupLogger()
	r := setupRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    getServiceAddr(),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func getServiceAddr() string {
	if service, ok := serviceInfo[serviceKey]; ok {
		return fmt.Sprintf("%s:%d", service.Host, service.Port)
	}
	return fmt.Sprintf("0.0.0.0:15777")
}

func setupLogger() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/v1").Subrouter()
	s.HandleFunc("/", health).Methods(http.MethodGet)
	s.HandleFunc("/", add).Methods(http.MethodPost)
	return r
}
