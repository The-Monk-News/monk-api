package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/The-Monk-News/monk-api/model"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", getNews)
	c.Start()
}

func getNews() {
	resp, _ := http.Get("https://newsapi.org/v2/top-headlines?country=us&apiKey=fabb056ff8594a2c9cd1ea680aa83aa7")
	defer resp.Body.Close()

	var nresp model.Obj
	if error := json.NewDecoder(resp.Body).Decode(&nresp); error != nil {
		log.Fatal(error)
	}

	fmt.Println(nresp)
}