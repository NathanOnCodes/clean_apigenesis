package handler

import (
	"clean_architecture/api_genesis/internal/use_cases"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


type ConversionHandler struct {
	newConversionUseCase *use_cases.UseCaseConversion
}

func NewHandler(property *use_cases.UseCaseConversion) *ConversionHandler {
	return &ConversionHandler{ newConversionUseCase: property,	}
}

func (h *ConversionHandler) IndexExchangeEndPoint(ctx echo.Context) error {
	message := "Bem vindo, para converter as moedas siga esse exemplo:\nhttp://localhost:8000/exchange/valor/moeda-atual/moeda-destino/cotação\nvoce tambem pode conferir o seu histórico de conversões na rota exchange/logs"
	return ctx.String(http.StatusOK, message)
}


func (h *ConversionHandler) CreateConversionEndPoint(ctx echo.Context) error {
	amount, amountErr := strconv.ParseFloat(ctx.Param("amount"), 64)
	rate, rateErr := strconv.ParseFloat(ctx.Param("rate"), 64)

	if amountErr != nil || rateErr != nil {
		log.Fatal("ocorreu um erro")
	}

	conversion, err := h.newConversionUseCase.Create(amount, rate, ctx.Param("to"))
	if err != nil {
		return err
	}

	response := map[string]interface{}{
		"moeda": conversion.SymbolCoin,
		"valor": conversion.Value,
	}

	return ctx.JSON(http.StatusOK, response)
}


func (h *ConversionHandler) FindAllEndPoint(ctx echo.Context) error {
	result, err := h.newConversionUseCase.FindAll()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}