package routes

import (
	"github.com/andre-haniak/go-api/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/Personalities", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
