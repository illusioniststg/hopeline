package main

import (
	"fmt"
	"net/http"
	"github.com/illusioniststg/hopeline/internal/routeplanning"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Project Hopeline!")
}

func main() {
	http.HandleFunc("/assign_boat", routeplanning.AssignBoat)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
