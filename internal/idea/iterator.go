package idea

import (
  "fmt"
)

type iterator struct {
  data []int
  current int
  isComplete bool
}

func NewIterator(data []int) iterator {
   return iterator{
    data: data,
    current: 0,
    isComplete: false,
  }
}

func (self *iterator) Next() (int, bool) {
  if self.isComplete {
    return -1, true
  }

  if self.current < len(self.data) {
    result := self.data[self.current] 
    self.current += 1
    return result, false
  } else {
    self.isComplete = true
    return -1, true
  }

}


func RunIterator() {
  var data = []int{5, 4, 3, 2}
  iter := NewIterator(([]int)(data))

  for {
    value, isComplete := iter.Next()
    if (isComplete) {
      fmt.Println("complete:", value)
      return
    }
    fmt.Println("value:", value)
  }
}
