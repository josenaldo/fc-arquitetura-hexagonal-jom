package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/josenaldo/fc-arquitetura-hexagonal-jom/adapters/web/handler"
	"github.com/josenaldo/fc-arquitetura-hexagonal-jom/application"
	"github.com/urfave/negroni"
)

type WebServer struct {
	ProductService application.ProductServiceInterface
	Port           string
}

func NewWebServer(productService application.ProductServiceInterface, port string) *WebServer {
	return &WebServer{ProductService: productService, Port: port}
}

func (s *WebServer) Serve() {

	r := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	handler.NewProductHandler(r, n, s.ProductService)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              s.Port,
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := http.ListenAndServe(server.Addr, nil)
	if err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}
