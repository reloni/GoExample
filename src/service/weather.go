package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	key := "03add4c278668e54404560d4ff48d568"
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", city, key)

	w.Header().Set("Content-Type", "application/json")
	hostName, hostErr := os.Hostname()
	if hostErr == nil {
		w.Header().Set("NodeName", hostName)
	}

	pool := newRedisPool()
	defer pool.Close()
	conn := pool.Get()
	defer conn.Close()

	cached := getWeatherFromRedis(conn, city)
	if cached != nil {
		log.Printf("Load cached weather")
		w.Write([]byte(*cached))
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := fmt.Sprintf("{ \"error\": \"Error calling weather service\", \"message\": \"%s\" }", err.Error())
		w.Write([]byte(message))
		log.Printf("Error: %s", err.Error())
		return
	}

	log.Println("Load remote weather")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	setWeatherToRedis(conn, city, body)

	w.Write(body)
}

func setWeatherToRedis(c redis.Conn, city string, data []byte) {
	setRedisString(c, city, string(data))
	setRedisExpire(c, city, 60*5)
}

func getWeatherFromRedis(c redis.Conn, city string) *string {
	str, _ := getRedisString(c, city)
	if str != nil {
		setRedisExpire(c, city, 60*5)
	}
	return str
}
