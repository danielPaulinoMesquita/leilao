package user_controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"leilao/configuration/rest_err"
	"leilao/internal/usecase/user_usecase"
	"net/http"
)

type UserController struct {
	userUseCase user_usecase.UserUseCaseInterface
}

func NewUserController(userUseCase user_usecase.UserUseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (u *UserController) FindUserById(c *gin.Context) {
	userId := c.Query("userId")

	if err := uuid.Validate(userId); err != nil {
		errReset := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "userId",
			Message: "Invalid UUID fields",
		})

		c.JSON(errReset.Code, errReset)
		return
	}

	userData, err := u.userUseCase.FindUserById(context.Background(), userId)
	if err != nil {
		restErr := rest_err.ConvertError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, userData)
}
