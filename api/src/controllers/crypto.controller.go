package controllers

import (
	"api/src/db"
	"api/src/models"
	repos "api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateCrypto(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var crypto models.Crypto
	if err = json.Unmarshal(reqBody, &crypto); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = crypto.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Conect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repos.NewCryptosRepo(db)
	cryptoId, err := repo.Create(crypto)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	crypto.ID = cryptoId
	responses.JSON(w, http.StatusCreated, crypto)
}

func FindAllCryptos(w http.ResponseWriter, r *http.Request) {
	db, err := db.Conect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repos.NewCryptosRepo(db)
	cryptos, err := repo.FindAll()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, cryptos)
}

func FindCryptoById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cryptoID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Conect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repos.NewCryptosRepo(db)
	crypto, err := repo.FindOne(cryptoID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, crypto)
}

func UpdateCrypto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cryptoID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var crypto models.Crypto
	if err = json.Unmarshal(reqBody, &crypto); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = crypto.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Conect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repos.NewCryptosRepo(db)
	if err = repo.Update(cryptoID, crypto); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteCrypto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cryptoID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Conect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repos.NewCryptosRepo(db)
	if err = repo.Delete(cryptoID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
