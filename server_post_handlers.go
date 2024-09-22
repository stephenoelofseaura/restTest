package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/username/rest-test/pkg/db"
)

func serverPostHandler(w http.ResponseWriter, r *http.Request, s *SocialMediaServer) error {
	if strings.HasPrefix(r.URL.Path, "/Users/") {
		fmt.Println("Creating User")
		body := r.Body
		var newUser User
		err := json.NewDecoder(body).Decode(&newUser)
		if err != nil {
			return errors.New("Error decoding user. Error: " + err.Error())
		}
		userCreationErr := s.database.CreateUser(newUser)
		if userCreationErr != nil {
			return errors.New("Error creating user. Error: " + userCreationErr.Error())
		} else {
			convertErr := db.ConvertAndWriteData(newUser, w)
			if convertErr != nil {
				return errors.New("Error converting data. Error: " + convertErr.Error())
			}
			return nil
		}
	}

	if strings.HasPrefix(r.URL.Path, "/Posts/") {
		pathWithoutPrefix := strings.TrimPrefix(r.URL.Path, "/Posts/")
		if pathWithoutPrefix == "" {
			fmt.Println("Creating Post")
			body := r.Body
			var newPost Post
			err := json.NewDecoder(body).Decode(&newPost)
			if err != nil {
				return errors.New("Error decoding user. Error: " + err.Error())
			}
			postCreationError := s.database.CreatePost(newPost)
			if postCreationError != nil {
				return errors.New("Error creating post. Error: " + postCreationError.Error())
			} else {
				convertErr := db.ConvertAndWriteData(newPost, w)
				if convertErr != nil {
					return errors.New("Error converting data. Error: " + convertErr.Error())
				}
			}
			return nil
		} else {
			if strings.HasSuffix(pathWithoutPrefix, "/Comments/") {
				postId := strings.TrimSuffix(pathWithoutPrefix, "/Comments/")
				fmt.Println("Creating Comment for Post: " + postId)
				body := r.Body
				var newComment Comment
				err := json.NewDecoder(body).Decode(&newComment)
				if err != nil {
					return errors.New("Error in Request Body. Error: " + err.Error())
				}
				postIdInt, conversionError := strconv.Atoi(postId)
				if conversionError != nil {
					return errors.New("Error getting Post Id. Error: " + conversionError.Error())
				}
				commentCreationError := s.database.CreateComment(newComment, postIdInt)
				if commentCreationError != nil {
					return errors.New("Database returned error creating comment. Error: " + commentCreationError.Error())
				} else {
					convertErr := db.ConvertAndWriteData(newComment, w)
					if convertErr != nil {
						return errors.New("Error converting data. Error: " + convertErr.Error())
					}
				}
				return nil
			}
		}
	}
	return errors.New("Invalid URL Path")
}
