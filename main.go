package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Abhinav7903/mongo/routes"
)

func main() {
	fmt.Println("Hello, World!")
	r:=routes.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server is running on port 8080")
}
