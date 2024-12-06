package main

import (
    "context"
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
    ctx := context.Background()
    code := r.URL.Query().Get("code")
    if code == "" {
        http.Error(w, "Code not found", http.StatusBadRequest)
        return
    }

    token, err := oauth2Config.Exchange(ctx, code)
    if err != nil {
        http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Use the token to get user information
    client := oauth2Config.Client(ctx, token)
    resp, err := client.Get("https://api.github.com/user/emails")
    if err != nil {
        http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Handle the response (e.g., read the body, parse JSON, etc.)
    // For demonstration, we'll just print the response status
    fmt.Fprintf(w, "Response Status: %s", resp.Status)
}

func main() {
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/api/auth/callback", callbackHandler)
    fmt.Println("Server is running on :3000")
    if err := http.ListenAndServe(":3000", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
