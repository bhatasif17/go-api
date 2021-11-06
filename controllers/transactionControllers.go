package controllers

import (
	"go-restful-api/services"
	u "go-restful-api/utils"
	"net/http"

	"github.com/gorilla/mux"
)

var GetTransactionData = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	network := vars["network"]
	txid := vars["txid"]

	if network != "BTC" && network != "LTC" && network != "DOGE" {
		w.WriteHeader(http.StatusBadRequest)
		res := u.Message(false, "NETWORK NOT SUPPORTED!")
		u.Respond(w, res)
		return
	}

	if len(txid) != 64 {
		w.WriteHeader(http.StatusBadRequest)
		res := u.Message(false, "INVALID TXID!")
		u.Respond(w, res)
		return
	}

	data := services.GetTransaction(network, txid)

	res := u.Message(true, "SUCCESS!")
	res["data"] = data
	u.Respond(w, res)
}
