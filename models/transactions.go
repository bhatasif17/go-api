package models

type TransactionDTO struct {
	Data struct {
		TXID      string `json:"txid"`
		Fee       string `json:"fee"`
		Time      int64  `json:"time"`
		SentValue string `json:"sent_value"`
	}
	Status string `json:"status"`
}

type TransactionResponse struct {
	TXID       string
	Fee        string
	Time       string
	Sent_Value string
}
