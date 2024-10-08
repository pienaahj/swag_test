package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pienaahj/swag_test/model"
)

type userResponse struct {
	Data []model.User `json:"data"`
}

type errResponse struct {
	Message string `json:"message"`
}

// GetUsers return list of all users from the database
// @Summary return all users
// @Description return list of all users from the database
// @Tags Users
// @Router /users [get]
func GetUsers(c *gin.Context) {
	users := model.ListUsers()
	c.JSON(http.StatusOK, userResponse{Data: users})
}

// CreateUser creates a new user
// @Summary creates a new user
// @Description creates a new user
// @Security bearerToken
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "User"
// @Success 201 {object} model.User
// @Failure 400 {object} errResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	err := model.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	auth := c.GetHeader("Authorization")
	c.Header("Authorization", auth)
	log.Println("authorization", auth)

	c.JSON(http.StatusCreated, req)
}
