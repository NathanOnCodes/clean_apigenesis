package entity

type Conversion struct {
	NameCoin string `json:"name" bson:"name"`
	SymbolCoin string `json:"symbol" bson:"symbol"`
	Value float64	`json:"value" bson:"value"`
}

func New(name, symbol string, value float64) *Conversion{
	return &Conversion{
		NameCoin: name,
		SymbolCoin: symbol,
		Value: value,	
	}
}