// Agent Node
package main

import (
  "os"
  "log"
  "github.com/jessevdk/go-flags"
  "github.com/wilseypa/rphash-golang/app/ds"
  "github.com/wilseypa/rphash-golang/app/server"
)

type AgentCommand struct {
  Master string `short:"m" long:"master" optional:"no" description:"Provide the IP of master server"`
}

var (
  agent = new(ds.Agent)
  agentCommand AgentCommand
  parser = flags.NewParser(nil, flags.Default)
)

func (this *AgentCommand) Execute(args []string) error {
  agent.ConnectMaster(this.Master)
  return nil
}

func main() {
  parser.AddCommand("master",
    "Add the IP of master",
    "Add the IP of the master server",
    &agentCommand)

  if _, err := parser.Parse(); err != nil {
    os.Exit(1)
  }

  // Get the config for the server.
  instance, _ := server.Server()

  // Set up routes.
  router := make(map[string]server.HTTPHandler)
  router["POST†/create"] = PostCreateStream
  router["POST†/send"] = PostSendStream
  router["GET/cluster"] = GetClusterStream
  instance.Server.Router = router

  // Server listening ...
  instance.Server.Listen()
}

func PostCreateStream(res http.ResponseWriter, req *http.Request) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    log.Println("Failed to parse incoming request with error:", err)
  }
  if agent.IsInitialized() {
    res.Write([]byte(`{ "error": "Stream could not be initialized. Stream already created" }`))
    return
  }
  agent.HandleCreateStream(body)
  res.Write([]byte(`{ "response": "Stream successfully created" }`))
}

func PostSendStream(res http.ResponseWriter, req *http.Request) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    log.Println("Failed to parse incoming request with error:", err)
  }
  if !agent.IsInitialized() {
    res.Write([]byte(`{ "error": "Stream does not exist" }`))
    return
  }
  agent.HandleSendDataStream(body)
  res.Write([]byte(`{ "response": "Vector successfully appended to stream" }`))
}

func GetClusterStream(res http.ResponseWriter, req *http.Request) {
  if !agent.IsInitialized() {
    res.Write([]byte(`{ "error": "Stream does not exist" }`))
    return
  }
  response := agent.HandleClusterStream()
  // ...
}
