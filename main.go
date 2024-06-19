package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Estructura del JSON de entrada de la API
type InputData struct {
	Date string        `json:"date"`
	Data []interface{} `json:"data"`
}

// Estructura del JSON de salida
type OutputData struct {
	Title string        `json:"titulo"`
	Dia   string        `json:"dia"`
	Info  []interface{} `json:"info"`
	ApiVersion string   `json:"api_version"`
}

// funcion para llamar a la api y consumir el servicio
func apiCall(url string, token string) (InputData, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return InputData{}, err
	}

	// headers de la llamada
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return InputData{}, err
	}
	defer resp.Body.Close()

	var inputData InputData
	err = json.NewDecoder(resp.Body).Decode(&inputData)
	if err != nil {
		return InputData{}, err
	}

	return inputData, nil
}

// Funcion para procesar la solicitud y modificar el JSON para devolver la respuesta solicitada
func handler(w http.ResponseWriter, r *http.Request) {

	// Obtiene las variables de entorno archivo .env
	apiURL := os.Getenv("API_URL")
	authToken := os.Getenv("AUTH_TOKEN")
	if apiURL == "" {
		apiURL = "https://us-central1-teamcore-retail.cloudfunctions.net/test_mobile/api/questions" // valor por defecto por sino puede obtenerlo de la variable de entorno
	}
	if authToken == "" {
		authToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2NzM0NzU4MTEsImV4cCI6MTcwNTAxMTgxMSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.9wqriO_2Q8Xfwc9VcgMpr-2c4WVdLRJ5G6NcNaXdpuY" // valor por defecto por sino puede obtenerlo de la variable de entorno
	}

	inputData, err := apiCall(apiURL, authToken)
	if err != nil {
		http.Error(w, "Failed to fetch API", http.StatusInternalServerError)
		return
	}

	// Construccion del JSON final para devolver
	outputData := OutputData{
		Title: "Test data questions",
		Dia:   time.Now().Format("02-01-2006"),
		Info:  inputData.Data,
		ApiVersion: "1",
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(outputData)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

//Funcion principal que carga al comenzar el programa
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Obtiene el puerto de las variables de entorno archivo .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // valor por defecto por sino puede obtenerlo de la variable de entorno
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
