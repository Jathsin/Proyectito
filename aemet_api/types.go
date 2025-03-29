package aemet_api

import "fmt"

//get_uv_prediction

type data_url_response struct {
	Descripcion string `json:"descripcion"`
	Estado      int    `json:"estado"`
	Datos       string `json:"datos"`
	Metadatos   string `json:"metadatos"`
}

type Data_response_data struct {
	Root Root `json:"root"`
}

type Root struct {
	FechaValidez     string   `json:"FECHA_VALIDEZ"`
	FechaMod         string   `json:"FECHA_MOD"`
	FechaElaboracion string   `json:"FECHA_ELABORACION"`
	Ciudad           []Ciudad `json:"CIUDAD"`
}

type Ciudad struct {
	Uv       int    `json:"uv"`
	Valor    string `json:"valor"`
	Canarias int    `json:"canarias"`
	Id       IDType `json:"id"`
}

type IDType any

// Parses ID into string
func ToStringValue(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return fmt.Sprintf("%d", v)
	case int64:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%g", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// get_uv_local_prediction
type Data_response_data_local []BeachForecast

type BeachForecast struct {
	Origen struct {
		Productor string `json:"productor"`
		Web       string `json:"web"`
		Language  string `json:"language"`
		Copyright string `json:"copyright"`
		NotaLegal string `json:"notaLegal"`
	} `json:"origen"`
	Elaborado  string          `json:"elaborado"`
	Nombre     string          `json:"nombre"`
	Localidad  int             `json:"localidad"`
	Prediccion BeachPrediction `json:"prediccion"`
	ID         int             `json:"id"`
}

type BeachPrediction struct {
	Dia []DayForecast `json:"dia"`
}

type DayForecast struct {
	EstadoCielo WeatherElement     `json:"estadoCielo"`
	Viento      WeatherElement     `json:"viento"`
	Oleaje      WeatherElement     `json:"oleaje"`
	TMaxima     ValueElement       `json:"tMaxima"`
	STermica    DescriptionElement `json:"sTermica"`
	TAgua       ValueElement       `json:"tAgua"`
	UvMax       ValueElement       `json:"uvMax"`
	Fecha       int                `json:"fecha"`
	// Lowercase duplicates
	TMMaxima  ValueElement       `json:"tmaxima"`
	STTermica DescriptionElement `json:"stermica"`
	TAgua2    ValueElement       `json:"tagua"`
}

type WeatherElement struct {
	Value        string `json:"value"`
	F1           int    `json:"f1"`
	Descripcion1 string `json:"descripcion1"`
	F2           int    `json:"f2"`
	Descripcion2 string `json:"descripcion2"`
}

type ValueElement struct {
	Value  string `json:"value"`
	Valor1 int    `json:"valor1"`
}

type DescriptionElement struct {
	Value        string `json:"value"`
	Valor1       int    `json:"valor1"`
	Descripcion1 string `json:"descripcion1"`
}
