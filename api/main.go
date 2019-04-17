package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/phedirko/go-home-accounting/models"
	"github.com/phedirko/go-home-accounting/month/repository"
)

func main() {

	http.HandleFunc("/api/months", createMonth)
	log.Print("running on :8000 . . .")
	http.ListenAndServe(":8000", nil)
}

func createMonth(w http.ResponseWriter, r *http.Request) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "post", "password", "postgres")

	_, pgErr := repository.NewPostgres(psqlInfo)
	body, _ := ioutil.ReadAll(r.Body)

	if pgErr != nil {
		println("new pg err:", pgErr)
	}

	var month models.Month
	err := json.Unmarshal(body, &month)

	if err != nil {
		println("unamrshal err:", err)
	}

	// repo.Insert(month)

	fmt.Printf("%+v", month)
}
