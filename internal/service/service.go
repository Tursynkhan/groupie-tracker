package service

import (
	"encoding/json"
	"log"
	"main/internal/models"
	"net/http"
)

type Service struct{}

func (s *Service) Group() ([]models.Artists, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	var artist []models.Artists

	client := http.Client{}

	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&artist); err != nil {
		log.Fatal(err)
	}
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// jsonErr := json.Unmarshal(body, &artist)
	// if jsonErr != nil {
	// 	log.Fatal(jsonErr)
	// }
	return artist, nil
}

func (s *Service) IdArtist(url string) (models.Artists, error) {
	concatUrl := "https://groupietrackers.herokuapp.com/api/artists/" + url
	var artist models.Artists
	client := http.Client{}
	res1, err := client.Get(concatUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res1.Body.Close()

	if err := json.NewDecoder(res1.Body).Decode(&artist); err != nil {
		log.Fatal(err)
	}

	return artist, nil
}

func (s *Service) Relations(url string) (models.Relations, error) {
	concatUrl := "https://groupietrackers.herokuapp.com/api/relation/" + url
	var artist models.Relations
	client := http.Client{}
	res, err := client.Get(concatUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&artist); err != nil {
		log.Fatal(err)
	}

	return artist, nil
}
