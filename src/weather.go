package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
)

func getWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	key := "03add4c278668e54404560d4ff48d568"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", city, key)

	w.Header().Set("Content-Type", "application/json")

	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := fmt.Sprintf("{ \"error\": \"Error calling weather service\", \"message\": \"%s\" }", err.Error())
		w.Write([]byte(message))
		glog.Error(err.Error())
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	w.Write(body)
}
