package main

import (
    "fmt"
    "net/http"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/github"
)

var (
    oauth2Config = &oauth2.Config{
        ClientID:     "YOUR_CLIENT_ID",
        ClientSecret: "YOUR_CLIENT_SECRET",
        RedirectURL:  "https://your-app-name.vercel.app/api/auth/callback",
        Scopes:       []string{"user:email"},
        Endpoint:     github.Endpoint,
    }
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
    url := oauth2Config.AuthCodeURL("state")
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
    // Handle the callback and exchange the code for a token
}

func main() {
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/api/auth/callback", callbackHandler)
    http.ListenAndServe(":3000", nil)
}
