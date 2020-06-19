package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)


func main() {

  url := "https://sandbox.wallets.africa/transfer/banks/all"
  method := "POST"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
  }
  req.Header.Add("Authorization", "Bearer uvjqzm5xl6bw")
  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  newfile := string(body)

  fmt.Println(newfile)
}
