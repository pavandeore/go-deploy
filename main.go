package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		var requestBody struct{ Name string }

		if err:= json.NewDecoder(r.Body).Decode(&requestBody); err!=nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := struct{ Message string } { fmt.Sprintf("Hello %s", requestBody.Name) }

		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("server is running")

	log.Fatal(http.ListenAndServe(":8080", nil))
}