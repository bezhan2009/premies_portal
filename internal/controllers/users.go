package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/internal/controllers/middlewares"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
	"strconv"
)

func GetAllUsers(c *gin.Context) {
	afterIDStr := c.Query("after")
	if afterIDStr == "" {
		afterIDStr = "0"
	}

	afterID, err := strconv.Atoi(afterIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	users, err := service.GetAllUsers(uint(afterID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] invalid user_id path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		logger.Error.Printf("[controllers.GetUserByID] invalid id: %s\n", c.Param("id"))
		return
	}

	user, err := service.GetUserByID(uint(id))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] error: %v\n", err)
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetMyDataUser(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)

	user, err := service.GetUserByID(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Error.Printf("[controllers.SignUp] error: %v\n", err)
		return
	}

	_, err := service.SignUp(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logger.Error.Printf("[controllers.SignUp] error: %v\n", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})
	logger.Info.Printf("[controllers.SignUp] message successfully\n data %v", user)
}
