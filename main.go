package main

import (
	"fmt"
	"net/http"

	"os"

	authcontroller "github.com/jeypc/go-auth/controllers"
)

func main() {

	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/", fs)

	fmt.Fprintln(os.Stdout, []any{"Server jalan di: http://localhost:3000"}...)
	http.ListenAndServe(":3000", nil)
}
