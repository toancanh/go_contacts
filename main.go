package main

import (
	"github.com/gorilla/mux"
	"go_contacts/app"
	"fmt"
	"go_contacts/controllers"
	"go_contacts/pkg/setting"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/contacts/{id:[0-9]+}", controllers.GetContactsFor).Methods("GET") //  user/2/contacts

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	err := http.ListenAndServe(":" + setting.AppSetting.Port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}