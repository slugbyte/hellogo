package main

import (
	"fmt"
	"github.com/slugbyte/hellogo/internal/app"
	// "github.com/slugbyte/hellogo/internal/server"
	// "github.com/slugbyte/hellogo/internal/util"
  // "github.com/google/uuid"
)


func main() {
  fmt.Println("...");

  if true {
    wat := true
    fmt.Println(wat)
  }

  app.LuckyNumberServer()

  // handle 5 request and then stop server

  // fmt.Println("sum:", util.Sum(1, 3, 5))
  // fmt.Println("sub:", util.Sub(100, []int{1, 3, 5, 50}))
  // util.RunIterator()

  // util.TellStorySync()
  // util.TellStoryAsync()

  // fmt.Println("\nrace 1")
  // util.FootRace();
  // fmt.Println("\nrace 2")
  // util.FootRace();
  // fmt.Println("\nrace 3")
  // util.FootRace();
  // fmt.Println("\nrace 4")
  // util.FootRace();
  // fmt.Println("\nrace 5")
  // util.FootRace();
  // fmt.Println("\nrace 6")
  // util.FootRace();

  // util.Mine()

  // counter := 0
  // for counter < 1000 {
  //   time.Sleep(time.Millisecond)
  //   fmt.Printf("\rLoading Server: %d%%", (counter / 10)+ 1)
  //   counter += 1
  // }
  // fmt.Println()

  // server.StartServer()
}
