package main

import (
  "errors"
  "encoding/json"
  "github.com/go-zoo/bone"
  _ "github.com/chrislusf/glow/driver"
  "github.com/chrislusf/glow/flow"
  "github.com/wilseypa/rphash-golang/app/server"
  "github.com/wilseypa/rphash-golang/api"
  "github.com/wilseypa/rphash-golang/stream"
  "io/ioutil"
  "log"
  "net/http"
)

type VectorChannel chan [][]float64

type Service struct {
  stream *stream.Stream
}

type Request struct {
  data     []float64 `json:"data"`
  clusters int       `json:"clusters"`
}

type Response struct {
  clusters [][]float64 `json:"data"`
}

var (
  streamingService = new(Service)
	f = flow.New()
)

func main() {
  // Get the config for the server.
  instance, _ := server.Server()

  // Set up routes.
  router := make(map[string]server.HTTPHandler)
  router["GET†/"] = GetRoot
  router["GET†/:file"] = GetPublic
  router["POST†/stream"] = PostDataStream
  router["POST†/map"] = PostMapStream
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

func PostDataStream(res http.ResponseWriter, req *http.Request) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    log.Println("Failed to parse incoming request with error:", err)
  }
  var data Request
  if err := json.Unmarshal(body, &data); err != nil {
    log.Println("Failed to parse json with error:", err)
  }

  if streamingService.stream == nil {
    streamingService.stream = api.NewStreamRPHash(len(data.data), int(data.clusters))
  } else {
    streamingService.stream.AppendVector(data.data)
  }

  if err != nil {
    log.Println("Failed to stream data with error:", err)
  }
  log.Println("Stream updated.")
  res.Write([]byte(200))
}

func PostMapStream(res http.ResponseWriter, req *http.Request) {
  // Distribute the original data among the nodes.
  centroids := flow.NewDataset(f, streamingService.stream).Map(func (distStream *stream.Stream) {
    return distStream.GetCentroids()
  })
  // Marshal the data and send it back.
}
