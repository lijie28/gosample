package route

import (
	"app/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

//登録
func RegistUser(w rest.ResponseWriter, r *rest.Request) {
	user := model.UserInfo{}
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

	//TODO: save in db

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
