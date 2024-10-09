package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pchchv/contacts/models"
	u "github.com/pchchv/contacts/utils"
)

func CreateContact(w http.ResponseWriter, r *http.Request) {
	// obtaining the ID of the user who sent the request
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}
	if err := json.NewDecoder(r.Body).Decode(contact); err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}

func GetContactsFor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetContacts(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
