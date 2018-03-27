package main
import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./view")))
	http.ListenAndServe(":3030", nil)
}