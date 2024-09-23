package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/username/rest-test/pkg/db"
)

func ServerGetHandler(w http.ResponseWriter, r *http.Request, s *SocialMediaServer) error {

	if strings.HasPrefix(r.URL.Path, "/Users/") {
		user := strings.TrimPrefix(r.URL.Path, "/Users/")
		//Get All Users
		if user == "" {
			fmt.Println("Getting Users")
			err := db.ConvertAndWriteData(s.database.GetUsers(), w)
			if err != nil {
				return errors.New("Error converting data. Error: " + err.Error())
			}
			return nil
		}
		//Get All Posts by User
		if strings.HasSuffix(user, "/Posts/") {
			user = strings.TrimSuffix(user, "/Posts/")
			fmt.Println("Getting Posts by User: " + user)
			err := db.ConvertAndWriteData(s.database.GetPostsByUser(user), w)
			if err != nil {
				return errors.New("Error converting data. Error: " + err.Error())
			}
			return nil
		}
		//Get User
		user = strings.TrimSuffix(user, "/")
		fmt.Println("Getting User: " + user)
		err := db.ConvertAndWriteData(s.database.GetUser(user), w)
		if err != nil {
			return errors.New("Error converting data. Error: " + err.Error())
		}
		//No Path exists
		return nil
	}

	if strings.HasPrefix(r.URL.Path, "/Posts/") {
		post := strings.TrimPrefix(r.URL.Path, "/Posts/")
		//Get All Posts
		if post == "" {
			fmt.Println("Getting Posts")
			err := db.ConvertAndWriteData(s.database.GetPosts(), w)
			if err != nil {
				return errors.New("Error converting data. Error: " + err.Error())
			}
			return nil
			//Get Single Post by ID
		}
		fmt.Println("Getting Post: " + post)
		postNum, err := strconv.Atoi(post)
		if err != nil {
			return errors.New("Error converting post number. Error: " + err.Error())
		}
		err = db.ConvertAndWriteData(s.database.GetPost(postNum), w)
		if err != nil {
			return errors.New("Error converting data. Error: " + err.Error())
		}
		//No Path exists
		return nil
	}
	//Invalid Path
	return errors.New("Invalid URL Path")
}
