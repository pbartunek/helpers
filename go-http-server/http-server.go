package main

import (
  "net/http"
  "flag"
  "fmt"
)

func main() {

  port := flag.String("p", "8090", "TCP port to bind")
  addr := flag.String("a", "", "IP address to bind")
  dir := flag.String("d", ".", "www root")
  flag.Parse()

  fs := http.FileServer(http.Dir(*dir))
  http.Handle("/", http.StripPrefix("/", fs))

  fmt.Println("Starting a web server on " + *addr + ":" + *port + " in " + *dir + " directory")
  http.ListenAndServe(*addr + ":" + *port, nil)
}
