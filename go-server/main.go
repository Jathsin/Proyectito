//Builds server

package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	mux := createMux()

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	err = server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}

}

type Data struct {
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

func createMux() *http.ServeMux {

	mux := http.NewServeMux()

	//Specify handlers

	//ANTARTIDA
	mux.HandleFunc("/antartida", func(w http.ResponseWriter, r *http.Request) { // Makes request, Unmarshalls data and prints it back as an image
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		//Call AEMET API
		apiURL := os.Getenv("API_ANTARTIDA")
		if apiURL == "" {
			http.Error(w, "API URL not set", http.StatusInternalServerError)
			return
		}

		resp, err := http.Get(apiURL)
		if err != nil {
			http.Error(w, "Failed request to AMEMET´s API", http.StatusInternalServerError)
			return
		}

		//Show response to the client
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}

		var data Data
		//@Todo:should manage error here!
		json.Unmarshal(body, &data)

		w.Header().Set("Content-Type", "json")

		responseData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(responseData)
		if err != nil {
			http.Error(w, "Failed to write body", http.StatusInternalServerError)
			return
		}
	})
	return mux

	//Should manage error here? What if the error is in the implementation of the mux?

}
