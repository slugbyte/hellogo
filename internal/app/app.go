package app

import (
  "fmt"
  "math/rand"
  "strconv"
  "net/http"
  "context"
	// "testing"
  "github.com/google/uuid"

)

type Server struct {
	id         uuid.UUID      
	httpServer *http.Server   
	mux        *http.ServeMux 

	addr    string          
	started bool            
	context context.Context 
}


func NewServer(addr string) *Server {
  mux := http.NewServeMux()
  return &Server{
    id: uuid.New(),
    addr: addr,
    mux: mux,
    context: nil,
    httpServer: &http.Server{
      Addr: addr,
      Handler: mux,
    },
  }
}


func (self *Server) AddRoute(path string, handler func(w http.ResponseWriter, r *http.Request)) {
  self.mux.HandleFunc(path, http.HandlerFunc(handler))
}

func (self *Server) Start() {
  fmt.Println("Server started", self.addr, self.id.String())
  fmt.Println("I will server 5 lucky numbers and then shutdown")
  self.httpServer.ListenAndServe()
}

func (self *Server) Stop() {
  self.httpServer.Shutdown(self.context)
  fmt.Println("Server stopped")
}

func LuckyNumberServer() {
  chLucky := make(chan int, 6)
  s := NewServer(":8080")
  s.AddRoute("/", func(w http.ResponseWriter, r *http.Request) { 
    lucky := rand.Int()
    chLucky <- lucky
    w.Write([]byte("lucky number: "))
    w.Write([]byte(strconv.Itoa(lucky)))
    w.Write([]byte("\n"))
  })


  go s.Start()
  total := 0
  for luckyNumber := range chLucky {
    total += 1
    fmt.Println("handle request", total, "lucky number:", luckyNumber)
    if (total == 5)  {
      s.Stop();
      close(chLucky)
    }
  }

}
