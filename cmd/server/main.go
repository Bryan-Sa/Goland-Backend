package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Bryan-Sa/Goland-Backend/api/route"
)

func main() {
	fmt.Println("Starting the server...")
	log.Fatal(http.ListenAndServe(":8080", route.SetupRoutes()))
	route.Tester()
}
