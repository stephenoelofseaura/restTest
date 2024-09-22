package main

import (
	"errors"
	"fmt"
	"github.com/username/rest-test/pkg/db"
	"net/http"
	"strconv"
	"strings"
)

func serverGetHandler(w http.ResponseWriter, r *http.Request, s *SocialMediaServer) error {
	if strings.HasPrefix(r.URL.Path, "/Users/") {
		user := strings.TrimPrefix(r.URL.Path, "/Users/")
		if user == "" {
			fmt.Println("Getting Users")
			err := db.convertAndWriteData(s.database.GetUsers(), w)
			if err != nil {
				return errors.New("Error converting data. Error: " + err.Error())
			}
			return nil
		} else {
			user = strings.TrimSuffix(user, "/")
			fmt.Println("Getting User: " + user)
			err := db.convertAndWriteData(s.database.GetUser(user), w)
			if err != nil {
				return errors.New("Error converting data. Error: " + err.Error())
			}
			return nil
		}
	}

	if strings.HasPrefix(r.URL.Path, "/Posts/") {
		post := strings.TrimPrefix(r.URL.Path, "/Posts/")
		if post == "" {
			fmt.Println("Getting Posts")
			err := db.convertAndWriteData(s.database.GetPosts(), w)
			if err != nil {
				return errors.New("Error converting data. Error: " + err.Error())
			}
			return nil
		} else {
			fmt.Println("Getting Post: " + post)
			postNum, err := strconv.Atoi(post)
			if err != nil {
				return errors.New("Error converting post number. Error: " + err.Error())
			}
			err = db.convertAndWriteData(s.database.GetPost(postNum), w)
			if err != nil {
				return errors.New("Error converting data. Error: " + err.Error())
			}
			return nil
		}
	}
	return errors.New("Invalid URL Path")
}
