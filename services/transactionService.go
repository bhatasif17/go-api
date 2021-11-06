package services

import (
	"encoding/json"
	"fmt"
	"go-restful-api/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetTransaction(network string, txid string) models.TransactionResponse {
	url := "https://sochain.com/api/v2/tx/" + network + "/" + txid
	fmt.Println("API Request: \n" + url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	fmt.Println("API Response: \n" + responseString)

	var todoStruct models.TransactionDTO
	json.Unmarshal(responseData, &todoStruct)

	var model = models.TransactionResponse{}
	if todoStruct.Status == "fail" {
		return model
	}

	model.TXID = todoStruct.Data.TXID
	model.Fee = todoStruct.Data.Fee
	model.Sent_Value = todoStruct.Data.SentValue
	model.Time = time.Unix(todoStruct.Data.Time, 0).UTC().Format("02/01/2006 15:04")

	return model
}
