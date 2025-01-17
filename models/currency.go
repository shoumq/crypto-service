package models

type AddCurrencyRequest struct {
	Coin string `json:"coin"`
}

type RemoveCurrencyRequest struct {
	Coin string `json:"coin"`
}

type GetCurrencyPriceRequest struct {
	Coin      string `json:"coin"`
	Timestamp int64  `json:"timestamp"`
}

type CurrencyPrice struct {
	CoinID    int     `json:"coin_id"`
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
}

type Currency struct {
	ID   int    `json:"id"`
	Coin string `json:"coin"`
}
