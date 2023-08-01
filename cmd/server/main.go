package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Bryan-Sa/Goland-Backend/api/route"
	"github.com/Bryan-Sa/Goland-Backend/db"
)

func main() {
	db.DbConnection()
	fmt.Println("Starting the server...")
	fmt.Println("je suis un petit test")
	log.Fatal(http.ListenAndServe(":8080", route.SetupRoutes()))
	route.Tester()
}
