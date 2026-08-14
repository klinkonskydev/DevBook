package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

// User represents a person using the social network
type User struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"createdAt"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

// GetFullUser makes 4 requests to the API to assemble the user
func GetFullUser(userID uint32, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	postsChannel := make(chan []Post)

	go GetUserData(userChannel, userID, r)
	go GetFollowers(followersChannel, userID, r)
	go GetFollowing(followingChannel, userID, r)
	go GetUserPosts(postsChannel, userID, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case loadedUser := <-userChannel:
			if loadedUser.ID == 0 {
				return User{}, errors.New("error fetching the user")
			}

			user = loadedUser

		case loadedFollowers := <-followersChannel:
			if loadedFollowers == nil {
				return User{}, errors.New("error fetching the followers")
			}

			followers = loadedFollowers

		case loadedFollowing := <-followingChannel:
			if loadedFollowing == nil {
				return User{}, errors.New("error fetching who the user is following")
			}

			following = loadedFollowing

		case loadedPosts := <-postsChannel:
			if loadedPosts == nil {
				return User{}, errors.New("error fetching the posts")
			}

			posts = loadedPosts
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

// GetUserData calls the API to fetch the user's base data
func GetUserData(channel chan<- User, userID uint32, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

// GetFollowers calls the API to fetch the user's followers
func GetFollowers(channel chan<- []User, userID uint32, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

// GetFollowing calls the API to fetch the users followed by an user
func GetFollowing(channel chan<- []User, userID uint32, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- following
}

// GetUserPosts calls the API to fetch the posts made by an user
func GetUserPosts(channel chan<- []Post, userID uint32, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.APIURL, userID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		channel <- nil
		return
	}

	if posts == nil {
		channel <- make([]Post, 0)
		return
	}

	channel <- posts
}
