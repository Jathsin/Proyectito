package aemet_api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var base_url string = "https://opendata.aemet.es/opendata"

var uv_prediction_endpoint string = "/api/prediccion/especifica/uvi/"

func get_uv_prediction(dia int, id string) (int, error) {

	err := godotenv.Load()
	if err != nil {
		return -1, err
	}

	api_key := os.Getenv("aemet_api_key")

	//Create client and set header with api key
	client := &http.Client{}
	url := base_url + uv_prediction_endpoint + strconv.Itoa(dia)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1, err
	}
	req.Header.Add("api_key", api_key)

	//Get first response with data url
	var data_url_response = data_url_response{}
	res, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return -1, err
	}
	err = json.Unmarshal(body, &data_url_response)
	if err != nil {
		return -1, err
	}

	//Get second response with actual data
	data_url := data_url_response.Datos
	data_response, err := client.Get(data_url)
	if err != nil {
		return -1, err
	}
	data_response_body, _ := io.ReadAll(data_response.Body)
	var data_response_struct = Data_response_data{}
	err = json.Unmarshal(data_response_body, &data_response_struct)
	if err != nil {
		return -1, err
	}

	//Search location in data
	for _, ciudad := range data_response_struct.Root.Ciudad {
		if ToStringValue(ciudad.Id) == id {
			return ciudad.Uv, nil
		}
	}
	return -1, errors.New("Dato no encontrado")

}

var uv_local_prediction_endpoint = "/api/prediccion/especifica/playa/"


func get_uv_local_prediction(id_playa int) ([]int, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	api_key := os.Getenv("aemet_api_key")

	//Create client and set header with api key
	client := &http.Client{}
	url := base_url + uv_local_prediction_endpoint + strconv.Itoa(id_playa)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("api_key", api_key)

	//First call
	var data_url_response = data_url_response{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &data_url_response)
	if err != nil {
		return nil, err
	}

	//Second call
	data_url := data_url_response.Datos
	data_response, err := client.Get(data_url)
	if err != nil {
		return nil, err
	}
	data_response_body, err := io.ReadAll(data_response.Body)
	if err != nil {
		return nil, err
	}
	var data_response_array Data_response_data_local
	err = json.Unmarshal(data_response_body, &data_response_array)
	if err != nil {
		return nil, err
	}

	//Search through the response
	var uv_array []int

	for _, beachForecast := range data_response_array {
		for _, day := range beachForecast.Prediccion.Dia {
			uv_array = append(uv_array, day.UvMax.Valor1)
		}

	}

	return uv_array, nil
}
