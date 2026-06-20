package scheduler

import (
	"sort"
	"sync"
	"time"

	"job-scheduler/internal/model"
)

type Scheduler struct {
	Jobs  []model.Job
	Nodes []model.Node
	Mu    sync.Mutex
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		Jobs: []model.Job{},
		Nodes: []model.Node{
			{ID: "node-1", TotalCPU: 4, TotalMem: 1024},
			{ID: "node-2", TotalCPU: 2, TotalMem: 512},
			{ID: "node-3", TotalCPU: 8, TotalMem: 2048},
		},
	}
}

func (s *Scheduler) ScheduleJob(job *model.Job) bool {
	for i := range s.Nodes {
		node := &s.Nodes[i]

		if node.TotalCPU-node.UsedCPU >= job.CPU &&
			node.TotalMem-node.UsedMemory >= job.Memory {

			node.UsedCPU += job.CPU
			node.UsedMemory += job.Memory
			job.AssignedNode = node.ID
			job.Status = "RUNNING"
			return true
		}
	}
	job.Status = "PENDING"
	return false
}

func (s *Scheduler) AddJob(job model.Job) model.Job {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	if s.ScheduleJob(&job) {
		s.Jobs = append(s.Jobs, job)
		s.startExecution(len(s.Jobs) - 1)
	} else {
		s.Jobs = append(s.Jobs, job)
	}

	return job
}

func (s *Scheduler) GetJobs() []model.Job {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	return s.Jobs
}

func (s *Scheduler) GetNodes() []model.Node {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	return s.Nodes
}

func (s *Scheduler) startExecution(index int) {
	go func(i int) {
		time.Sleep(10 * time.Second)

		s.Mu.Lock()
		defer s.Mu.Unlock()

		if i >= len(s.Jobs) {
			return
		}

		if s.Jobs[i].Status == "RUNNING" {

			for n := range s.Nodes {
				if s.Nodes[n].ID == s.Jobs[i].AssignedNode {
					s.Nodes[n].UsedCPU -= s.Jobs[i].CPU
					s.Nodes[n].UsedMemory -= s.Jobs[i].Memory
				}
			}

			s.Jobs[i].Status = "COMPLETED"
			s.reschedulePending()
		}
	}(index)
}

func (s *Scheduler) reschedulePending() {

	var pendingIndexes []int
	for i := range s.Jobs {
		if s.Jobs[i].Status == "PENDING" {
			pendingIndexes = append(pendingIndexes, i)
		}
	}

	sort.Slice(pendingIndexes, func(a, b int) bool {
		return s.Jobs[pendingIndexes[a]].Priority >
			s.Jobs[pendingIndexes[b]].Priority
	})

	for _, i := range pendingIndexes {
		if s.ScheduleJob(&s.Jobs[i]) {
			s.startExecution(i)
		}
	}
}
