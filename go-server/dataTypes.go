package main

type Antartida struct {
	Identificacion string  `json:"identificacion"`
	Nombre         string  `json:"nombre"`
	Latitud        float64 `json:"latitud"`
	Longitud       float64 `json:"longitud"`
	Altitud        float64 `json:"altitud"`
	Srs            string  `json:"srs"`
	AltNieve       string  `json:"alt_nieve"`
	DDD            int     `json:"ddd"`
	DDDStd         int     `json:"dddstd"`
	DDDx           int     `json:"dddx"`
	Fhora          string  `json:"fhora"`
	HR             int     `json:"hr"`
	Ins            float64 `json:"ins"`
	Lluv           float64 `json:"lluv"`
	Pres           float64 `json:"pres"`
	RadKjM2        string  `json:"rad_kj_m2"`
	RadWM2         float64 `json:"rad_w_m2"`
	Rec            string  `json:"rec"`
	Temp           float64 `json:"temp"`
	Tmn            float64 `json:"tmn"`
	Tmx            float64 `json:"tmx"`
	Ts             float64 `json:"ts"`
	Tsb            string  `json:"tsb"`
	Tsmn           string  `json:"tsmn"`
	Tsmx           string  `json:"tsmx"`
	Vel            float64 `json:"vel"`
	Velx           float64 `json:"velx"`
	Albedo         float64 `json:"albedo"`
	Difusa         float64 `json:"difusa"`
	Directa        float64 `json:"directa"`
	Global         float64 `json:"global"`
	IrSolar        float64 `json:"ir_solar"`
	Neta           float64 `json:"neta"`
	Par            float64 `json:"par"`
	Tcielo         float64 `json:"tcielo"`
	Ttierra        float64 `json:"ttierra"`
	Uvab           float64 `json:"uvab"`
	Uvb            float64 `json:"uvb"`
	Uvi            float64 `json:"uvi"`
	Qdato          int     `json:"qdato"`
}
