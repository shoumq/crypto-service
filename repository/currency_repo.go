package repository

import (
	"crypto-service/models"
	"database/sql"
	"time"
)

type CurrencyRepo struct {
	db *sql.DB
}

func NewCurrencyRepo(db *sql.DB) *CurrencyRepo {
	return &CurrencyRepo{db: db}
}

func (r *CurrencyRepo) AddCurrency(coin string) error {
	_, err := r.db.Exec("INSERT INTO currencies (coin) VALUES ($1) ON CONFLICT (coin) DO NOTHING", coin)
	return err
}

func (r *CurrencyRepo) RemoveCurrency(coin string) error {
	_, err := r.db.Exec("DELETE FROM currencies WHERE coin = $1", coin)
	return err
}

func (r *CurrencyRepo) GetCurrencyPrice(coin string, timestamp int64) (models.CurrencyPrice, error) {
	var price models.CurrencyPrice
	err := r.db.QueryRow(`
        SELECT coin_id, price, timestamp
        FROM prices
        WHERE coin_id = (SELECT id FROM currencies WHERE coin = $1)
        AND timestamp = (
            SELECT timestamp
            FROM prices
            WHERE coin_id = (SELECT id FROM currencies WHERE coin = $1)
            AND timestamp <= $2
            ORDER BY ABS(timestamp - $2)
            LIMIT 1
        )
        LIMIT 1`, coin, timestamp).Scan(&price.CoinID, &price.Price, &price.Timestamp)
	return price, err
}

func (r *CurrencyRepo) GetAllCurrencies() ([]models.Currency, error) {
	rows, err := r.db.Query("SELECT id, coin FROM currencies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var currencies []models.Currency
	for rows.Next() {
		var currency models.Currency
		if err := rows.Scan(&currency.ID, &currency.Coin); err != nil {
			return nil, err
		}
		currencies = append(currencies, currency)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return currencies, nil
}

func (r *CurrencyRepo) AddCurrencyPrice(coinID int, price float64) error {
	timestamp := time.Now().Unix()
	_, err := r.db.Exec("INSERT INTO prices (coin_id, price, timestamp) VALUES ($1, $2, $3)", coinID, price, timestamp)
	return err
}

func (r *CurrencyRepo) GetAllPrices() ([]models.CurrencyPrice, error) {
	rows, err := r.db.Query("SELECT coin_id, price, timestamp FROM prices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prices []models.CurrencyPrice
	for rows.Next() {
		var price models.CurrencyPrice
		if err := rows.Scan(&price.CoinID, &price.Price, &price.Timestamp); err != nil {
			return nil, err
		}
		prices = append(prices, price)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return prices, nil
}
