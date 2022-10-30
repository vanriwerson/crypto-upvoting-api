package routes

import (
	"api/src/controllers"
	"net/http"
)

var voteRoutes = []Route{
	{
		URI:          "/vote",
		Method:       http.MethodPost,
		Controller:   controllers.CreateVote,
		RequiresAuth: true,
	},
	{
		URI:          "/ranking",
		Method:       http.MethodGet,
		Controller:   controllers.GetGeneralRanking,
		RequiresAuth: false,
	},
	{
		URI:          "/ranking/{id}",
		Method:       http.MethodGet,
		Controller:   controllers.GetVotesByUser,
		RequiresAuth: false,
	},
}
