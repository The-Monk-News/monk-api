package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/The-Monk-News/monk-api/model"
	"github.com/robfig/cron"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func fetchNewsEveryMinute() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", getNews)
	c.Start()
	time.Sleep(2 * time.Minute)
}

func getNews() {
	resp, err := http.Get("https://newsapi.org/v2/top-headlines?country=us&apiKey=fabb056ff8594a2c9cd1ea680aa83aa7")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var nresp model.Obj
	if error := json.NewDecoder(resp.Body).Decode(&nresp); error != nil {
		log.Fatal(error)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://monkBro:FUCKOFF@cluster0.nb9eh.mongodb.net/NewsDB?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	//Time out err
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}
