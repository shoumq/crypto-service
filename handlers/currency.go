package handlers

import (
	"crypto-service/models"
	"crypto-service/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CurrencyHandler struct {
	db *sql.DB
}

func NewCurrencyHandler(db *sql.DB) *CurrencyHandler {
	return &CurrencyHandler{db: db}
}

// AddCurrency godoc
// @Summary Add a new currency to the watchlist
// @Description Add a new currency to the watchlist
// @Accept  json
// @Produce  json
// @Param   coin     body    models.AddCurrencyRequest  true  "Currency to add"
// @Success 200 {object} models.AddCurrencyRequest
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /currency/add [post]
func (h *CurrencyHandler) AddCurrency(c *gin.Context) {
	var req models.AddCurrencyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := services.NewCurrencyService(h.db)
	if err := service.AddCurrency(req.Coin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// RemoveCurrency godoc
// @Summary Remove a currency from the watchlist
// @Description Remove a currency from the watchlist
// @Accept  json
// @Produce  json
// @Param   coin     body    models.RemoveCurrencyRequest  true  "Currency to remove"
// @Success 200 {object} models.RemoveCurrencyRequest
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /currency/remove [post]
func (h *CurrencyHandler) RemoveCurrency(c *gin.Context) {
	var req models.RemoveCurrencyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := services.NewCurrencyService(h.db)
	if err := service.RemoveCurrency(req.Coin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// GetCurrencyPrice godoc
// @Summary Get the price of a currency at a specific timestamp
// @Description Get the price of a currency at a specific timestamp
// @Accept  json
// @Produce  json
// @Param   request     body    models.GetCurrencyPriceRequest  true  "Currency and timestamp"
// @Success 200 {object} models.CurrencyPrice
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /currency/price [post]
func (h *CurrencyHandler) GetCurrencyPrice(c *gin.Context) {
	var req models.GetCurrencyPriceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := services.NewCurrencyService(h.db)
	price, err := service.GetCurrencyPrice(req.Coin, req.Timestamp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, price)
}

// GetAllCurrencies godoc
// @Summary Get all currencies in the watchlist
// @Description Get all currencies in the watchlist
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Currency
// @Failure 500 {object} string
// @Router /currency/all [get]
func (h *CurrencyHandler) GetAllCurrencies(c *gin.Context) {
	service := services.NewCurrencyService(h.db)
	currencies, err := service.GetAllCurrencies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, currencies)
}

// GetAllPrices godoc
// @Summary Get all prices
// @Description Get all prices
// @Accept  json
// @Produce  json
// @Success 200 {array} models.CurrencyPrice
// @Failure 500 {object} string
// @Router /currency/prices [get]
func (h *CurrencyHandler) GetAllPrices(c *gin.Context) {
	service := services.NewCurrencyService(h.db)
	prices, err := service.GetAllPrices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prices)
}
