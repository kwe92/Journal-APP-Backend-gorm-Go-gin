package handler

import (
	"errors"
	"fmt"
	"journal_api/database"
	"journal_api/model"
	"journal_api/utility"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register: validates JSON request, creates new user,
// and generates a jwt for registered user and writes jwt to JSON response body.
func Register(ctx *gin.Context) {

	// expected authentication input from request body
	var registrationInput model.RegistrationInput

	// unmarshal request body into expected input
	err := ctx.ShouldBindJSON(&registrationInput)

	if err != nil {
		utility.SendBadRequestResponse(ctx, err)
		return
	}

	// instantiate user
	user := model.User{
		Fname:    registrationInput.Fname,
		Lname:    registrationInput.Lname,
		Email:    registrationInput.Email,
		Phone:    registrationInput.Phone,
		Password: registrationInput.Password,
	}

	// save user to database
	savedUser, err := user.Save(database.Database)

	if err != nil {
		utility.SendBadRequestResponse(ctx, err)
		return
	}

	// generate JWT based on newlyregistered user
	jwt, err := utility.GenerateJWT(*savedUser)

	if err != nil {
		utility.SendBadRequestResponse(ctx, err)
		return
	}

	// write jwt and user to response body
	ctx.JSON(http.StatusOK, gin.H{
		"jwt": jwt,
		"user": gin.H{
			"first_name":   &savedUser.Fname,
			"last_name":    &savedUser.Lname,
			"email":        &savedUser.Email,
			"phone_number": &savedUser.Phone,
		},
	})

	log.Printf("\nnew user registration: %+v\n\n", *savedUser)
}

// Login: validates request, locates user if exists, validates password, generates JWT and writes the token to response body.
func Login(ctx *gin.Context) {

	// define expected authentication input from request body
	var loginInput model.LoginInput

	// unmarshal request body into expected input
	err := ctx.ShouldBindJSON(&loginInput)

	fmt.Println("\n\nAUTH Input:", loginInput)

	fmt.Println("\n\nAUTH Email:", loginInput.Email)

	fmt.Println("\n\nAUTH Password:", loginInput.Password)

	if err != nil {
		utility.SendBadRequestResponse(ctx, err)
		return
	}
	// locate existing user by email
	user, err := model.FindUserByEmail(loginInput.Email, database.Database)

	if err != nil {
		utility.SendBadRequestResponse(ctx, err)
		return
	}

	// validate user password
	err = user.ValidatePassword(loginInput.Password)

	if err != nil {
		utility.SendBadRequestResponse(ctx, err)
		return
	}
	// generate JWT based on the user attempting to signin
	jwt, err := utility.GenerateJWT(user)

	if err != nil {
		utility.SendBadRequestResponse(ctx, err)
		return
	}

	// write jwt and user to response body
	ctx.JSON(http.StatusOK, gin.H{
		"jwt": jwt,
		"user": gin.H{
			"first_name":   user.Fname,
			"last_name":    user.Lname,
			"email":        user.Email,
			"phone_number": user.Phone,
		}})

}

// CheckAvailableEmail: checks database for available email
func CheckAvailableEmail(ctx *gin.Context) {
	var userEmail model.UserEmail

	if err := ctx.ShouldBindJSON(&userEmail); err != nil {
		utility.SendBadRequestResponse(ctx, err)

	}

	user, err := model.FindUserByEmail(userEmail.Email, database.Database)

	// return error if a user was found
	if err == nil {
		utility.SendBadRequestResponse(ctx, errors.New("user already exists"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": fmt.Sprintf("email %s is available", userEmail.Email)})

	fmt.Println("user from CheckAvailableEmail:", user)

}
