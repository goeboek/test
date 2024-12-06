package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

func main() {
	// Set up the OAuth2 configuration
	config := &oauth2.Config{
		ClientID:     "1ac49467-c2c2-4096-a654-875e1486e5a8",
		ClientSecret: "3f331af883d35a27dc1b7c131025725c",
		Scopes:       []string{"read", "write"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://xyz.accurate.id/accurate/api/auth-info.do",
			TokenURL: "https://xyz.accurate.id/accurate/api/auth-info.do",
		},
		RedirectURL: "https://your-vercel-app.vercel.app/callback",
	}

	// Set up the HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := config.AuthCodeURL("state-token")
		http.Redirect(w, r, url, http.StatusFound)
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Use the token to make API calls
		client := config.Client(context.Background(), token)
		resp, err := client.Get("https://xyz.accurate.id/accurate/api/user-info.do")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var userInfo map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "User info: %v", userInfo)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Printf("Starting server on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
