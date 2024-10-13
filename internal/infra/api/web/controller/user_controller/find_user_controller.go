package user_controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"leilao/configuration/rest_err"
	"leilao/internal/usecase/user_usecase"
	"net/http"
)

type userController struct {
	userUseCase user_usecase.UserUsecase
}

func NewUserController(userUseCase user_usecase.UserUsecase) *userController {
	return &userController{
		userUseCase: userUseCase,
	}
}

func (u *userController) FindUserById(c *gin.Context) {
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
