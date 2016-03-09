package ds

import (
  "sync"
  "log"
  "bytes"
  "encoding/json"
  "net/http"
)

type DataChannel chan []byte
type Master struct {
  agents []string
}

func (this *Master) ConnectAgent(ip string) {
  this.agents = append(this.agents, ip)
}

func (this *Master) SendStreamCreateToAgents(body []byte) []map[string]interface{} {
  sendBody := bytes.NewReader(body)
  resChan := make(DataChannel, len(this.agents))
  waitGroup := new(sync.WaitGroup)
  responses := make([]map[string]interface{}, len(this.agents))

  // Send a request to each of the agents.
  for _, agent := range this.agents {
    waitGroup.Add(1)
    go func(waitGroup *sync.WaitGroup, agent string, sendBody io.Reader, resChan DataChannel) {
      defer waitGroup.Done()
      res, err := http.Post("http://"+agent+"/create", "application/octet-stream", sendBody)
      if err != nil {
        log.Println("There was an error sending to %s with error %s", agent, err)
      }
      body, err := ioutil.ReadAll(res.Body)
      if err != nil {
        log.Println("There was an error reading a response from %s with error %s", agent, err)
      }
      resChan <- body
    }(waitGroup, agent, sendBody, resChan)
  }
  waitGroup.Wait()

  // Serialize the responses
  i := 0
  for response := range resChan {
    jsonResponse := make(map[string]interface{})
    if err := json.Unmarshal(response, &jsonResponse); err != nil {
      log.Println("There was an error serializing a response from %s with error %s", agent, err)
    }
    responses[i] = jsonResponse
    i++
  }
  return responses
}

func (this *Master) SendStreamDataToAgents(body []byte) map[string]interface{} {
  sendBody := bytes.NewReader(body)
  resChan := make(DataChannel, len(this.agents))
  waitGroup := new(sync.WaitGroup)
  responses := make([]map[string]interface{}, len(this.agents))

  // Send a request to each of the agents.
  for _, agent := range this.agents {
    waitGroup.Add(1)
    go func(waitGroup *sync.WaitGroup, agent string, sendBody io.Reader, resChan DataChannel) {
      defer waitGroup.Done()
      res, err := http.Post("http://"+agent+"/send", "application/octet-stream", sendBody)
      if err != nil {
        log.Println("There was an error sending to %s with error %s", agent, err)
      }
      body, err := ioutil.ReadAll(res.Body)
      if err != nil {
        log.Println("There was an error reading a response from %s with error %s", agent, err)
      }
      resChan <- body
    }(waitGroup, agent, sendBody, resChan)
  }
  waitGroup.Wait()

  // Serialize the responses
  i := 0
  for response := range resChan {
    jsonResponse := make(map[string]interface{})
    if err := json.Unmarshal(response, &jsonResponse); err != nil {
      log.Println("There was an error serializing a response from %s with error %s", agent, err)
    }
    responses[i] = jsonResponse
    i++
  }
  return responses
}

func (this *Master) SendStreamClusterToAgents() map[string]interface{} {
  waitGroup := new(sync.WaitGroup)
  resChan := make(DataChannel, len(this.agents))
  responses := make([]map[string]interface{}, len(this.agents))
  for _, agent := range this.agents {
    waitGroup.Add(1)
    go func(waitGroup *sync.WaitGroup, agent string, resChan DataChannel) {
      defer waitGroup.Done()
      res, err := http.Get("http://"+agent+"/cluster")
      if err != nil {
        log.Println("There was an error sending to %s with error %s", agent, err)
      }
      body, err := ioutil.ReadAll(res.Body)
      if err != nil {
        log.Println("There was an error reading a response from %s with error %s", agent, err)
      }
      resChan <- res
    }(waitGroup, agent, resChan)
  }
  waitGroup.Wait()

  // Serialize the responses
  i := 0
  for response := range resChan {
    jsonResponse := make(map[string]interface{})
    if err := json.Unmarshal(response, &jsonResponse); err != nil {
      log.Println("There was an error serializing a response from %s with error %s", agent, err)
    }
    responses[i] = jsonResponse
    i++
  }
  return responses
}
