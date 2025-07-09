package controllers

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"net/http"
	"os"
	"premiesPortal/internal/app/models"
	"premiesPortal/internal/app/service"
	"premiesPortal/pkg/errs"
	"premiesPortal/pkg/logger"
	utils2 "premiesPortal/pkg/utils"
	"time"
)

const emptyInt = 0

// GoogleLogin godoc
// @Summary Начать авторизацию через Google
// @Description Этот эндпоинт перенаправляет пользователя на страницу авторизации Google.
// @Tags auth
// @Produce  json
// @Success 302 {string} string "Redirect to Google"
// @Failure 500 {object} models.ErrorResponse
// @Router /auth/google [get]
func GoogleLogin(c *gin.Context) {
	c.Request.URL.Query().Set("provider", "google")
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

// GoogleCallback godoc
// @Summary Авторизация через Google с созданием пользователя
// @Description Этот эндпоинт авторизует пользователя через Google OAuth. Сразу вызывается регистрация пользователя (SignUp), и независимо от результата регистрации происходит вход (SignIn) для получения JWT токенов.
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} models.TokenResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /auth/google/callback [get]
func GoogleCallback(c *gin.Context) {
	// Устанавливаем провайдера "google" в контекст запроса
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", "google"))

	// Завершаем авторизацию через Google
	userData, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		logger.Error.Printf("Google auth error: %v", err)
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	logger.Info.Printf("Получены данные от Google: %+v", userData)

	// Формируем модель пользователя для регистрации/входа
	user := models.User{
		Email:    userData.Email,
		Username: userData.Name,
	}

	signedUserID, accessToken, refreshToken, err := service.SignInWithGoogle(user)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrYouHaveNotRegisteredYet)
			return
		}

		HandleError(c, err)
		return
	}

	// Возвращаем токены
	c.JSON(http.StatusOK, models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       signedUserID,
	})
}

// SignUp godoc
// @Summary Register a new user
// @Description This endpoint registers a new user with a username, email, and password.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.UserRequest true "User information"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/sign-up [post]
func SignUp(c *gin.Context) {
	var user models.UserRequest

	if err := c.BindJSON(&user); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	userID, err := service.SignUp(
		models.User{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
			Phone:    user.Phone,
			FullName: user.FullName,
			RoleID:   int(user.RoleID),
		},

		models.Worker{
			Salary:        user.Salary,
			Position:      user.Position,
			Plan:          user.Plan,
			SalaryProject: user.SalaryProject,
			PlaceWork:     user.PlaceWork,
		},
	)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrIncorrectUsernameOrPassword)
			return
		}

		HandleError(c, err)
		return
	}

	accessToken, refreshToken, err := utils2.GenerateToken(userID, user.RoleID, user.Username)
	if err != nil {
		logger.Error.Printf("Error generating access token: %s", err)
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       userID,
	})
}

// SignIn godoc
// @Summary User login
// @Description This endpoint logs in an existing user using their username, email, and password.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.UserLogin true "User login information"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/sign-in [post]
func SignIn(c *gin.Context) {
	var user models.UserLogin

	if err := c.BindJSON(&user); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if user.Username == "" {
		HandleError(c, errs.ErrUsernameIsEmpty)
		return
	}

	if user.Password == "" {
		HandleError(c, errs.ErrPasswordIsEmpty)
		return
	}

	user.Password = utils2.GenerateHash(user.Password)

	signedUser, accessToken, refreshToken, err := service.SignIn(user.Username, user.Password)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrIncorrectUsernameOrPassword)
			return
		}

		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       signedUser.ID,
		RoleID:       uint(signedUser.RoleID),
	})
}

// RefreshToken godoc
// @Summary Refresh Token
// @Description This endpoint refreshes the access token.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.RefreshRequest true "User login information"
// @Success 200 {object} models.RefreshTokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/refresh [post]
func RefreshToken(c *gin.Context) {
	var requestBody struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	// Проверка валидности refresh_token
	token, err := jwt.ParseWithClaims(requestBody.RefreshToken, &utils2.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		HandleError(c, errs.ErrInvalidToken)
		return
	}

	// Генерация нового access_token
	claims, ok := token.Claims.(*utils2.CustomClaims)
	if !ok || claims.ExpiresAt < time.Now().Unix() {
		HandleError(c, errs.ErrRefreshTokenExpired)
		return
	}

	accessToken, refreshToken, err := utils2.GenerateToken(claims.UserID, claims.RoleID, claims.Username)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}
