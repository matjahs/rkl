package main

import (
  "log"

  "github.com/matjahs/rkl/cmd"
)

func main() {
  if err := cmd.Execute(); err != nil {
    log.Fatal(err)
  }
}
