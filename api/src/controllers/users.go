package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"

	"github.com/gorilla/mux"
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

	if err = user.Prepare("register"); err != nil {
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
	params := mux.Vars(r)

	userIDu64, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDFromToken, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	userID := uint32(userIDu64)
	if userIDFromToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("you can't find another user that is not yours"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.UsersRepository(db)
	user, err := repository.GetUserByID(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// EditUser edit user by id from users table
func EditUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDu64, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDFromToken, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	userID := uint32(userIDu64)

	if userIDFromToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("you can't update another user that is not yours"))
		return
	}

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

	if err = user.Prepare("update"); err != nil {
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
	if err = repository.EditUser(userID, user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser remove an user in the users table
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDu64, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDFromToken, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	userID := uint32(userIDu64)
	if userIDFromToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("you can't delete another user that is not yours"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.UsersRepository(db)
	if err = repository.DeleteUser(userID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FollowUser follows an user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userIDu64, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userID := uint32(userIDu64)

	if userID == followerID {
		responses.Error(w, http.StatusForbidden, errors.New("isn't possible follow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.UsersRepository(db)
	if err := repository.FollowUser(userID, followerID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, nil)
}

// UpdatePassword will change the user password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIDFromToken, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userIDu64, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userID := uint32(userIDu64)

	if userIDFromToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("isn't possible update an user that is not yours"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	var password models.Password
	if err = json.Unmarshal(requestBody, &password); err != nil {
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
	savedPasswordFromDatabase, err := repository.GetPasswordFromUserID(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(savedPasswordFromDatabase, password.Old); err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("the old password is invalid"))
		return
	}

	newHashedPassword, err := security.Hash(password.New)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.UpdatePassword(userID, string(newHashedPassword)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUser unfollows an user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userIDu64, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userID := uint32(userIDu64)

	if userID == followerID {
		responses.Error(w, http.StatusForbidden, errors.New("isn't possible unfollow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.UsersRepository(db)
	if err := repository.UnfollowUser(userID, followerID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Followers show all user followers
func Followers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDu64, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userID := uint32(userIDu64)

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.UsersRepository(db)
	followers, err := repository.Followers(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

// Following show all user following
func Following(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userIDu64, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userID := uint32(userIDu64)

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.UsersRepository(db)
	following, err := repository.Following(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, following)
}
