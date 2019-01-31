package main

import (
  cyrand "crypto/rand"
  "encoding/binary"
  "fmt"
  "math/rand"
  "time"
)

var z int

func func_z(z *int, y int, mrng *rand.Rand) {
  amt := time.Duration(mrng.Intn(200))
  time.Sleep(time.Millisecond * amt)
  *z += y
}

func rand_generator() rand.Source {
  var code int64
  err := binary.Read(cyrand.Reader, binary.BigEndian, &code)
  if err != nil {
    fmt.Println(err)
  }
  return rand.NewSource(code)
}

func main() {
  genrated := rand_generator()
  rng  := rand.New(genrated)
  go func_z(&z, 4, rng)
  go func_z(&z, 10, rng)
  amt := time.Duration(rng.Intn(100))
  time.Sleep(time.Millisecond * amt)
  fmt.Println(z)
}

// Race Conditions
// In this code the func_z function is called twice.
// There is a random wait time before the funcution modifies
//  the global z value.
// There is a random wait time after both functions are dispatched
//  before the global z is printed.
// Sample output of multiple runs:
//   E:\Coursera>go run gorutines.go
//   10
//   E:\Coursera>go run gorutines.go
//   14

