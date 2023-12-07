package dtos

type TradeInput struct {
	OrderID       string  `json:"orderId"`
	InvestorID    string  `json:"investorId"`
	AssetID       string  `json:"assetId"`
	CurrentShares int     `json:"currentShares"`
	Shares        int     `json:"shares"`
	Price         float64 `json:"price"`
	OrderType     string  `json:"orderType"`
}

type OrderOutput struct {
	OrderID           string               `json:"orderId"`
	InvestorID        string               `json:"investorId"`
	AssetID           string               `json:"assetId"`
	OrderType         string               `json:"orderType"`
	Status            string               `json:"status"`
	Partial           int                  `json:"partial"`
	Shares            int                  `json:"shares"`
	TransactionOutput []*TransactionOutput `json:"transactions"`
}

type TransactionOutput struct {
	TransactionID string  `json:"transactionId"`
	BuyerId       string  `json:"buyerId"`
	SellerId      string  `json:"sellerId"`
	AssetID       string  `json:"assetId"`
	Price         float64 `json:"price"`
	Shares        int     `json:"shares"`
}
