package handlers

import (
	"MyGramProject/entities"
	"MyGramProject/helpers"
	"MyGramProject/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *userHandler {
	return &userHandler{service}
}

func (handler *userHandler) RegisterUser(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	// Validasi email
	if !helpers.ValidateEmail(user.Email) {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid email format")
		return
	}

	// Call service to register user
	if err := handler.service.RegisterUser(&user); err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to register user")
		return
	}

	helpers.SuccessResponse(c, http.StatusCreated, user)
}

func (handler *userHandler) LoginUser(c *gin.Context) {
	// Parse request body
	var loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authenticate user
	token, err := handler.service.LoginUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// If authenticated, return JWT token
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (handler *userHandler) UpdateUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userId"))
	var updateUser entities.User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid input data")
		return
	}

	err := handler.service.UpdateUser(uint(userID), &updateUser)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to update user")
		return
	}

	helpers.SuccessResponse(c, http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (handler *userHandler) DeleteUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userId"))

	err := handler.service.DeleteUser(uint(userID))
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	helpers.SuccessResponse(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}
