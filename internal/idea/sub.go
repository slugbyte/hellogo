package idea
import (
  // "fmt"
)

func Sub(initial int, argList []int) int {

  var result = initial
  for index := range argList {
    // fmt.Println("cool:", argList[index])
    result -= argList[index]
  }
  return result
}
