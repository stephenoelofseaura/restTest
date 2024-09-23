package main

import (
	"errors"
	"fmt"
	"regexp"
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

func (s *StubDatabase) UpdateUser(name string, updatedUser User) error {
	fmt.Println(name)
	re := regexp.MustCompile("([A-Z][a-z0-9]+)")
	dataBaseFriendlyName := strings.Join(re.FindAllString(name, -1), " ")
	for i, user := range s.users {
		// fmt.Println(dataBaseFriendlyName)
		// fmt.Println(user)
		if user.Name == dataBaseFriendlyName {
			s.users[i] = updatedUser
			return nil
		}
	}
	return errors.New("User with this name does not exist, could not update user")
}

func (s *StubDatabase) GetPosts() []Post {
	return s.posts
}

func (s *StubDatabase) GetPostsByUser(name string) []Post {
	var posts []Post
	re := regexp.MustCompile("([A-Z][a-z0-9]+)")
	dataBaseFriendlyName := strings.Join(re.FindAllString(name, -1), " ")
	for _, post := range s.posts {
		if post.User.Name == dataBaseFriendlyName {
			posts = append(posts, post)
		}
	}
	return posts
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
