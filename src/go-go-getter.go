// example development usage: `go run go-go-getter.go "https://data.cityofchicago.org/resource/f7f2-ggz5.json?fuel_type_code=LPG"`
// one can use the built version like so: `./go-go-getter "https://data.cityofchicago.org/resource/f7f2-ggz5.json?fuel_type_code=LPG"`
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "log"
    "net/http"
)

func main() {
    url := os.Args[1] // takes url as input
    fmt.Println("You typed: ", url)

    resp, err := http.Get(url)

    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    err = ioutil.WriteFile("output.json", body, 0644)

    if err != nil {
        log.Fatal(err) // exits
    }
}
