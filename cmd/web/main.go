package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Artists struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int32
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	ts, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	res := group()
	// fmt.Println(res)
	err = ts.Execute(w, res)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func group() []Artists {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))
	if err != nil {
		log.Fatal(err)
	}
	var artist []Artists
	jsonErr := json.Unmarshal(body, &artist)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return artist
	// fmt.Println(artist)
}
