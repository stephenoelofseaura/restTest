package main

import (
	"errors"
	"strings"
)

type StubDatabase struct {
	users []User
	posts []Post
}

func (s *StubDatabase) GetUsers() []User {
	return s.users
}

func (s *StubDatabase) CreateUser(newUser User) error {
	for _, user := range s.users {
		if user.Name == newUser.Name {
			return errors.New("User with this name already exists")
		}
	}
	s.users = append(s.users, newUser)
	return nil
}

func (s *StubDatabase) GetPosts() []Post {
	return s.posts
}

func (s *StubDatabase) CreatePost(newPost Post) error {
	for _, post := range s.posts {
		if post.Id == newPost.Id {
			return errors.New("Post with this Id already exists")
		}
	}
	s.posts = append(s.posts, newPost)
	return nil
}

func (s *StubDatabase) GetPost(Id int) Post {
	for _, post := range s.posts {
		if post.Id == Id {
			return post
		}
	}
	return Post{}
}
func (s *StubDatabase) GetUser(name string) User {
	for _, user := range s.users {
		newName := strings.Replace(user.Name, " ", "", -1)
		if newName == name {
			return user
		}
	}
	return User{}
}

func (s *StubDatabase) CreateComment(comment Comment, postId int) error {
	for _, post := range s.posts {
		if post.Id == postId {
			post.Comments = append(post.Comments, comment)
			return nil
		}
	}
	return errors.New("Post with this Id does not exist, could not create comment")
}
