// Agent Node
package main

import (
  "os"
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
  router["POST†/create"] = HandleCreateStream
  router["POST†/send"] = HandleSendStream
  router["POST†/cluster"] = HandleClusterStream
  instance.Server.Router = router

  // Server listening ...
  instance.Server.Listen()
}

func HandleCreateStream(res http.ResponseWriter, req *http.Request) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    log.Println("Failed to parse incoming request with error:", err)
  }
  agent.HandleCreateStream(/* ... */)
  res.Write([]byte("200"))
}

func HandleSendStream(res http.ResponseWriter, req *http.Request) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    log.Println("Failed to parse incoming request with error:", err)
  }

  agent.HandleSendStream(/* ... */)
  res.Write([]byte("200"))
}

func HandleClusterStream(res http.ResponseWriter, req *http.Request) {
  clusters := agent.HandleClusterStream()
  // Send the data back...
}
