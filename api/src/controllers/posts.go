package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePost create a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userIDFromToken, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = uint32(userIDFromToken)

	if err = post.Prepare("create"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.PostsRepository(db)
	post.ID, err = repository.Create(post)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

// GetPosts get all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.PostsRepository(db)
	posts, err := repository.Get(uint32(userID))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

// GetPost get specific post
func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.PostsRepository(db)
	post, err := repository.GetByID(uint32(postID))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, post)
}

// EditPost edit a post
func EditPost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.PostsRepository(db)
	savedPostInTheDatabase, err := repository.GetByID(uint32(postID))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if savedPostInTheDatabase.AuthorID != uint32(userID) {
		responses.Error(w, http.StatusForbidden, errors.New("isn't possible update a post that is not yours"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = post.Prepare("update"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.Update(uint32(postID), post); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeletePost delete a post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.PostsRepository(db)
	savedPostInTheDatabase, err := repository.GetByID(uint32(postID))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if savedPostInTheDatabase.AuthorID != uint32(userID) {
		responses.Error(w, http.StatusForbidden, errors.New("isn't possible delete a post that is not yours"))
		return
	}

	if err = repository.Delete(uint32(postID)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// GetUserPosts get posts from an user account
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID , err := strconv.ParseUint(params["userId"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.PostsRepository(db)
	posts, err := repository.GetPostsByUser(uint32(userID))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

// Upvote upvotes an user post
func Upvote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.PostsRepository(db)
	if err = repository.Upvote(uint32(postID)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Downvote downvotes an user post
func Downvote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.PostsRepository(db)
	if err = repository.Downvote(uint32(postID)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
