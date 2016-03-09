package ds

import (
  "log"
)

type Agent struct {
  master string
}

func (this *Agent) ConnectMaster(ip string) {
  this.master = ip
}

func (this *Agent) HandleCreateStream(cluster int) {
  log.Println(cluster)
}

func (this *Agent) HandleSendStream(data interface{}) {
  log.Println(data)
}

func (this *Agent) HandleClusterStream() [][]float64 {
  return make([][]float64, 5)
}
