package main

import (
  "fmt"
  "golang.org/x/net/context"
  "golang.org/x/oauth2/google"
)

func main() {
  // Use oauth2.NoContext if there isn't a good context to pass in.
  ctx := context.TODO()

  ts, err := google.DefaultTokenSource(ctx)
  if err != nil {
    fmt.Print(err)
    return
  }

  token,tErr := ts.Token()
  if tErr != nil {
    fmt.Print(tErr)
    return
  }

  fmt.Println("Successfully retrieved token!")
  fmt.Println(token)
}