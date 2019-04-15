package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"home-accounting/models"
	"home-accounting/month/repository"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/months", createMonth)
	server := &http.Server{
		Handler: router,
		Addr:    "localhost:8000", // trailing litteral coma
	}

	log.Println("running on: ", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil{
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 1000)
	defer cancel()
	// Doesn’t block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func createMonth(w http.ResponseWriter, r *http.Request )  {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "post", "password", "db")

	repo, pgErr := repository.NewPostgres(psqlInfo)
	body, _ := ioutil.ReadAll(r.Body)

	if pgErr != nil {
		println("new pg err:", pgErr)
	}

	var month models.Month
	err := json.Unmarshal(body, &month)

	if err != nil {
		println("unamrshal err:", err)
	}

	repo.Insert(month)

	fmt.Printf("%+v", month)
}
