package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/shadowdevfr/naobluesky/bluesky"
	"github.com/shadowdevfr/naobluesky/mastodon"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	log.Println("Requête à APIFY...")
	var start time.Time = time.Now()
	requestBody, err := json.Marshal(ApifyRequest{
		From:      "naolib_direct",
		MaxItems:  20,
		QueryType: "Latest",
	})
	if err != nil {
		log.Println(err)
		return
	}

	apifyToken := os.Getenv("APIFY_TOKEN")
	if rand.Intn(2) == 1 {
		apifyToken = os.Getenv("APIFY_TOKEN2")
		log.Println("Utilisation d'APIFY 2")
	} else {
		log.Println("Utilisation d'APIFY 1")
	}

	resp, err := http.Post(fmt.Sprintf("https://api.apify.com/v2/acts/CJdippxWmn9uRfooo/runs?token=%s", apifyToken), "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println(err)
		return
	}
	respstr, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var response ApifyResponse
	err = json.Unmarshal(respstr, &response)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("[APIFY] Dataset: %s", response.Data.DefaultDatasetID)
	if response.Data.DefaultDatasetID == "" {
		panic("Empty dataset: apify limit reached?")
	}
	log.Println("Attente d'une réponse...")
	time.Sleep(15 * time.Second)

	resp, err = http.Get(fmt.Sprintf("https://api.apify.com/v2/datasets/%s/items?token=%s", response.Data.DefaultDatasetID, apifyToken))
	if err != nil {
		log.Println(err)
		return
	}
	respstr, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var tweets []Tweet
	err = json.Unmarshal(respstr, &tweets)
	if err != nil {
		log.Println(err)
		return
	}

	if len(tweets) <= 0 {
		return
	}

	log.Printf("Tweet récupéré en %s.", time.Since(start))
	log.Printf("Vérification du dernier tweet stocké, si ce n'est pas celui ci, alors on poste")

	lasttweet, err := os.ReadFile("lasttweet.txt")
	if err != nil {
		panic(err)
	}

	if strings.HasPrefix(tweets[0].Text, "@") { // Surement une réponse à un tweet, on ignore
		log.Println("Le tweet est une réponse à quelqu'un, on ignore.")
		os.Exit(0)
		return
	}

	if string(lasttweet) == tweets[0].ID {
		log.Println("C'est le même, on peut ignorer. bye")
		os.Exit(0)
		return
	}

	err = os.WriteFile("lasttweet.txt", []byte(tweets[0].ID), 0644)
	if err != nil {
		panic(err)
	}

	session := bluesky.CreateSession(os.Getenv("BSKY_NAME"), os.Getenv("BSKY_PASSWORD"))
	bluesky.Post(session, tweets[0].Text)
	log.Println("Posté sur Bluesky")

	ma := mastodon.Mastodon{Instance: os.Getenv("MASTODON_INSTANCE"), Token: os.Getenv("MASTODON_TOKEN")}
	ma.Post(tweets[0].Text)
	log.Printf("Posté sur Mastodon: %s\n", os.Getenv("MASTODON_INSTANCE"))
}
