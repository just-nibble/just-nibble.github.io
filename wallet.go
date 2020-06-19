package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "log"
  "os"
)

const (
  Port = ":8080"
  )

func makefile() {

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

  f, err := os.Create("test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString(newfile)
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}

func serveDynamic(w http.ResponseWriter, r *http.Request) {
  file, err := ioutil.ReadFile("test.txt")
  if err != nil {
      log.Fatal(err)
  }
  text := string(file)

  fmt.Fprintln(w,text)
  }

func main() {
  http.HandleFunc("/",serveDynamic)
  http.ListenAndServe(Port,nil)
  }