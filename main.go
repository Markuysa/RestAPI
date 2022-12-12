package main

import (
	handlers "RestAPI/handlers"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type handlerStruct struct {
	usersCount int
	//dbConnection sql.DB
}

func (h *handlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	command := strings.Join(strings.Split(r.URL.Path, "/api/users"), "")
	var IDint int
	if strings.Contains(command, "/getUser/") {
		ID := strings.Join(strings.Split(command, "/getUser/"), "")
		IDint, _ = strconv.Atoi(ID)
		log.Print(IDint)
		command = "/getUser/"
	}
	switch command {
	case "/getAll":
		users := handlers.GetAll()
		fmt.Fprintln(w, users)
	case "/getUser/":
		fmt.Fprintln(w, handlers.GetByID(IDint))
	}
}

func main() {

	muxer := http.NewServeMux()

	server := http.Server{
		Addr:         ":4040",
		Handler:      muxer,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	usersHandler := handlerStruct{
		usersCount: 10,
	}
	muxer.Handle("/api/users/", &usersHandler)
	log.Print("Started server at ", server.Addr, " host")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
