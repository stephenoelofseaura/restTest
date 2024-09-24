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

func ServerPutHandler(w http.ResponseWriter, r *http.Request, s *SocialMediaServer) error {
	//Update a Single User
	if strings.HasPrefix(r.URL.Path, "/Users/") {
		pathWithoutPrefix := strings.TrimPrefix(r.URL.Path, "/Users/")
		if pathWithoutPrefix == "" {
			return errors.New("Invalid URL")
		} else {
			fmt.Println("Updating User")
			userName := strings.TrimSuffix(pathWithoutPrefix, "/")
			body := r.Body
			var updatedUser User
			err := json.NewDecoder(body).Decode(&updatedUser)
			if err != nil {
				return errors.New("Error in request body. Error: " + err.Error())
			}
			userUpdateErr := s.database.UpdateUser(userName, updatedUser)
			if userUpdateErr != nil {
				return errors.New("Error updating user. Error: " + userUpdateErr.Error())
			} else {
				convertErr := db.ConvertAndWriteData(updatedUser, w)
				if convertErr != nil {
					return errors.New("Error converting data. Error: " + convertErr.Error())
				}
			}
			return nil
		}
	}
	if strings.HasPrefix(r.URL.Path, "/Posts/") {
		pathWithoutPrefix := strings.TrimPrefix(r.URL.Path, "/Posts/")
		if pathWithoutPrefix == "" {
			return errors.New("Invalid URL")
		} else {
			fmt.Println("Updating Post")
			body := r.Body
			var updatedPost Post
			err := json.NewDecoder(body).Decode(&updatedPost)
			if err != nil {
				return errors.New("Error reading request body")
			}
			postNumber := strings.TrimSuffix(pathWithoutPrefix, "/")
			postNumber := strconv.Atoi(postNumber)
			postUpdateErr := s.database.UpdatePost(postNumber, updatedPost)
			if postUpdateErr != nil {
				return errors.New("Error updating post. Error: " + postUpdateErr)
			} else {
				convertErr := db.ConvertAndWriteData(updatedPost, w)
				if convertErr != nil {
					return errors.New("Error converting data. Error: " + convertErr.Error())
				}
			}
		}
		return nil
	}
	return errors.New("Invalid URL Path")
}
