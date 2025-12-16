package models

type Difference struct {
	Date            string  `json:"date"`
	Reparto         string  `json:"reparto"`
	Diferencia      float64 `json:"diferencia"`
	Tipo            string  `json:"tipo"`
	UserName        string  `json:"user_name"`
	DepositEsperado float64 `json:"deposit_esperado"`
	TotalAmount     float64 `json:"total_amount"`
}

type Statistics struct {
	TotalDiferencias int `json:"total_diferencias"`
	TotalFaltantes   int `json:"total_faltantes"`
	TotalSobrantes   int `json:"total_sobrantes"`
	Consolidados     int `json:"consolidados"`
}

type DifferenceResponse struct {
	Status       string       `json:"status"`
	Desde        string       `json:"desde"`
	Hasta        string       `json:"hasta"`
	Estadisticas Statistics   `json:"estadisticas"`
	Items        []Difference `json:"items"`
}
