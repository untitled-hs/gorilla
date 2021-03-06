package controller

import (
	"errors"

	"github.com/dj-hirrot/gorilla/src/domain/models"
	"github.com/dj-hirrot/gorilla/src/interface/db"
	"github.com/dj-hirrot/gorilla/src/usecase"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler db.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &db.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// Show godoc
// @Summary      Show an user
// @Description  Get user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string        true  "User ID"
// @Success      200  {object}  models.User
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/{id} [get]
func (controller *UserController) Show(c echo.Context) (err error) {
	id, _ := uuid.FromString(c.Param("id"))
	user, err := controller.Interactor.Show(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, NewError(err))
		return
	}
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

// Index godoc
// @Summary      List users
// @Description  Get all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Users
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users [get]
func (controller *UserController) Index(c echo.Context) (err error) {
	users, err := controller.Interactor.Index()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

// Create godoc
// @Summary      Create user
// @Description  Create user by body
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        parameter body      models.UserAttributes true "User attributes"
// @Success      201       {object}  models.User
// @Failure      400       {object}  Error
// @Failure      404       {object}  Error
// @Failure      500       {object}  Error
// @Router       /users [post]
func (controller *UserController) Create(c echo.Context) (err error) {
	u := models.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Create(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

// Update godoc
// @Summary      Update user
// @Description  Update user by body
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id        path      string                true "User ID"
// @Param        parameter body      models.UserAttributes true "User attributes"
// @Success      200       {object}  models.User
// @Failure      400       {object}  Error
// @Failure      404       {object}  Error
// @Failure      500       {object}  Error
// @Router       /users/{id} [put]
func (controller *UserController) Update(c echo.Context) (err error) {
	id, _ := uuid.FromString(c.Param("id"))
	u := models.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Update(id, u)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, NewError(err))
		return
	}
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}

// Delete godoc
// @Summary      Delete user
// @Description  Delete user by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      204  {object}  nil
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/{id} [delete]
func (controller *UserController) Delete(c echo.Context) (err error) {
	id, _ := uuid.FromString(c.Param("id"))
	err = controller.Interactor.Delete(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, NewError(err))
		return
	}
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.NoContent(204)
	return
}
