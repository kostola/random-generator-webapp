package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strconv"

	"github.com/Pallinder/go-randomdata"
	"github.com/gorilla/mux"
)

type Response struct {
	Status  string   `json:"status"`
	Data    []string `json:"data,omitempty"`
	Message string   `json:"message,omitempty"`
}

var validItems = []string{
	"adjective",
	"address",
	"city",
	"country",
	"currency",
	"day",
	"email",
	"firstnamem",
	"firstnamef",
	"fulldate",
	"fullnamem",
	"fullnamef",
	"ipv4address",
	"ipv6address",
	"lastname",
	"locale",
	"month",
	"paragraph",
	"phonenumber",
	"sillyname",
	"state",
	"street",
	"titlem",
	"titlef",
	"useragent",
}

func getRandomItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item := vars["item"]

	// Check if the "item" is in the valid list
	if !isValidItem(item) {
		// Return a Bad Request error response
		response := Response{
			Status:  "error",
			Data:    nil,
			Message: "Item is not valid.",
		}

		if err := sendJSONResponse(w, response, http.StatusBadRequest); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Get the "count" query parameter or set it to 1 if not present
	countStr := r.URL.Query().Get("count")
	count := 1
	if len(countStr) > 0 {
		countInt, err := strconv.Atoi(countStr)
		if err != nil || countInt < 1 || countInt > 10 {
			// Return an error response
			response := Response{
				Status:  "error",
				Data:    nil,
				Message: "Invalid count parameter. Count must be between 1 and 5.",
			}

			if err := sendJSONResponse(w, response, http.StatusBadRequest); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		count = countInt
	}

	// Create a response struct for a successful response
	data := []string{}
	for i := 0; i < count; i++ {
		value := ""
		switch item {
		case "adjective":
			value = randomdata.Adjective()
		case "address":
			value = randomdata.Address()
		case "city":
			value = randomdata.City()
		case "country":
			value = randomdata.Country(randomdata.FullCountry)
		case "currency":
			value = randomdata.Currency()
		case "day":
			value = randomdata.Day()
		case "email":
			value = randomdata.Email()
		case "firstnamem":
			value = randomdata.FirstName(randomdata.Male)
		case "firstnamef":
			value = randomdata.FirstName(randomdata.Female)
		case "fulldate":
			value = randomdata.FullDate()
		case "fullnamem":
			value = randomdata.FullName(randomdata.Male)
		case "fullnamef":
			value = randomdata.FullName(randomdata.Female)
		case "ipv4address":
			value = randomdata.IpV4Address()
		case "ipv6address":
			value = randomdata.IpV6Address()
		case "lastname":
			value = randomdata.LastName()
		case "locale":
			value = randomdata.Locale()
		case "month":
			value = randomdata.Month()
		case "paragraph":
			value = randomdata.Paragraph()
		case "phonenumber":
			value = randomdata.PhoneNumber()
		case "sillyname":
			value = randomdata.SillyName()
		case "state":
			value = randomdata.State(randomdata.FullCountry)
		case "street":
			value = randomdata.Street()
		case "titlem":
			value = randomdata.Title(randomdata.Male)
		case "titlef":
			value = randomdata.Title(randomdata.Female)
		case "useragent":
			value = randomdata.UserAgentString()
		}
		data = append(data, value)
	}

	if err := sendJSONResponse(w, Response{Status: "ok", Data: data}, http.StatusOK); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func isValidItem(item string) bool {
	for _, validItem := range validItems {
		if validItem == item {
			return true
		}
	}
	return false
}

func sendJSONResponse(w http.ResponseWriter, response Response, statusCode int) error {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, err = w.Write(jsonResponse)
	return err
}

func startServer(port int) error {
	r := mux.NewRouter()
	r.HandleFunc("/random/{item}", getRandomItem).Methods("GET")

	http.Handle("/", r)

	// Start the server and handle the error
	fmt.Printf("Server is starting on port %d...\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getPort() (int, error) {
	viper.AutomaticEnv()
	viper.SetConfigName("server")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	var portFromFlag int
	pflag.IntVar(&portFromFlag, "port", 0, "Port number")
	pflag.Parse()

	if err := viper.BindPFlag("port", pflag.Lookup("port")); err != nil {
		return 0, err
	}

	if err := viper.ReadInConfig(); err != nil {
		return 0, err
	}

	return viper.GetInt("port"), nil
}

func main() {
	port, err := getPort()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = startServer(port); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
