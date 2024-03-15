package main

import (
	"fmt"
	"github.com/slugbyte/hellogo/internal/app"
	// "github.com/slugbyte/hellogo/internal/server"
	// "github.com/slugbyte/hellogo/internal/idea"
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

  // fmt.Println("sum:", idea.Sum(1, 3, 5))
  // fmt.Println("sub:", idea.Sub(100, []int{1, 3, 5, 50}))
  // idea.RunIterator()

  // idea.TellStorySync()
  // idea.TellStoryAsync()

  // fmt.Println("\nrace 1")
  // idea.FootRace();
  // fmt.Println("\nrace 2")
  // idea.FootRace();
  // fmt.Println("\nrace 3")
  // idea.FootRace();
  // fmt.Println("\nrace 4")
  // idea.FootRace();
  // fmt.Println("\nrace 5")
  // idea.FootRace();
  // fmt.Println("\nrace 6")
  // idea.FootRace();

  // idea.Mine()

  // counter := 0
  // for counter < 1000 {
  //   time.Sleep(time.Millisecond)
  //   fmt.Printf("\rLoading Server: %d%%", (counter / 10)+ 1)
  //   counter += 1
  // }
  // fmt.Println()

  // server.StartServer()
}
