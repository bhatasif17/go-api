package models

type BlockDTO struct {
	Data struct {
		Network           string   `json:"network"`
		BlockNo           int64    `json:"block_no"`
		Time              int64    `json:"time"`
		PreviousBlockHash string   `json:"previous_blockhash"`
		NextBlockHash     string   `json:"next_blockhash"`
		Size              int      `json:"size"`
		Transactions      []string `json:"txs"`
	}
	Status string `json:"status"`
}

type BlockResponse struct {
	Network             string
	BlockNo             int64
	Time                string
	Previous_Block_Hash string
	Next_Block_Hash     string
	Size                int
	Transactions        [10]TransactionResponse
}
