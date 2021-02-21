package main

import (
	"fmt"
  "math/rand"
	"net/http"
	"time"
  "strings"
)

var num = 0
var alphabet = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func request(cnt int) {

  num = cnt+1

  url := "http://127.0.0.1:8090"
  fmt.Println("URL:>", url)

  jsn := strings.NewReader(`
		{
			"counter" : ` + fmt.Sprintf("%d", num) + `
		}
	`)

  req, err := http.NewRequest("POST", url, jsn)
  req.Header.Set("X-RANDOM", generate(8))
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      panic(err)
  }
  defer resp.Body.Close()

}


func generate(n int) string {
  s := make([]rune, n)
  for i := range s {
      s[i] = alphabet[rand.Intn(len(alphabet))]
  }
  return string(s)
}

func main() {

	request(num)

  for {
    time.Sleep(1 * time.Minute)
    request(num)
  }

}
