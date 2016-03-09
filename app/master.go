// Master Node
package main

import (
  "github.com/jessevdk/go-flags"
  "encoding/json"
  "github.com/go-zoo/bone"
  "github.com/wilseypa/rphash-golang/app/server"
  "github.com/wilseypa/rphash-golang/app/ds"
  "io/ioutil"
  "log"
  "os"
  "net/http"
)

type MasterCommand struct {
  Agents []string `short:"a" long:"agents" optional:"no" description:"Provide the IP of server(s) in cluster"`
}

var (
  master = new(ds.Master)
  masterCommand MasterCommand
  parser = flags.NewParser(nil, flags.Default)
)

func (this *MasterCommand) Execute(args []string) error {
  for _, node := range args {
    master.ConnectAgent(node)
  }
  return nil
}

func main() {
  parser.AddCommand("agents",
    "Add a list of agents",
    "Add a list of IPs associated with agents in cluster",
    &masterCommand)

  if _, err := parser.Parse(); err != nil {
    os.Exit(1)
  }

  // Get the config for the server.
  instance, _ := server.Server()

  // Set up routes.
  router := make(map[string]server.HTTPHandler)

  // Client application.
  router["GET†/"] = GetRoot
  router["GET†/:file"] = GetPublic

  // Server application.
  router["POST†/api/stream/create"] = PostCreateStream
  router["POST†/api/stream/send"] = PostSendStream
  router["GET†/api/stream/cluster"] = GetClusterStream
  instance.Server.Router = router

  // Server listening ...
  instance.Server.Listen()
}

func GetRoot(res http.ResponseWriter, req *http.Request) {
  http.ServeFile(res, req, "public/index.html")
}

func GetPublic(res http.ResponseWriter, req *http.Request) {
  file := bone.GetValue(req, "file")
  http.ServeFile(res, req, "public/"+file)
}

func PostCreateStream(res http.ResponseWriter, req *http.Request) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    log.Println("Failed to parse incoming request with error:", err)
  }
  finalResponse := make(map[string][]map[string]interface{})
  finalResponse["data"] := master.SendStreamCreateToAgents(body)
  finalJSONResponse, err := json.Marshal(finalResponse)
  if err != nil {
    log.Println("Failed to parse outgoing request with error:", err)
  }
  res.Write([]byte(string(finalJSONResponse)))
}

func PostSendStream(res http.ResponseWriter, req *http.Request) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    log.Println("Failed to parse incoming request with error:", err)
  }
  finalResponse := make(map[string][]map[string]interface{})
  finalResponse["data"] := master.SendStreamDataToAgents(body)
  finalJSONResponse, err := json.Marshal(finalResponse)
  if err != nil {
    log.Println("Failed to parse outgoing request with error:", err)
  }
  res.Write([]byte(string(finalJSONResponse)))
}

func GetClusterStream(res http.ResponseWriter, req *http.Request) {
  finalResponse := make(map[string][]map[string]interface{})
  finalResponse["data"] := master.SendStreamClusterToAgents()
  finalJSONResponse, err := json.Marshal(finalResponse)
  if err != nil {
    log.Println("Failed to parse outgoing request with error:", err)
  }
  res.Write([]byte(string(finalJSONResponse)))
}
