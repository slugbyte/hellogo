package util

import  (
  "time"
  "fmt"
  "math/rand"
)

type racer struct {
  name string
  time int64
}


func ragisterRacer(name string, lane chan<-racer, racerCount *int) {
  *racerCount += 1
  go func() {
    random := rand.Int63n(499) + 1
    time.Sleep(time.Duration(random) * time.Millisecond)
    lane <- racer{
      name: name,
      time: random,
    }
  }()
}

func FootRace() {
  racerCount := 0

  lane1 := make(chan racer)
  lane2 := make(chan racer)
  lane3 := make(chan racer)

  fmt.Println("on your mark, get set, go")
  ragisterRacer("tina", lane1, &racerCount)
  ragisterRacer("corey", lane2, &racerCount)
  ragisterRacer("monica", lane3, &racerCount)

  var dnfCount = 0
  var place = 1
  for i := 0; i<racerCount; i++ {
    select {
      case a:= <-lane1:
        fmt.Println("place", place, "->", a.name, "time", a.time)
        place++
      case a:= <-lane2:
        fmt.Println("place", place, "->", a.name, "time", a.time)
        place++
      case a:= <-lane3:
        fmt.Println("place", place, "->", a.name, "time", a.time)
        place++
      case <- time.After(time.Duration(150) * time.Millisecond): // this is a timeout  
        dnfCount++
    }
  }
  if dnfCount > 0 {
    fmt.Println(dnfCount,  "racers DNF")
  } else {
    fmt.Println("what a performance")
  }
}
