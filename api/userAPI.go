package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DendiAnugerah/Todo/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (A *API) RegisterUser(c echo.Context) error {
	var user model.User

	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return c.JSON(400, model.ErrorResponse{Error: err.Error()})
	}

	if user.Username == "" || user.Password == "" {
		return c.JSON(400, model.ErrorResponse{Error: "Username or password cannot be empty!"})
	}

	if A.UserRepo.CheckPasswordLength(user.Password) {
		return c.JSON(400, model.ErrorResponse{Error: "Password length must be greater than 5!"})
	}

	err = A.UserRepo.AddUser(user)
	if err != nil {
		return c.JSON(500, model.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(200, model.SuccessResponse{Username: user.Username, Message: "User registered successfully!"})
}

func (A *API) Login(c echo.Context) error {
	var user model.User
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return c.JSON(400, model.ErrorResponse{Error: err.Error()})
	}

	if user.Username == "" || user.Password == "" {
		return c.JSON(400, model.ErrorResponse{Error: "Username or password cannot be empty!"})
	}

	err = A.UserRepo.IsUsernameAvail(user)
	if err != nil {
		return c.JSON(401, model.ErrorResponse{Error: "Wrong username or password!"})
	}

	sessionToken := uuid.New().String()
	Expiry := time.Now().Add(time.Hour * 3)
	session := model.Session{Token: sessionToken, Username: user.Username, Expiry: Expiry}

	_, err = A.SessionRepo.SessionNameAvail(session.Username)
	if err != nil {
		err = A.SessionRepo.AddSession(session)
		if err != nil {
			return c.JSON(500, model.ErrorResponse{Error: err.Error()})
		}
	} else {
		err = A.SessionRepo.UpdateSession(session)
		if err != nil {
			return c.JSON(500, model.ErrorResponse{Error: err.Error()})
		}
	}

	c.SetCookie(&http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: Expiry,
	})

	return c.JSON(200, model.SuccessResponse{Username: user.Username, Message: "Login successfully!"})
}
