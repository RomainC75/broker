package dto

type BinanceMessageDto struct {
	E      string      `json:"e" validate:"required"`
	Ee     int64       `json:"E" validate:"required"`
	Symbol string      `json:"s" validate:"required"`
	U      int64       `json:"U" validate:"required"`
	Uu     int64       `json:"u"`
	B      [][]float64 `json:"b"`
	A      int64       `json:"a"`
	Price  string      `json:"p"`
	Qty    string      `json:"q"`
	Time   int64       `json:"T"`
}
