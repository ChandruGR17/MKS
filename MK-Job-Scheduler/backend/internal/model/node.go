package model

type Node struct {
	ID         string `json:"id"`
	TotalCPU   int    `json:"total_cpu"`
	TotalMem   int    `json:"total_memory"`
	UsedCPU    int    `json:"used_cpu"`
	UsedMemory int    `json:"used_memory"`
}
