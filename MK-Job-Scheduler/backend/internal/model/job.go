package model

type Job struct {
	ID       string `json:"id"`
	CPU      int    `json:"cpu"`
	Memory   int    `json:"memory"`
	Priority int    `json:"priority"`
	Status   string `json:"status"`
        AssignedNode string `json:"assigned_node"`
}
