package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Code de configuration et d'initialisation de votre application

	// Exemple de gestionnaire d'API pour une route /hello
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	// Démarrer le serveur HTTP sur le port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
