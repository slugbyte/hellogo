package idea

import (
  "fmt"
  // "time"
)

func writeStorySync(book chan<- string, doneWrite chan <-bool) {
  fmt.Println("a")
  book <- "it was raining"
  fmt.Println("b")
  book <- "it was poring"
  fmt.Println("c")
  book <- "the old man was snoring"
  fmt.Println("d")
  doneWrite <- true
}

func readStorySync(book <-chan string, doneWrite <-chan bool, doneRead chan<- bool) {
  <-doneWrite
  fmt.Println("1|", <-book)
  fmt.Println("2|", <-book)
  fmt.Println("3|", <-book)
  doneRead <- true
}

func TellStorySync() {
  fmt.Println("---ChOneRun---")

  book := make(chan string, 3)
  doneWrite:= make(chan bool) 
  doneRead:= make(chan bool) 

  go writeStorySync(book, doneWrite)
  go readStorySync(book, doneWrite, doneRead)
  <-doneRead
  fmt.Println("the end")
}

func writeStoryAsync(book chan<- string) {
  fmt.Println("a")
  book <- "it was raining"
  fmt.Println("b")
  book <- "it was poring"
  fmt.Println("c")
  book <- "the old man was snoring"
  fmt.Println("d")
}

func readStoryAsync(book <-chan string,  doneRead chan<- bool) {
  fmt.Println("1|", <-book)
  fmt.Println("2|", <-book)
  fmt.Println("3|", <-book)
  doneRead <- true
}


func TellStoryAsync() {
  fmt.Println("---ChOneRun---")

  book := make(chan string)
  doneRead:= make(chan bool) 

  go writeStoryAsync(book)
  go readStoryAsync(book, doneRead)
  <-doneRead
  fmt.Println("the end")
}
