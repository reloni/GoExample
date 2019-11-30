package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getWeather(w http.ResponseWriter, r *http.Request) {
	log.Println("Load weather")

	city := r.URL.Query().Get("city")
	key := "03add4c278668e54404560d4ff48d568"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", city, key)

	// pool := newPool()
	// conn := pool.Get()
	// str, err := getString(conn, city)
	// if err == nil {
	// 	log.Printf("Read string: %s", *str)
	// }
	// setString(conn, city, "string")

	w.Header().Set("Content-Type", "application/json")

	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := fmt.Sprintf("{ \"error\": \"Error calling weather service\", \"message\": \"%s\" }", err.Error())
		w.Write([]byte(message))
		log.Printf("Error: %s", err.Error())
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	w.Write(body)
}
