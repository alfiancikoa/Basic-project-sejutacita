package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	response "rest-example/api/common"
	"rest-example/api/middlewares"
	"rest-example/models"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	userModel models.UserModel
}

func NewController(userModel models.UserModel) *Controller {
	return &Controller{
		userModel,
	}
}

// Fungsi Register User
func (controller Controller) RegisterController(c echo.Context) error {
	newUser := models.User{}
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, response.StatusFailed("bad request"))
	}
	// Regex
	var pattern string
	var matched bool
	// Check Format Name
	pattern = `^(\w+ ?){4}$`
	regex, _ := regexp.Compile(pattern)
	matched = regex.Match([]byte(newUser.Username))
	if !matched {
		return c.JSON(http.StatusBadRequest, response.StatusFailed("username cannot less than 5 characters or invalid format"))
	}
	// Check Format Password
	pattern = `^([a-zA-Z0-9()@:%_\+.~#?&//=\n"'\t\\;<>!$*-{}]+ ?){8}$`
	matched, _ = regexp.Match(pattern, []byte(newUser.Password))
	if !matched {
		return c.JSON(http.StatusBadRequest, response.StatusFailed("password must contain password format and more than equals 8 characters"))
	}
	// Check Format Email
	emailLower := strings.ToLower(newUser.Email)
	pattern = `^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$`
	matched, _ = regexp.Match(pattern, []byte(emailLower))
	if !matched {
		return c.JSON(http.StatusBadRequest, response.StatusFailed("email must contain email format"))
	}
	fmt.Println(newUser)
	rowName, _ := controller.userModel.CheckDatabase("username", newUser.Username)
	if rowName > 0 {
		return c.JSON(http.StatusBadRequest, response.StatusFailed("username is already used"))
	}
	rowEmail, _ := controller.userModel.CheckDatabase("email", newUser.Email)
	if rowEmail > 0 {
		return c.JSON(http.StatusBadRequest, response.StatusFailed("email is already used"))
	}
	newUser.Role = "user"
	newUser.Password, _ = models.GeneratehashPassword(newUser.Password)
	if _, err := controller.userModel.Insert(newUser); err != nil {
		return c.JSON(http.StatusInternalServerError, response.StatusFailed("internal server error"))
	}
	return c.JSON(http.StatusCreated, response.StatusSuccess("Success Registered New Account"))
}

// Fungsi untuk Login User
func (controller Controller) LoginController(c echo.Context) error {
	dataLogin := models.LoginUser{}
	// Bind all data from JSON
	c.Bind(&dataLogin)
	dataLogin.Username = strings.ToLower(dataLogin.Username)
	user, err := controller.userModel.Login(dataLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.StatusFailed("internal server error"))
	}
	if user == nil {
		return c.JSON(http.StatusBadRequest, response.StatusFailed("invalid username or password"))
	}
	token, _ := middlewares.CreateToken(user.ID, user.Role)
	return c.JSON(http.StatusCreated, response.StatusSuccessLogin("login success", user.ID, token))
}

// Fungsi Edit User
func (controller Controller) UpdateController(c echo.Context) error {
	newDataUser := models.User{}
	if err := c.Bind(&newDataUser); err != nil {
		return c.JSON(http.StatusBadRequest, response.StatusFailed("bad request"))
	}
	idlogin, _ := middlewares.ExtractTokenId(c)
	userData, _ := controller.userModel.FindUserBy("id", idlogin)
	// Regex
	var pattern string
	var matched bool
	// Check Format Name
	pattern = `^(\w+ ?){4}$`
	regex, _ := regexp.Compile(pattern)
	matched = regex.Match([]byte(newDataUser.Username))
	if !matched {
		return c.JSON(http.StatusBadRequest, response.StatusFailed("username cannot less than 5 characters or invalid format"))
	}

	// Check Format Email
	emailLower := strings.ToLower(newDataUser.Email)
	pattern = `^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$`
	matched, _ = regexp.Match(pattern, []byte(emailLower))
	if !matched {
		return c.JSON(http.StatusBadRequest, response.StatusFailed("email must contain email format"))
	}
	// Cek apakah username sudah digunakan atau belum
	if userData.Username != newDataUser.Username {
		rowName, _ := controller.userModel.CheckDatabase("username", newDataUser.Username)
		if rowName > 0 {
			return c.JSON(http.StatusBadRequest, response.StatusFailed("username is already used try another one"))
		}
	}
	// Cek email apakahh sudah digunakan atau belum
	if userData.Email != newDataUser.Email {
		rowEmail, _ := controller.userModel.CheckDatabase("email", newDataUser.Email)
		if rowEmail > 0 {
			return c.JSON(http.StatusBadRequest, response.StatusFailed("email is already used try another one"))
		}
	}
	// Cek Password
	if newDataUser.Password == "" {
		// Gunakan password lama
		newDataUser.Password = userData.Password
	} else {
		// Check Format Password
		pattern = `^([a-zA-Z0-9()@:%_\+.~#?&//=\n"'\t\\;<>!$*-{}]+ ?){8}$`
		matched, _ = regexp.Match(pattern, []byte(newDataUser.Password))
		if !matched {
			return c.JSON(http.StatusBadRequest, response.StatusFailed("password must contain password format and more than equals 8 characters"))
		}
		// Bycrypt new Password
		newDataUser.Password, _ = models.GeneratehashPassword(newDataUser.Password)
	}
	if _, err := controller.userModel.Edit(newDataUser, idlogin); err != nil {
		return c.JSON(http.StatusInternalServerError, response.StatusFailed("internal server error"))
	}
	return c.JSON(http.StatusCreated, response.StatusSuccess("Success Edit Profile"))

}

// Fungsi untuk menghapus user khusus admin
func (controller Controller) DeleteController(c echo.Context) error {
	idlogin, role := middlewares.ExtractTokenId(c)
	user_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.StatusFailed("false param"))
	}
	if role == "admin" || idlogin == user_id {
		if _, err := controller.userModel.Delete(user_id); err != nil {
			return c.JSON(http.StatusInternalServerError, response.StatusFailed("internal server error"))
		}
		return c.JSON(http.StatusOK, response.StatusSuccess("success delete user"))
	} else {
		return c.JSON(http.StatusUnauthorized, response.StatusFailed("unauthorized access"))
	}
}

// Fungsi untuk melihat profil untuk user, dan melihat seluruh data khusus admin
func (controller Controller) GetUserController(c echo.Context) error {
	user_id, role := middlewares.ExtractTokenId(c)
	if role == "admin" {
		users, err := controller.userModel.FindUsers()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.StatusFailed("internal server error"))
		}
		return c.JSON(http.StatusOK, response.StatusSuccessData("success get all user", users))
	}
	user, err := controller.userModel.FindUserBy("id", user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.StatusFailed("internal server error"))
	}
	return c.JSON(http.StatusOK, response.StatusSuccessData("success get profile user", user))
}
