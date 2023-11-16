package handler

import (
	"diary_api/helper"
	"diary_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register: validates JSON request, creates new user,
// and writes details of saved user to JSON response.
func Register(context *gin.Context) {

	// define expected authentication input from request body
	var authInput model.AuthenticationInput

	// unmarshal request body into expected authentication input
	err := context.ShouldBindJSON(&authInput)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// instantiate user
	user := model.User{
		Username: authInput.Username,
		Password: authInput.Password,
	}

	// ! should user.beforeSave be called here?

	// save user to database
	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// write saved user to response body
	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

// Login: validates request, locates user if exists, validates password, generates JWT and writes the token to response body.
func Login(context *gin.Context) {

	//define expected authentication input from request body
	var authInput model.AuthenticationInput

	// unmarshal request body into expected input
	err := context.ShouldBindJSON(&authInput)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// locate existing user by username
	user, err := model.FindUserByUsername(authInput.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate user password
	err = user.ValidatePassword(authInput.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// generate JWT based on the user attempting to signin
	jwt, err := helper.GenerateJWT(user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// write jwt to response body
	context.JSON(http.StatusOK, gin.H{"jwt": jwt})

}