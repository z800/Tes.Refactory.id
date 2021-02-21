package main

import (
  "fmt"
  "net/http"
  "log"
  "os"
  "encoding/json"
  "time"
  "net"
)

type RequestCounter struct {
	Counter int    `json:"counter"`
	Header  string `json:"X-RANDOM"`
}

const (
	LOG_FORMAT = "[%s] Success: POST http://%s %s"
)

func LogToFile(dateTime time.Time, ip string, data RequestCounter) {
	f, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	timeFormattedString := dateTime.Format(time.RFC3339)
	defer f.Close()
	log.SetFlags(0)
	log.SetOutput(f)
	json, _ := json.Marshal(data)
	dataString := string(json)
	log.Printf(LOG_FORMAT, timeFormattedString, ip, dataString)
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func handler (w http.ResponseWriter, r *http.Request){

  if r.Method == "POST" {

  	data := RequestCounter{}
  	err := json.NewDecoder(r.Body).Decode(&data)
  	if err != nil {
  		http.Error(w, err.Error(), http.StatusBadRequest)
  		return
  	}

  	data.Header = r.Header.Get("X-RANDOM")

  	LogToFile(time.Now(), GetLocalIP(), data)
  	w.WriteHeader(http.StatusCreated)

  } else {
  	w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "Bad Request")
  }

}

func main() {

  http.HandleFunc("/", handler)
	fmt.Println("SERVER LISTEN AT PORT: 8090")

  if err := http.ListenAndServe(":8090", nil); err != nil {
    log.Fatal(err)
  }

}
