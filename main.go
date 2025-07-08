package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/", servePage)
    http.HandleFunc("/overview", serveOverview)
    http.HandleFunc("/about", serveAbout)
    http.HandleFunc("/solve", serveSolve)
    http.HandleFunc("/calculate", calculateHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
    fmt.Printf("🚀 Server running on port %s\n", port)
    http.ListenAndServe(":"+port, nil)
}
