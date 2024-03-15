package server

import (
  "math"
  "net/http"
  "fmt"
  "strings"
  "errors"
)

type HSLColor struct {
  H float64
  S float64
  L float64
}

type RGBColor struct {
  R uint8
  G uint8
  B uint8
}

func minFloat(floats... float64) float64 {
  min := floats[0]
  for _, value := range floats {
    if value < min {
      min = value
    }
  }
  return min
}

func maxFloat(floats... float64) float64 {
  max := floats[0]
  for _, value := range floats {
    if value > max {
      max = value
    }
  }
  return max
}

func (self *RGBColor) String() string {
  return fmt.Sprintf("rgb(%d, %d, %d)", self.R, self.G, self.B)
}

func (self *RGBColor) RGBToHSL() HSLColor {
  r := float64(self.R) / 255.0
  g := float64(self.G) / 255.0
  b := float64(self.B) / 255.0

  max := maxFloat(r, g, b)
  min := minFloat(r, g, b)

  delta := max - min

  h := 0.0
  s := 0.0
  l := (max + min) / 2.0

  if delta != 0 {
    if l < 0.5 {
      s = delta / (max + min)
    } else {
      s = delta / (2 - max - min)
    }

    if r == max {
      h = (g - b) / delta
    } else if g == max {
      h = 2 + (b - r) / delta
    } else {
      h = 4 + (r - g) / delta
    }

    h *= 60
    if h < 0 {
      h += 360
    }
  }

  return HSLColor{
    H: h,
    S: s,
    L: l,
  }
}

// HSLColor to RGBColor 
func (self *HSLColor) HSLToRGB() RGBColor {
  c := (1 - (2 * self.L - 1)) * self.S
  x := c * (1 - math.Abs(math.Mod(self.H / 60, 2) - 1))
  m := self.L - c / 2

  r := 0.0
  g := 0.0
  b := 0.0

  if self.H >= 0 && self.H < 60 {
    r = c
    g = x
  } else if self.H >= 60 && self.H < 120 {
    r = x
    g = c
  } else if self.H >= 120 && self.H < 180 {
    g = c
    b = x
  } else if self.H >= 180 && self.H < 240 {
    g = x
    b = c
  } else if self.H >= 240 && self.H < 300 {
    r = x
    b = c
  } else if self.H >= 300 && self.H < 360 {
    r = c
    b = x
  }

  return RGBColor{
    R: uint8((r + m) * 255),
    G: uint8((g + m) * 255),
    B: uint8((b + m) * 255),
  }
}



type BodyBuilder struct {
  builder *strings.Builder
}

func NewBodyBuilder() *BodyBuilder {
  builder := &strings.Builder{}
  return &BodyBuilder{
    builder: builder,
  }
}

func (self *BodyBuilder) WriteString(str string) {
  self.builder.Write(([]byte)(str))
}

func (self *BodyBuilder) String() string {
  return self.builder.String()
}

func (self *BodyBuilder) Bytes() []byte {
  return ([]byte)(self.builder.String())
}

type AppState struct {
  counter int
  emailList []string
}

func containsEmail(state *AppState, email string) (int, error) {
  index := -1
  for i, value := range state.emailList {
    if value == email {
      index = i
    }
  }

  if index == -1 {
    return -1, errors.New("email not found")
  }

  return index, nil
}

func addEmail(state *AppState, email string) error {
  _, err := containsEmail(state, email)
  if (err == nil) {
    return errors.New("email already exists")
  }
  state.emailList = append(state.emailList, email)
  return nil
}

func removeEmail(state *AppState, email string) error {
  index, err := containsEmail(state, email)
  if err != nil {
    return err
  }

  state.emailList = append(state.emailList[:index], state.emailList[index+1:]...)
  return nil
}


func createHandlerFunc(state *AppState) http.HandlerFunc {
  return http.HandlerFunc(func (w http.ResponseWriter, req *http.Request) {
    state.counter += 1


    fmt.Println("state.emailList", state.emailList)

    fmt.Println("req count:", state.counter)
    url := req.URL.String()

    fmt.Printf("\n[%d] %s %s\n", state.counter, req.Method, url)

    routeName  := "/email-add/"
    if strings.HasPrefix(url, routeName) {
      email := url[(len(routeName)):]
      err := addEmail(state, email)
      if err != nil {
        fmt.Println("allready has email")
        w.WriteHeader(409)
        return
      }
      body := NewBodyBuilder()
      body.WriteString("added email: ")
      body.WriteString(email)
      body.WriteString("\n")
      fmt.Print(body.String())
      w.Write(body.Bytes())
      // w.WriteHeader(200)
      return
    }

    routeName  = "/email-remove/"
    if strings.HasPrefix(url, routeName) {
      email := url[(len(routeName)):]
      err := removeEmail(state, email)
      if err != nil {
        fmt.Println("email", email,  "not found, cannot remove")
        w.WriteHeader(404)
        return
      }
      body := NewBodyBuilder()
      body.WriteString("removed email: ")
      body.WriteString(email)
      body.WriteString("\n")
      fmt.Print(body.String())
      w.Write(body.Bytes())
      // w.WriteHeader(200)
      return
    }

    if strings.HasPrefix(url, "/email-list") {
      body := NewBodyBuilder()
      if len(state.emailList) == 0 {
        body.WriteString("email list empty\n")
      } else {
        body.WriteString("email list: \n")
        for _, email := range state.emailList {
          body.WriteString("  -- ")
          body.WriteString(email)
          body.WriteString("\n")
        }
        body.WriteString("\n")
      }
      fmt.Print(body.String())
      w.Write(body.Bytes())
      // w.WriteHeader(200)
      return
    }

    w.WriteHeader(404)
  })

}


func StartServer() {
  port := ":6666"

  state := &AppState{
    counter: 0,
    emailList: make([]string, 0),
  }

  server := &http.Server{
    Addr: port,
    Handler: createHandlerFunc(state),
  }

  fmt.Println("serving on", port)
  err := server.ListenAndServe()
  if err != nil {
    panic(fmt.Sprintf("server failed to server: (err [%e])", err))
  }
}

