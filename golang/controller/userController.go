package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sanbertutor/models"
	"sanbertutor/repository"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(c echo.Context) error {
	user := models.User{}

	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Printf("Error read request body with err: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error read request body",
			"success": false,
		})
	}
	defer c.Request().Body.Close()

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password with err: %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error hash",
			"success": false,
		})
	}
	user.Password = string(password)

	user.Guid_user = uuid.NewString()
	user.Date_joined = time.Now()

	guid, err := repository.AddUserRepository(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error register user",
			"success": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("Guid %s was added to the database\n", guid),
		"success": true,
	})
}

func LoginController(c echo.Context) error {
	input := models.LoginInput{}

	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		log.Printf("Error read request body with err: %s", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error read request body",
			"success": false,
		})
	}
	defer c.Request().Body.Close()

	user, err := repository.GetUserLoginByEmailRepository(input.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
			"success": false,
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		fmt.Printf("Password didn't match with err: %s\n", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Password didn't match",
			"success": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login Successful",
		"greet":   fmt.Sprintf("Hello, %s %s! Your email is %s", user.Firstname, user.Lastname, user.Email),
		"success": true,
	})
}
