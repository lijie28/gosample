package route

import (
	"app/model"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

var usersList []model.UserInfo

func GetUsers(w rest.ResponseWriter, r *rest.Request) {
	user := model.UserInfo{}
	err := r.DecodeJsonPayload(&user)

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	result, err := json.Marshal(usersList)
	fmt.Println(string(result))

	for _, u := range usersList {
		if user.Token == u.Token {
			result, err := json.Marshal(u)
			if err != nil {
				fmt.Println("error:%s", err)
				return
			}
			fmt.Println(string(result))
			w.WriteJson(string(result))
			return
		}
	}

	w.WriteJson("no this account")

}

//登録
func RegistUser(w rest.ResponseWriter, r *rest.Request) {
	user := model.UserInfo{}
	err := r.DecodeJsonPayload(&user)

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO: Id       intを認証

	//TODO: Name     stringを認証

	//TODO: Email    stringを認証
	if user.Email == "" {
		rest.Error(w, "email required", 400)
		return
	}

	//TODO: Phone    stringを認証

	//tokenを作成
	user.Token = tokenGenerator()

	respond := SuccessRespond{Message: "OK"}

	//TODO: save in db
	usersList = append(usersList, user)

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

func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
