package restApi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type StubDatabase struct {
	users []User
	posts []Post
}

func (s *StubDatabase) GetUsers() []User {
	return s.users
}
func (s *StubDatabase) GetPosts() []Post {
	return s.posts
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

func TestServer(t *testing.T) {

	getTests := []struct {
		name        string
		database    *StubDatabase
		requestPath string
		want        string
	}{
		{
			name: "Get all users from database",
			database: &StubDatabase{
				users: []User{{Name: "James Clarke"}, {Name: "Jane Smith"}},
				posts: []Post{},
			},
			requestPath: "/Users/",
			want:        "[{\"Name\":\"James Clarke\"},{\"Name\":\"Jane Smith\"}]",
		},
		{
			name: "Get all posts from database",
			database: &StubDatabase{
				users: []User{},
				posts: []Post{{User: User{Name: "James Clarke"}, Text: "Hello, World!", Id: 1}, {User: User{Name: "Jane Smith"}, Text: "Hello, World!", Id: 2}},
			},
			requestPath: "/Posts/",
			want:        "[{\"Text\":\"Hello, World!\",\"User\":{\"Name\":\"James Clarke\"},\"Id\":1},{\"Text\":\"Hello, World!\",\"User\":{\"Name\":\"Jane Smith\"},\"Id\":2}]",
		},
		{
			name: "Get a user from database",
			database: &StubDatabase{
				users: []User{{Name: "James Clarke"}, {Name: "Jane Smith"}},
				posts: []Post{},
			},
			requestPath: "/Users/JamesClarke",
			want:        "{\"Name\":\"James Clarke\"}",
		},
		{
			name: "Get a post from database",
			database: &StubDatabase{
				users: []User{},
				posts: []Post{{User: User{Name: "James Clarke"}, Text: "Hello, World!", Id: 1}, {User: User{Name: "Jane Smith"}, Text: "Hello, World!", Id: 2}},
			},
			requestPath: "/Posts/1",
			want:        "{\"Text\":\"Hello, World!\",\"User\":{\"Name\":\"James Clarke\"},\"Id\":1}",
		},
	}

	for _, tt := range getTests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, tt.requestPath, nil)
			response := httptest.NewRecorder()

			server := &SocialMediaServer{database: tt.database}

			err := server.ServeHTTP(response, request)

			if err != nil {
				t.Fatalf("ServeHTTP has returned an error %v", err)
			}

			got := response.Body.String()

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
