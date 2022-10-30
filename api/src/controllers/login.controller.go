package controllers

import (
	"api/src/auth"
	"api/src/db"
	"api/src/models"
	repos "api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Conect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repos.NewUsersRepo(db)
	loggedUser, err := repo.FindByMail(user.Email)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(loggedUser.Password, user.Password); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, erro := auth.GenerateToken(loggedUser.ID)
	if erro != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
