package restApi

//Add User
//Add Post
//Get User
//Get Post
//Add Likes
//Get Likes
//Get Comments
//Add Comments

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Name string `json:"Name"`
}

type Post struct {
	Text string `json:"Text"`
	User User   `json:"User"`
	Id   int    `json:"Id"`
}

type Database interface {
	GetUsers() []User
	GetPosts() []Post
	GetPost(id int) Post
	GetUser(name string) User
}

type SocialMediaServer struct {
	database Database
}

func (s *SocialMediaServer) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {

	case http.MethodGet:

		if strings.HasPrefix(r.URL.Path, "/Users/") {
			user := strings.TrimPrefix(r.URL.Path, "/Users/")
			if user == "" {
				fmt.Println("Getting Users")
				err := convertAndWriteData(s.database.GetUsers(), w)
				if err != nil {
					return errors.New("Error converting data. Error: " + err.Error())
				}
				return nil
			} else {
				user = strings.TrimSuffix(user, "/")
				fmt.Println("Getting User: " + user)
				err := convertAndWriteData(s.database.GetUser(user), w)
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
				err := convertAndWriteData(s.database.GetPosts(), w)
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
				err = convertAndWriteData(s.database.GetPost(postNum), w)
				if err != nil {
					return errors.New("Error converting data. Error: " + err.Error())
				}
				return nil
			}
		}
	}

	return errors.New("Unsupported method")
}

func convertAndWriteData(data any, w http.ResponseWriter) error {
	convertedData, err := json.Marshal(data)
	if err != nil {
		return errors.New("Error marshalling data, Error: " + err.Error())
	}
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(convertedData)
	if writeErr != nil {
		return errors.New("Error writing response")
	}
	return nil
}

func main() {
}
