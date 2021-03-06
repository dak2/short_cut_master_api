package controllers

import (
	entity "short_cut_master_api/src/entities"
	repository "short_cut_master_api/src/interfaces/database"
	usecase "short_cut_master_api/src/usecases"
	"github.com/labstack/echo"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUsersController(sqlHandler repository.SqlHandler) *UserController {
	return &UserController {
		Interactor: usecase.UserInteractor {
			UserRepository: &repository.UserRepository {
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c echo.Context) {
	u := entity.User{}
	c.Bind(&u)
	controller.Interactor.Add(u)
	createdUsers := controller.Interactor.GetInfo()
	c.JSON(201, createdUsers)
	return
}

func (controller *UserController) GetUser() []entity.User {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}
