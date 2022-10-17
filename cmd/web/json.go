package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func group() []Artists {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var artist []Artists
	jsonErr := json.Unmarshal(body, &artist)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return artist
}

func idArtist(url string) Artists {
	concatUrl := "https://groupietrackers.herokuapp.com/api/artists/" + url
	res1, err := http.Get(concatUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res1.Body.Close()

	body, err := ioutil.ReadAll(res1.Body)
	if err != nil {
		log.Fatal(err)
	}
	var artist Artists
	jsonErr := json.Unmarshal(body, &artist)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return artist
}

func relations(url string) Relations {
	concatUrl := "https://groupietrackers.herokuapp.com/api/relation/" + url

	res, err := http.Get(concatUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var artist Relations
	jsonErr := json.Unmarshal(body, &artist)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return artist
}
