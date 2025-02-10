package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const apiKey = "3ca4317bb1cbb0a5e37f2ac1b5715e0d"

func fetchWeather(city string, ch chan string, wg *sync.WaitGroup) interface{} {

	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	defer wg.Done()

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching weather for %s: %s\n", city, err)
		return data
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.Printf("Error decoding weather data for %s: %s", city, err)
		return data
	}

	ch <- fmt.Sprintf("This is %s temperature: %v\n", city, data)
	return data
}

func main() {

	startNow := time.Now()

	cities := []string{"Algiers", "London", "Damascus", "Cairo", "Bagdad"}

	ch := make(chan string)

	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go fetchWeather(city, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Printf("This operation took %v\n", time.Since(startNow))

}
