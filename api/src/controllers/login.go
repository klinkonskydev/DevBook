package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

// Login is responsible for authenticate an user
func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.UsersRepository(db)
	savedUserFromDatabse, err := repository.GetUserByEmail(&user.Email, &user.Password)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(savedUserFromDatabse.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("invalid password"))
		return
	}

	token, err := authentication.CreateToken(savedUserFromDatabse.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userID := strconv.Itoa(int(savedUserFromDatabse.ID))
	responses.JSON(w, http.StatusOK, models.Authentication{ ID: userID, Token: token })
}
