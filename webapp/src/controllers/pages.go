package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// LoadLoginScreen renders the login screen
func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoadUserRegistrationPage loads the user registration page
func LoadUserRegistrationPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

// LoadMainPage loads the main page with the posts
func LoadMainPage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeError(w, response)
		return
	}

	var posts []models.Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.APIError{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 32)

	utils.ExecuteTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserID uint32
	}{
		Posts:  posts,
		UserID: uint32(userID),
	})
}

// LoadEditPostPage loads the post edit page
func LoadEditPostPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 32)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.APIURL, postID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeError(w, response)
		return
	}

	var post models.Post
	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.APIError{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "edit-post.html", post)
}

// LoadUsersPage loads the page with the users matching the given filter
func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNickname := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/users?q=%s", config.APIURL, nameOrNickname)

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleStatusCodeError(w, response)
		return
	}

	var users []models.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.APIError{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "users.html", users)
}

// LoadUserProfilePage loads the user's profile page
func LoadUserProfilePage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 32)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	loggedUserID, _ := strconv.ParseUint(cookie["id"], 10, 32)

	if uint32(userID) == uint32(loggedUserID) {
		http.Redirect(w, r, "/profile", 302)
		return
	}

	user, err := models.GetFullUser(uint32(userID), r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "user.html", struct {
		User         models.User
		LoggedUserID uint32
	}{
		User:         user,
		LoggedUserID: uint32(loggedUserID),
	})
}

// LoadLoggedUserProfilePage loads the logged user's profile page
func LoadLoggedUserProfilePage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 32)

	user, err := models.GetFullUser(uint32(userID), r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "profile.html", user)
}

// LoadEditUserPage loads the page to edit the user's data
func LoadEditUserPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 32)

	channel := make(chan models.User)
	go models.GetUserData(channel, uint32(userID), r)
	user := <-channel

	if user.ID == 0 {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: "error fetching the user"})
		return
	}

	utils.ExecuteTemplate(w, "edit-user.html", user)
}

// LoadUpdatePasswordPage loads the page to update the user's password
func LoadUpdatePasswordPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "update-password.html", nil)
}
