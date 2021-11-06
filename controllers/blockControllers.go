package controllers

import (
	"go-restful-api/services"
	u "go-restful-api/utils"
	"net/http"

	"github.com/gorilla/mux"
)

var GetBlockData = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	network := vars["network"]
	blockhash := vars["blockhash"]

	if network != "BTC" && network != "LTC" && network != "DOGE" {
		w.WriteHeader(http.StatusBadRequest)
		res := u.Message(false, "NETWORK NOT SUPPORTED!")
		u.Respond(w, res)
		return
	}

	if len(blockhash) != 64 {
		w.WriteHeader(http.StatusBadRequest)
		res := u.Message(false, "INVALID BLOCKHASH!")
		u.Respond(w, res)
		return
	}
	data := services.Block(network, blockhash)

	res := u.Message(true, "Success")
	res["data"] = data
	u.Respond(w, res)
}
