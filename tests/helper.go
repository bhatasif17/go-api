package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type ResponseModel struct {
	Status bool `json:"status"`
}

type LoginModel struct {
	Account struct {
		ID    int16  `json:"ID"`
		Token string `json:"token"`
	}
	Status bool `json:"status"`
}

func GetAuthorizationToken() string {
	apiurl := "http://localhost:8000/api/user/login"

	reqBody, err := json.Marshal(map[string]string{
		"email":    "a@a.com",
		"password": "Asif@123",
	})
	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post(apiurl,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		print(err)
	}

	blockStruct := LoginModel{}
	json.Unmarshal(body, &blockStruct)

	return blockStruct.Account.Token
}

func ApiHelper(t *testing.T, url string) bool {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fail()
	}
	token := GetAuthorizationToken()

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer " + token},
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fail()
	}

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fail()
	}

	blockStruct := ResponseModel{}
	json.Unmarshal(response, &blockStruct)
	return blockStruct.Status
}
