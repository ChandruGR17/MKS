package handler

import (
	"encoding/json"
	"net/http"

	"job-scheduler/internal/model"
	"job-scheduler/internal/scheduler"
)

type Handler struct {
	Scheduler *scheduler.Scheduler
}

func NewHandler(s *scheduler.Scheduler) *Handler {
	return &Handler{Scheduler: s}
}

func (h *Handler) Jobs(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		var job model.Job
		err := json.NewDecoder(r.Body).Decode(&job)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result := h.Scheduler.AddJob(job)
		json.NewEncoder(w).Encode(result)

	case http.MethodGet:
		json.NewEncoder(w).Encode(h.Scheduler.GetJobs())

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Nodes(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	json.NewEncoder(w).Encode(h.Scheduler.GetNodes())
}
