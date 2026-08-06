package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
)

// CreateUser insert a new user within users table
func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if err = user.Prepare(); err != nil {
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
	user.ID, err = repository.CreateUser(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, &user)
}

// GetUsers get all users from users table
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNickname := strings.ToLower(r.URL.Query().Get("q"))

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.UsersRepository(db)
	users, err := repository.GetUsers(nameOrNickname)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, &users)
}

// GetUser get user by id from users table
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user"))
}

// EditUser edit user by id from users table
func EditUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Edit user"))
}

// DeleteUser remove an user in the users table
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
