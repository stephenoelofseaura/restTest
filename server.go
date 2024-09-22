package main

//TODO
//Read All Posts by User

//Update User
//Update Comment
//Update Post
//Update Likes

//Delete User
//Delete Comment
//Delete Post

import (
	"errors"
	"net/http"
)

type User struct {
	Name string `json:"Name"`
}

type Comment struct {
	Text string `json:"Text"`
	User User   `json:"User"`
	Id   int    `json:"Id"`
}

type Comments []Comment

type Post struct {
	Text     string   `json:"Text"`
	User     User     `json:"User"`
	Id       int      `json:"Id"`
	Likes    int      `json:"Likes"`
	Comments Comments `json:"Comments"`
}

type Database interface {
	GetUsers() []User
	GetUser(name string) User
	CreateUser(user User) error

	CreateComment(comment Comment, postId int) error

	GetPosts() []Post
	GetPost(id int) Post
	CreatePost(post Post) error
}

type SocialMediaServer struct {
	database Database
}

func (s *SocialMediaServer) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {

	//Read Single User
	//Read All Users
	//Read Post by Id
	//Read All Posts
	case http.MethodGet:
		return serverGetHandler(w, r, s)

	//Create User
	//Create Post
	case http.MethodPost:
		return serverPostHandler(w, r, s)
	}
	return errors.New("Unsupported method")
}
