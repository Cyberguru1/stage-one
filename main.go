////////////////////////////
// Code by Cyb3rGuru      //
// APi server             //
//                        //
///////////////////////////

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Details struct {
	Client_ip string `json:"client_ip"`
	Location  string `json:"location"`
	Greeting  string `json:"greeting"`
}

var WEATHER_API_KEY string
var LOC_API_KEY string

func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// passed details
	user_name := r.URL.Query().Get("visitor_name")
	client_ip := r.Header.Get("X-Forwarded-For")

	var detail Details
	myloc := [2]float64{}

	// query through api
	location := GetLoc(&myloc, client_ip)
	tempt := GetTempt(myloc)

	// populate detail struct
	detail.Greeting = fmt.Sprint("Hello, ",user_name,"!, the temperature is ",tempt," degrees Celcius in ",location)
	detail.Client_ip = client_ip
	detail.Location = location

	json.NewEncoder(w).Encode(detail)
	return
}

func GetLoc(coord *[2]float64, ip_address string) string {

	if strings.Contains(ip_address, "::") || strings.Contains(ip_address, "127.0.0.1"){
		return "localhost"
	}

	url := fmt.Sprintf("https://api.geoapify.com/v1/ipinfo?ip=%s&apiKey=%s", ip_address, LOC_API_KEY)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return ""
	}
	res, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var response interface{}

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println(err)
	}

	output_city := response.(map[string]interface{})["city"].(map[string]interface{})["name"].(string)
	output_location := response.(map[string]interface{})["location"].(map[string]interface{})

	coord[0], coord[1] = output_location["latitude"].(float64), output_location["longitude"].(float64)

	return output_city
}

func GetTempt(location [2]float64) string {

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric", location[0], location[1], WEATHER_API_KEY)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return ""
	}
	res, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var response interface{}

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println(err)
	}

	tempt := fmt.Sprint(response.(map[string]interface{})["main"].(map[string]interface{})["temp"].(float64))

	return tempt
}

func main() {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Print("Error loading .env file")
	}

	WEATHER_API_KEY, LOC_API_KEY = os.Getenv("WEATHER_API_KEY"), os.Getenv("LOC_API_KEY")

	r := mux.NewRouter()
	r.HandleFunc("/api/hello", getInfo).Methods("GET")

	fmt.Printf("Starting server at port:8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
	
}
