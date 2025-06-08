package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		HandleError(c, errs.ErrInvalidAfterID)
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
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] invalid user_id path parameter: %s\n", c.Param("id"))
		HandleError(c, errs.ErrInvalidID)
		return
	}

	roleID, err := service.GetRoleByUserID(uint(userID))
	if err != nil {
		HandleError(c, err)
		return
	}

	user, err := service.GetUserByID(uint(userID), roleID)
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] error: %v\n", err)
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetMyDataUser(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	roleID := c.GetUint(middlewares.UserRoleIDCtx)

	user, err := service.GetUserByID(userID, roleID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
