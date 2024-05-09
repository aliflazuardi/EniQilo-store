package main

import "fmt"

func main() {
	fmt.Println("Welcome to EniQilo Store")


  // http server on port 8080
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to EniQilo Store")
  })
  http.ListenAndServe(":8080", nil)
}
