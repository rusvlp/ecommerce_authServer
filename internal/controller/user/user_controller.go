package user

import (
	"ca_kobercams/internal/entity"
	"ca_kobercams/internal/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserUsecase usecase.UserInterface
}

type CreateUserRequest struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
	Email                string `json:"email"`
}

func (r *CreateUserRequest) String() string {
	return "User " + r.Username + ", password: " + r.Password + ", passwordConf: " + r.PasswordConfirmation + ", email: " + r.Email
}

func (controller *UserController) CreateUserEndpoint(context *gin.Context) {
	createUserRequest := &CreateUserRequest{}

	if err := context.ShouldBindJSON(createUserRequest); err != nil {
		context.String(http.StatusBadRequest, err.Error())
	}

	fmt.Println("Got a request for creating user" + createUserRequest.String())

	user := entity.User{
		Username: createUserRequest.Username,
		Password: createUserRequest.Password,
		Email:    createUserRequest.Email,
	}

	if err := controller.UserUsecase.CreateUser(&user, createUserRequest.PasswordConfirmation); err != nil {
		context.String(http.StatusBadRequest, err.Error())
	}

}

func (controller *UserController) FindAllEndpoin(context *gin.Context) {
	err, users := controller.UserUsecase.FindAll()
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
	}

	//jsonResponse, err := json.Marshal(users)

	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
	}

	//fmt.Println(string(jsonResponse))

	context.JSON(http.StatusOK, users)
}
