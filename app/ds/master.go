package ds

import (
  "log"
)

type Master struct {
  agents []string
}

func (this *Master) ConnectAgent(ip string) {
  this.agents = append(this.agents, ip)
}

func (this *Master) SendStreamCreateToAgents(clusters int) {
  log.Println(clusters)
}

func (this *Master) SendStreamDataToAgents(data interface{}) {
  log.Println(data)
}

func (this *Master) SendStreamClusterToAgents() [][]float64 {
  // Wait for responses...

  return make([][]float64, 5)
}
