package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/ngs/go-amazon-product-advertising-api/amazon"
	"fmt"
	"encoding/json"
)

var items []Item

type Item struct {
	Title string
	URL string
}

// Display all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	client, err := amazon.New("**", "**", "ngsio-22", amazon.RegionIndia)
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.ItemSearch(amazon.ItemSearchParameters{
		SearchIndex: amazon.SearchIndexAll,
		Keywords:    "watch",
	}).Do()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d results found\n\n", res.Items.TotalResults)
	for _, item := range res.Items.Item {
		fmt.Printf(`-------------------------------
[Title] %v
[URL]   %v
`, item.ItemAttributes.Title, item.DetailPageURL)
	}
	json.NewEncoder(w).Encode(res.Items)
}

func main() {
	// set Router
	router := mux.NewRouter()
	router.Handle("/", http.FileServer(http.Dir("./view")))
	router.HandleFunc("/items", GetItems).Methods("GET")
	http.ListenAndServe(":3030", router)
}
