package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pchchv/contacts/models"
	u "github.com/pchchv/contacts/utils"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	// decode the request body into struct and failed if any error occur
	if err := json.NewDecoder(r.Body).Decode(account); err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create()
	u.Respond(w, resp)
}
