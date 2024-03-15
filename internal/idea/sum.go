package idea

import (
  "strconv"
  "errors"
  "fmt"
)

func Sum(argList ...int) int {
  var acc int = 0

  for arg := range argList {
    acc += arg
  }

  return acc
}

func Max(argList ...int) int {
  var max int = argList[0]

  for _, arg := range argList {
    if arg > max {
      max = arg
    }
  }

  return max
}

// create a variadic min function
func Min(argList ...int) int {
  var min int = argList[0]

  for _, arg := range argList {
    if arg < min {
      min = arg
    }
  }

  return min
}


func fibonacci(n int) int {
  if n <= 1 {
    return n
  }

  return fibonacci(n-1) + fibonacci(n-2)
}

// type for RGBColor

type RGBColor struct {
  R int
  G int
  B int
}

// this is a method for RGBColor
func (c RGBColor) Hex() string {
  return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

// hextring to RGBColor
func HexToRGB(hex string) (RGBColor, error) {
  var c RGBColor

  if len(hex) != 7 {
    return c, errors.New("invalid length")
  }

  if hex[0] != '#' {
    return c, errors.New("missing #")
  }

  r, err := strconv.ParseInt(hex[1:3], 16, 0)
  if err != nil {
    return c, err
  }

  g, err := strconv.ParseInt(hex[3:5], 16, 0)
  if err != nil {
    return c, err
  }

  b, err := strconv.ParseInt(hex[5:7], 16, 0)
  if err != nil {
    return c, err
  }

  c.R = int(r)

  c.G = int(g)

  c.B = int(b)

  return c, nil
}

//


// func for convert hexcode to RGBColor

