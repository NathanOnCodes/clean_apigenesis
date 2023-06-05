package use_cases

import (
	"clean_architecture/api_genesis/internal/entity"
	"clean_architecture/api_genesis/internal/repository"
	"strings"
	"errors"
)

type ConversionUseCase struct {
	ConversionRepository repository.InterfaceConversionRepository
}

func NewConversionUseCase(conversionRepository repository.InterfaceConversionRepository) *ConversionUseCase {
	return &ConversionUseCase{ConversionRepository: conversionRepository}
}

func (service *ConversionUseCase) Create(amount, rate float64, to string) (*entity.Conversion, error) {
	calcParamsResult := amount * rate

	var symbol string

	coins := map[string]string{
		"BRL": "R$",
		"USD": "$",
		"BTC": "₿",
		"EUR": "€",
	}

	for i := range coins {
		if i == strings.ToUpper(to) {
			symbol = coins[i]
		}

	}

	conversion := entity.NewConversion(to, symbol, calcParamsResult)
	service.ConversionRepository.Create(conversion)

	return conversion, nil
}

func (service *ConversionUseCase) FindAll() ([]entity.Conversion, error) {
	data, err := service.ConversionRepository.FindAll()

	if err != nil {
		return nil, errors.New("not found database")
	}

	return data, nil
}
