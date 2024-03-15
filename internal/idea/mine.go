package idea

import (
  "fmt"
  "time"
  "math/rand"
  "sync"
)

type employe struct {
  mutex sync.Mutex
  name string
  minutesWorked int
  oreCollected int
  locationCount int
  locationVisitList []string
}

func workAtLocation(miner *employe, whereToWork <-chan string,  wg *sync.WaitGroup ){
  for location := range whereToWork {
    minutesWorked := rand.Intn(2000)
    ore := rand.Intn(100)
    time.Sleep(time.Duration((int)(minutesWorked)) * time.Millisecond)

    miner.mutex.Lock()
    miner.minutesWorked += minutesWorked
    miner.locationCount += 1
    miner.oreCollected += ore
    miner.locationVisitList = append(miner.locationVisitList, location)
    miner.mutex.Unlock()
    fmt.Println(miner.name, "found", ore, "kg of ore at the", location)
    // oreBucket <- ore
  }
  wg.Done()
}

func Mine(){
  // 3 miners need to mine as much as they can from the locations below
  // each location should only be visited once by 1 miner
  // miners can go to more than 1 location 

  minerList := []employe{
    {name: "john"},
    {name: "betsy"},
    {name: "tony"},
  }

  locaitonList := []string {
    "hillside",
    "creek bed",
    "river bason",
    "old well",
    "dark cave",
    "miners mesa",
    "spookey holler",
  }

  var wg sync.WaitGroup
  whereToWork := make(chan string, len(locaitonList))
  // oreBucket := make(chan int, len(locaitonList))

  for i := range minerList {
    wg.Add(1)
    go workAtLocation(&minerList[i], whereToWork, &wg)
  }

  for _, location := range locaitonList {
    whereToWork <- location 
  }
  close(whereToWork)


  totalOre := 0
  wg.Wait()

  // for range locaitonList {
  //   totalOre  += <-oreBucket
  // }


  fmt.Println()
  for _, miner := range minerList {
    fmt.Println(miner.name, "worked for", miner.minutesWorked / 60, "hours at", miner.locationCount, "locations and found", miner.oreCollected, "kg of ore")
    fmt.Print(miner.name, "worked at ( ")
    for _, location := range miner.locationVisitList {
      fmt.Print("\"", location, "\" ")
    }
    fmt.Println(")");
    totalOre += miner.oreCollected
  }
  fmt.Println()
  fmt.Println(len(minerList), "workers found", totalOre, "kg of ore at", len(locaitonList), "locations")
}
