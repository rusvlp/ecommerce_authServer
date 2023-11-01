package app

import (
	"ca_kobercams/config"
	v1 "ca_kobercams/internal/controller/http/v1"
	user2 "ca_kobercams/internal/controller/user"
	mysql3 "ca_kobercams/internal/storage/mysql"
	"ca_kobercams/internal/usecase/user"
	mysql2 "ca_kobercams/pkg/mysql"
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

	err, mysql := mysql2.InitMysql(cfg.Database)

	if err != nil {
		return err, nil
	}

	err, repos := mysql3.InitMysqlRepositories(mysql.Database)

	if err != nil {
		return err, nil
	}

	//Set Up Usecases
	userUseCase := &user.User{
		Repo: repos.UserRepository,
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
