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

func Block(network string, blockhash string) models.BlockResponse {

	url := "https://sochain.com/api/v2/get_block/" + network + "/" + blockhash
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

	// Convert response body to string
	responseString := string(responseData)
	fmt.Println("API Response: \n" + responseString)

	//Convert response body to Todo struct
	blockStruct := models.BlockDTO{}
	json.Unmarshal(responseData, &blockStruct)

	var model = models.BlockResponse{}
	if blockStruct.Status == "fail" {
		return model
	}

	model.Network = blockStruct.Data.Network
	model.BlockNo = blockStruct.Data.BlockNo
	model.Next_Block_Hash = blockStruct.Data.NextBlockHash
	model.Previous_Block_Hash = blockStruct.Data.PreviousBlockHash
	model.Size = blockStruct.Data.Size
	model.Time = time.Unix(blockStruct.Data.Time, 0).UTC().Format("02/01/2006 15:04")

	for i := 0; i < 10; i++ {
		model.Transactions[i].TXID = blockStruct.Data.Transactions[i]
		data := GetTransaction(network, blockStruct.Data.Transactions[i])
		model.Transactions[i].Fee = data.Fee
		model.Transactions[i].Sent_Value = data.Sent_Value
		model.Transactions[i].Time = data.Time
	}
	return model
}
