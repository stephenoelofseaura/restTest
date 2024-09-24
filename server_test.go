package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServerPost(t *testing.T) {
	createTests := []struct {
		name          string
		database      *StubDatabase
		path          string
		requestString string
		want          string
	}{
		{
			name: "Create User",
			database: &StubDatabase{
				users: []User{},
				posts: []Post{},
			},
			path:          "/Users/",
			requestString: "{\"Name\":\"James Clarke\"}",
			want:          "{\"Name\":\"James Clarke\"}",
		},
		{
			name: "Create Post",
			database: &StubDatabase{
				users: []User{{Name: "James Clarke"}},
				posts: []Post{},
			},
			path:          "/Posts/",
			requestString: "{\"Text\":\"Hello, World!\",\"User\":{\"Name\":\"James Clarke\"},\"Id\":1,\"Likes\":1,\"Comments\":null}",
			want:          "{\"Text\":\"Hello, World!\",\"User\":{\"Name\":\"James Clarke\"},\"Id\":1,\"Likes\":1,\"Comments\":null}",
		},
		{
			name: "Create Comment",
			database: &StubDatabase{
				users: []User{{Name: "James Clarke"}},
				posts: []Post{{User: User{Name: "James Clarke"}, Text: "Hello, World!", Id: 1, Likes: 1}},
			},
			path:          "/Posts/1/Comments/",
			requestString: "{\"Text\":\"This is a comment\",\"User\":{\"Name\":\"James Clarke\"},\"Id\":5}",
			want:          "{\"Text\":\"This is a comment\",\"User\":{\"Name\":\"James Clarke\"},\"Id\":5}",
		},
	}

	for _, tt := range createTests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, tt.path, strings.NewReader(tt.requestString))
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

func TestServerGet(t *testing.T) {
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
				posts: []Post{{User: User{Name: "James Clarke"}, Text: "Hello, World!", Id: 1, Likes: 1, Comments: Comments{{Text: "This is a comment", User: User{Name: "Jane Smith"}, Id: 5}}}, {User: User{Name: "Jane Smith"}, Text: "Hello, World!", Id: 2, Likes: 0}},
			},
			requestPath: "/Posts/",
			want:        "[{\"Text\":\"Hello, World!\",\"User\":{\"Name\":\"James Clarke\"},\"Id\":1,\"Likes\":1,\"Comments\":[{\"Text\":\"This is a comment\",\"User\":{\"Name\":\"Jane Smith\"},\"Id\":5}]},{\"Text\":\"Hello, World!\",\"User\":{\"Name\":\"Jane Smith\"},\"Id\":2,\"Likes\":0,\"Comments\":null}]",
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
				posts: []Post{{User: User{Name: "James Clarke"}, Text: "Hello, World!", Id: 1, Likes: 1}, {User: User{Name: "Jane Smith"}, Text: "Hello, World!", Id: 2, Likes: 0}},
			},
			requestPath: "/Posts/1",
			want:        "{\"Text\":\"Hello, World!\",\"User\":{\"Name\":\"James Clarke\"},\"Id\":1,\"Likes\":1,\"Comments\":null}",
		},
		{
			name: "Get all posts from a user",
			database: &StubDatabase{
				users: []User{{Name: "James Clarke"}, {Name: "Jane Smith"}},
				posts: []Post{{User: User{Name: "James Clarke"}, Text: "Hello, World!", Id: 1, Likes: 1}, {User: User{Name: "Jane Smith"}, Text: "Hello, World!", Id: 2, Likes: 0}, {User: User{Name: "James Clarke"}, Text: "Hello, World!", Id: 3, Likes: 0}},
			},
			requestPath: "/Users/JamesClarke/Posts/",
			want:        "[{\"Text\":\"Hello, World!\",\"User\":{\"Name\":\"James Clarke\"},\"Id\":1,\"Likes\":1,\"Comments\":null},{\"Text\":\"Hello, World!\",\"User\":{\"Name\":\"James Clarke\"},\"Id\":3,\"Likes\":0,\"Comments\":null}]",
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

	updateTests := []struct {
		name          string
		updatedString string
		database      *StubDatabase
		requestPath   string
		body          string
		want          string
	}{
		{
			name:          "Update user",
			updatedString: "{\"Name\":\"James Mark\"}",
			database: &StubDatabase{
				users: []User{{Name: "James Clarke"}, {Name: "Jane Smith"}},
				posts: []Post{},
			},
			requestPath: "/Users/JamesClarke/",
			body:        "{\"Name\":\"James Mark\"}",
			want:        "{\"Name\":\"James Mark\"}",
		},
		{
			name:          "Update post",
			updatedString: "{\"Text\":\"Hello, World!\",\"User\":{\"Name\":\"James Clarke\"},\"Id\":1,\"Likes\":1,\"Comments\":null}",
			database: &StubDatabase{
				users: []User{{Name: "James Clarke"}, {Name: "Jane Smith"}},
				posts: []Post{{User: User{Name: "James Clarke"}, Text: "Hello, World!", Id: 1, Likes: 1, Comments: Comments{{Text: "This is a comment", User: User{Name: "Jane Smith"}, Id: 5}}}, {User: User{Name: "Jane Smith"}, Text: "Hello, World!", Id: 2, Likes: 0}},
			},
			requestPath: "/Posts/1/",
			body:        "{\"Text\":\"Hello, World!\",\"User\":{\"Name\":\"James Clarke\"},\"Id\":1,\"Likes\":1,\"Comments\":[{\"Text\":\"This is a comment\",\"User\":{\"Name\":\"Jane Smith\"},\"Id\":5}]}",
			want:        "{\"Text\":\"Hello, World!\",\"User\":{\"Name\":\"James Clarke\"},\"Id\":1,\"Likes\":1,\"Comments\":[{\"Text\":\"This is a comment\",\"User\":{\"Name\":\"Jane Smith\"},\"Id\":5}]}",
		},
	}

	for _, tt := range updateTests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPut, tt.requestPath, strings.NewReader(tt.body))
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
