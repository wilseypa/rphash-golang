package server

import (
  "encoding/json"
  "github.com/go-zoo/bone"
  "io/ioutil"
  "log"
  "net/http"
  "path/filepath"
  "strings"
)

type HTTPHandler func(http.ResponseWriter, *http.Request)

type HTTPServer struct {
  Address  string `json:"address"`
  Port     string `json:"port"`
  SSL      bool   `json:"ssl"`
  Router   map[string]HTTPHandler
  KeyFile  string `json:"keyfile"`
  CertFile string `json:"certfile"`
}

type Config struct {
  Server HTTPServer `json:"server"`
}

func Server() (*Config, error) {
  log.Printf("Config loading ... \n")
  config := &Config{}
  var (
    data []byte
    err  error
  )
  path := filepath.Join("server/config.json")
  if data, err = ioutil.ReadFile(path); err != nil {
    return nil, err
  }
  return config, json.Unmarshal(data, config)
}

func (server *HTTPServer) Listen() error {
  if server.SSL {
    log.Printf("Server HTTPS listening on " + server.Address + server.Port)
    return http.ListenAndServeTLS(server.Port, server.CertFile, server.KeyFile, nil)
  }
  log.Printf("Server HTTP listening on " + server.Address + server.Port)
  return http.ListenAndServe(server.Port, NewRouter(server.Router))
}

func NewRouter(handler map[string]HTTPHandler) *bone.Mux {
  mux := bone.New()
  for key := range handler {
    reqestHandler := handler[key]
    requesyKeys := strings.Split(key, "†")
    requestType, requestUrl := requesyKeys[0], requesyKeys[1]
    switch requestType {
    case "GET":
      mux.Get(requestUrl, http.HandlerFunc(reqestHandler))
      break
    case "POST":
      mux.Post(requestUrl, http.HandlerFunc(reqestHandler))
      break
    case "HANDLE":
      mux.Handle(requestUrl, http.HandlerFunc(reqestHandler))
      break
    }
  }
  return mux
}
