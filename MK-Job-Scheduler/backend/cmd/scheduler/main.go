package main

import (
	"log"
	"net/http"

	"job-scheduler/internal/handler"
	"job-scheduler/internal/scheduler"
)

func main() {

	s := scheduler.NewScheduler()
	h := handler.NewHandler(s)

	http.HandleFunc("/jobs", h.Jobs)
	http.HandleFunc("/nodes", h.Nodes)
	http.HandleFunc("/register", h.RegisterNode)


	log.Println("Scheduler running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
