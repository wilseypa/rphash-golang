package ds

import (
  "log"
  "github.com/wilseypa/rphash-golang/stream"
  "github.com/wilseypa/rphash-golang/api"
)

type Agent struct {
  master string
  initialized bool
  streaming *stream.Stream
}

func (this *Agent) IsInitialized() bool {
  return this.initialized
}

func (this *Agent) ConnectMaster(ip string) {
  this.master = ip
}

func (this *Agent) HandleCreateStream(body []byte) {
  this.initialized = true
  data := make(map[string]interface{})
  if err := json.Unmarshal(body, &data); err != nil {
    log.Println("There was an error serializing a response with error %s", err)
  }

  clusters := int(data["clusters"].(float64))
  dimensionality := int(data["dimensionality"].(float64))
  this.streaming = api.NewStreamRPHash(dimensionality,  clusters)
}

func (this *Agent) HandleSendDataStream(body []byte) {
  // Get the vector from the body
  // Append it to the stream
}

func (this *Agent) HandleClusterStream(body []byte) {
  // Run the clusterer
  // Get centroids and return them
}
