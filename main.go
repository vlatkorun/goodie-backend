package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have visited %s in `%s`!", r.URL.Path, "Orders")
}

func main() {

	port := func() string {

		value := os.Getenv("HTTP_PORT")

		if len(value) == 0 {
			return "8080"
		}
		return value
	}()

	StartServer(port, handle)
}

func StartServer(port string, handlerFunc http.HandlerFunc) {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting web server on port " + port)
	fmt.Println("Starting web server on port " + os.Getenv("DB_HOST"))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
