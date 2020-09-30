package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Get("https://localhost:5001/WeatherForecast")
	fmt.Println(err)
	if err != nil {
		fmt.Println(resp.StatusCode)
		return
	}

}
