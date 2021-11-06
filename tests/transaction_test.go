package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TransactionModel struct {
	Status bool `json:"status"`
}

func TestTransaction_ValidResponse(t *testing.T) {
	url := "http://localhost:8000/api/transaction/BTC/dbaf14e1c476e76ea05a8b71921a46d6b06f0a950f17c5f9f1a03b8fae467f10"

	actual := ApiHelper(t, url)

	assert.Equal(t, true, actual, "Should be same")
}

func TestTransaction_InValidNetwork(t *testing.T) {
	url := "http://localhost:8000/api/transaction/BTCCV/dbaf14e1c476e76ea05a8b71921a46d6b06f0a950f17c5f9f1a03b8fae467f10"

	actual := ApiHelper(t, url)

	assert.Equal(t, false, actual, "Should be same")
}
func TestTransaction_InValidTXID(t *testing.T) {
	url := "http://localhost:8000/api/transaction/BTC/AAAAAAdbaf14e1c476e76ea05a8b71921a46d6b06f0a950f17c5f9f1a03b8fae467f10"

	actual := ApiHelper(t, url)

	assert.Equal(t, false, actual, "Should be same")
}
func TestTransaction_InValidNetworkAndTXID(t *testing.T) {
	url := "http://localhost:8000/api/transaction/BTCCC/AAAAAdbaf14e1c476e76ea05a8b71921a46d6b06f0a950f17c5f9f1a03b8fae467f10"

	actual := ApiHelper(t, url)

	assert.Equal(t, false, actual, "Should be same")
}
