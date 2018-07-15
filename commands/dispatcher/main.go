package main

import (
    "github.com/joho/godotenv"
    "log"
    "walrus/dispatcher"
    "fmt"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  d := dispatcher.NewDispatcher()
  fmt.Println(d)

}
