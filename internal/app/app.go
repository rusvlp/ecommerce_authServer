package app

import (
	"ca_kobercams/config"
	v1 "ca_kobercams/internal/controller/http/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Run(cfg *config.Config) error {
	handler := gin.New()
	v1.NewRouter(handler)

	server := &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
