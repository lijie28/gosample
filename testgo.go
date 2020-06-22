package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	// "sync"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/user/regist", RegistUser),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type UserInfo struct {
	Id       int
	Name     string
	Email    string
	Birthday string
}

func RegistUser(w rest.ResponseWriter, r *rest.Request) {
	user := UserInfo{}
	err := r.DecodeJsonPayload(&user)

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.Email == "" {
		rest.Error(w, "email required", 400)
		return
	}

	respond := SuccessRespond{Message: "OK"}

	fmt.Println(respond.getJsonString())
	w.WriteJson(respond.getJsonString())
}

type SuccessRespond struct {
	Message string `json: "mes"`
}

func (sucRes SuccessRespond) getJsonString() string {
	jsonBytes, err := json.Marshal(sucRes)
	if err != nil {
		fmt.Print(err)
		return "err.statement"
	}
	return string(jsonBytes)
}
