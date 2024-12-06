package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    // Define the URL for the API endpoint
    url := "https://xyz.accurate.id/accurate//api/auth-info.do" // Replace with the actual endpoint

    // Create a new HTTP request
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    // Set the headers
    req.Header.Set("Authorization", "3f337af883d45a27dc1b7c131035708c") // Replace with your actual token
    req.Header.Set("X-Session-ID", "1ac49467-c4c4-4096-a654-875d1486e5a8") // Replace with your actual session ID

    // Create an HTTP client and make the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    // Read and print the response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }

    fmt.Println("Response Status:", resp.Status)
    fmt.Println("Response Body:", string(body))
}
