package app

import (
	"ca_kobercams/config"
	v1 "ca_kobercams/internal/controller/http/v1"
	user2 "ca_kobercams/internal/controller/user"
	"ca_kobercams/internal/entity"
	"ca_kobercams/internal/storage/mock"
	"ca_kobercams/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Run(cfg *config.Config) error {

	err, server := setUpDependencies(cfg)
	if err != nil {
		return err
	}

	server.ListenAndServe()
	return nil
}

func setUpDependencies(cfg *config.Config) (error, *http.Server) {

	handler := gin.New()

	server := &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//Set Up Storages
	userRepository := &mock.MockUserRepository{
		Users: make([]entity.User, 0),
	}

	//Set Up Usecases
	userUseCase := &user.User{
		Repo: userRepository,
	}

	userController := user2.UserController{
		userUseCase,
	}

	//Set Up Router
	router := v1.Router{
		userController,
		handler,
	}

	router.ConfigureRouter()
	return nil, server
}
