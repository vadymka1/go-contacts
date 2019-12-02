package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-crud/models"

	u "github.com/go-crud/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account)

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid Request"))

		return
	}

	resp := account.Create()
	u.Respond(w, resp)

}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account)

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid Request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)

}
