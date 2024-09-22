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
