package config

import (
	"net/http"

	"github.com/Zarar-Azeem/golang-chi/pkg/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{
		addr: addr,
	}
}

func (s *ApiServer) Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	routes.RegisterBookStoreRoutes(r)
	server := http.Server{
		Addr:    s.addr,
		Handler: r,
	}
	return server.ListenAndServe()
}
