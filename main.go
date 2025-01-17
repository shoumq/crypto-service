package main

import (
	_ "crypto-service/docs"
	"crypto-service/handlers"
	"crypto-service/services"
	"crypto-service/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	db, err := utils.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	r := gin.Default()
	currencyHandler := handlers.NewCurrencyHandler(db)

	r.POST("/currency/add", currencyHandler.AddCurrency)
	r.POST("/currency/remove", currencyHandler.RemoveCurrency)
	r.POST("/currency/price", currencyHandler.GetCurrencyPrice)
	r.GET("/currency/all", currencyHandler.GetAllCurrencies)
	r.GET("/currency/prices", currencyHandler.GetAllPrices)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Server started at :8080")

	currencyService := services.NewCurrencyService(db)
	go currencyService.UpdateCurrencyPrices()

	log.Fatal(r.Run(":8080"))
}
