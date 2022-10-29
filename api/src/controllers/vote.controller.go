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

func CreateVote(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var vote models.Vote
	if err = json.Unmarshal(reqBody, &vote); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = vote.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Conect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repos.NewVotesRepo(db)
	voteId, err := repo.Create(vote)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	vote.ID = voteId
	responses.JSON(w, http.StatusCreated, vote)
}

func GetGeneralRanking(w http.ResponseWriter, r *http.Request) {
	db, err := db.Conect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repos.NewVotesRepo(db)
	ranking, err := repo.FindAll()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, ranking)
}

func GetVotesByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
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

	repo := repos.NewVotesRepo(db)
	votes, err := repo.FindVotesByUser(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, votes)
}
