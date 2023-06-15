package main

import (
	"clean_architecture/api_genesis/handler"
	"clean_architecture/api_genesis/internal/infra"
	"clean_architecture/api_genesis/internal/repository"
	"clean_architecture/api_genesis/internal/use_cases"
	"log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
func main(){

	client, err := infra.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	
	conversionRepository := repository.NewMongoDB(client.GetClient())
	conversionUseCase := use_cases.NewUseCaseConversion(conversionRepository)
	conversionHandler := handler.NewHandler(conversionUseCase)

	app := echo.New()
	app.Use(middleware.Recover())

	configureRoutes(app, conversionHandler)

	app.Start(":8080")
}


func configureRoutes(app *echo.Echo, ctx *handler.ConversionHandler) {
	app.GET("/api/v1/exchange", ctx.IndexExchangeEndPoint)
	app.POST("/api/v1/exchange/:amount/:from/:to/:rate", ctx.CreateConversionEndPoint)
	app.GET("/api/v1/exchange/logs", ctx.FindAllEndPoint)
}