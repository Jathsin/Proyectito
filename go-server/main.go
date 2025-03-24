package main

import (
	"encoding/json"
	"fmt"
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

func createMux() *http.ServeMux {

	mux := http.NewServeMux()

	//Specify handlers

	//ANTARTIDA
	mux.HandleFunc("/antartida", func(w http.ResponseWriter, r *http.Request) { // Makes request, Unmarshalls data and prints it back as an image
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		//@TODO: do not use default client!
		data, err := fetchAPIData[*Antartida]("API_ANTARTIDA", http.DefaultClient)
		if err != nil {
			http.Error(w, "Failed to fetch dara from API", http.StatusInternalServerError)
		}

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

// Makes API reusable using generics
func fetchAPIData[T any](apiKey string, client *http.Client) (*T, error) {

	var apiURL = os.Getenv(apiKey)
	if apiURL == "" {
		return nil, fmt.Errorf("Could not access API Key")
	}

	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch data from API: %w", err)
	}

	//Show response to the client
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not access Body: %w", err)
	}

	var data T
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("JSON could not be unmarshalled: %w", err)
	}
	return &data, nil
}
