package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/golang/glog"
)

func main() {
	fmt.Println(os.Args[1])
	key := "03add4c278668e54404560d4ff48d568"
	city := os.Args[1]
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", city, key)
	resp, err := http.Get(url)

	if err != nil {
		glog.Error(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Body error\n")
		return
	}

	fmt.Printf("%s", body)

	glog.Flush()
}
