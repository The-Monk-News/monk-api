package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/The-Monk-News/monk-api/model"
	"github.com/robfig/cron"
)

func fetchNewsEveryMinute() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", getNews)
	c.Start()
	time.Sleep(2 * time.Minute)
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