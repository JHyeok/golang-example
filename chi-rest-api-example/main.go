package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", HelloWorld)
	r.Get("/ping", Pong)

	r.Mount("/users", usersResource{}.Routes())
	r.Mount("/posts", postsResource{}.Routes())

	log.Fatal(http.ListenAndServe(":3000", r))
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func Pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
