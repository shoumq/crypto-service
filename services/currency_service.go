package services

import (
	"crypto-service/models"
	"crypto-service/repository"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type CurrencyService struct {
	repo *repository.CurrencyRepo
	mu   sync.Mutex
}

func NewCurrencyService(db *sql.DB) *CurrencyService {
	return &CurrencyService{repo: repository.NewCurrencyRepo(db)}
}

func (s *CurrencyService) AddCurrency(coin string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if price, _ := s.GetCurrencyPriceUSD(coin); price != 0 {
		return s.repo.AddCurrency(coin)
	}
	return nil
}

func (s *CurrencyService) RemoveCurrency(coin string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.repo.RemoveCurrency(coin)
}

func (s *CurrencyService) GetCurrencyPrice(coin string, timestamp int64) (models.CurrencyPrice, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.repo.GetCurrencyPrice(coin, timestamp)
}

func (s *CurrencyService) GetAllCurrencies() ([]models.Currency, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.repo.GetAllCurrencies()
}

func (s *CurrencyService) GetCurrencyPriceUSD(coin string) (float64, error) {
	apiKey := "ba12b76a-a714-4098-b451-9a688072ed44"
	url := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"
	parameters := "start=1&limit=100&convert=USD"
	req, err := http.NewRequest("GET", url+"?"+parameters, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Set("X-CMC_PRO_API_KEY", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}

	data, ok := result["data"].([]interface{})
	if !ok {
		return 0, nil
	}

	for _, item := range data {
		currency, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		symbol, ok := currency["symbol"].(string)
		if !ok {
			continue
		}
		if symbol == coin {
			quote, ok := currency["quote"].(map[string]interface{})
			if !ok {
				continue
			}
			usd, ok := quote["USD"].(map[string]interface{})
			if !ok {
				continue
			}
			price, ok := usd["price"].(float64)
			if !ok {
				continue
			}
			return price, nil
		}
	}

	return 0, nil
}

func (s *CurrencyService) UpdateCurrencyPrices() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		currencies, err := s.GetAllCurrencies()
		if err != nil {
			continue
		}

		for _, currency := range currencies {
			price, err := s.GetCurrencyPriceUSD(currency.Coin)
			if err != nil {
				continue
			}
			if err := s.repo.AddCurrencyPrice(currency.ID, price); err != nil {
				continue
			}
		}
	}
}

func (s *CurrencyService) GetAllPrices() ([]models.CurrencyPrice, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.repo.GetAllPrices()
}
