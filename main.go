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
	client, err := amazon.NewFromEnvionment()
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.ItemSearch(amazon.ItemSearchParameters{
		SearchIndex: amazon.SearchIndexBooks,
		Keywords:    "Go 言語",
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
	json.NewEncoder(w).Encode(items)
}

func main() {
	// set Router
	router := mux.NewRouter()
	router.Handle("/", http.FileServer(http.Dir("./view")))
	router.HandleFunc("/items", GetItems).Methods("GET")
	http.ListenAndServe(":3030", router)
}


// Configuration
//export AWS_ACCESS_KEY_ID=${YOUR_AWS_ACCESS_KEY_ID}
//export AWS_SECRET_ACCESS_KEY=${YOUR_AWS_SECRET_ACCESS_KEY}
//export AWS_PRODUCT_REGION=JP
//export AWS_ASSOCIATE_TAG=ngsio-22