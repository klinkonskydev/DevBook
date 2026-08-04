package controllers

import (
	"net/http"
)

// CreateUser insert a new user within users table
func CreateUser(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Creating user"))
}

// GetUsers get all users from users table
func GetUsers(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Get all users"))
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
